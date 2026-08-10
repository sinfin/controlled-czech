package lint

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	controlledczech "github.com/sinfin/controlled-czech"
)

const defaultMaxWords = 25

type Options struct {
	MaxWords int
}

type Finding struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
	Text    string `json:"text,omitempty"`
}

type phraseRule struct {
	Rule   string
	Phrase string
}

type terminologyRule struct {
	Discouraged string
	Preferred   string
}

var (
	wordRegexp     = regexp.MustCompile(`[\p{L}\p{N}]+`)
	sentenceRegexp = regexp.MustCompile(`[.!?]+`)
	markdownPrefix = regexp.MustCompile(`^\s*(#{1,6}\s+|[-*+]\s+|\d+[.)]\s+|>\s*)`)
)

var vagueTerms = mustLoadLines("pravidla/neurcite-vyrazy.txt")

var aiSlopPhrases = mustLoadPhraseRules("pravidla/ai-slop.txt")

var copulaParaphrases = mustLoadLines("pravidla/opisy-slovesa-byt.txt")

var actorlessPhrases = []string{
	"se vytvoří", "se odešle", "se provede", "se uloží", "se zpracuje",
	"se nastaví", "se přidá", "se odstraní",
}

var discouragedTerms = mustLoadTerminology("pravidla/terminologie.json")

func mustLoadLines(path string) []string {
	data, err := controlledczech.Rules.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var lines []string
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			lines = append(lines, line)
		}
	}
	return lines
}

func mustLoadPhraseRules(path string) []phraseRule {
	data, err := controlledczech.Rules.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var rules []phraseRule
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
			panic(fmt.Sprintf("neplatné pravidlo fráze v %s: %q", path, line))
		}
		rules = append(rules, phraseRule{Rule: strings.TrimSpace(parts[0]), Phrase: strings.TrimSpace(parts[1])})
	}
	return rules
}

func mustLoadTerminology(path string) []terminologyRule {
	data, err := controlledczech.Rules.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var source map[string][]string
	if err := json.Unmarshal(data, &source); err != nil {
		panic(err)
	}
	var result []terminologyRule
	for preferred, discouraged := range source {
		for _, term := range discouraged {
			result = append(result, terminologyRule{Discouraged: term, Preferred: preferred})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Discouraged == result[j].Discouraged {
			return result[i].Preferred < result[j].Preferred
		}
		return result[i].Discouraged < result[j].Discouraged
	})
	return result
}

func Check(file, text string, options Options) []Finding {
	maxWords := options.MaxWords
	if maxWords <= 0 {
		maxWords = defaultMaxWords
	}

	var findings []Finding
	inFence := false
	previousSentence := ""

	for lineIndex, raw := range strings.Split(text, "\n") {
		lineNo := lineIndex + 1
		trimmed := strings.TrimSpace(raw)

		if strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~") {
			inFence = !inFence
			continue
		}
		if inFence || trimmed == "" {
			continue
		}

		content := markdownPrefix.ReplaceAllString(raw, "")
		lower := strings.ToLower(content)

		for _, term := range vagueTerms {
			if containsTerm(lower, term) {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC603",
					Message: fmt.Sprintf("Neurčitý výraz: %q", term), Text: trimmed,
				})
			}
		}

		for _, phrase := range aiSlopPhrases {
			if containsTerm(lower, phrase.Phrase) {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: phrase.Rule,
					Message: fmt.Sprintf("Metatextová nebo AI-slop fráze: %q", phrase.Phrase), Text: trimmed,
				})
			}
		}

		for _, phrase := range copulaParaphrases {
			if containsTerm(lower, phrase) {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC204",
					Message: fmt.Sprintf("Opis slovesa být: %q; pokud opis nepřidává význam, použij prosté „je“", phrase), Text: trimmed,
				})
			}
		}

		for _, phrase := range actorlessPhrases {
			if containsTerm(lower, phrase) {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC201",
					Message: "Pravděpodobně chybí explicitní aktér", Text: trimmed,
				})
				break
			}
		}

		for _, term := range discouragedTerms {
			if containsTerm(lower, term.Discouraged) {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC301",
					Message: fmt.Sprintf("Nepreferovaný termín %q; preferuj %q", term.Discouraged, term.Preferred), Text: trimmed,
				})
			}
		}

		for _, sentence := range splitSentences(content) {
			words := wordRegexp.FindAllString(sentence, -1)
			if len(words) > maxWords {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC102",
					Message: fmt.Sprintf("Dlouhá věta: %d slov (doporučené maximum %d)", len(words), maxWords), Text: strings.TrimSpace(sentence),
				})
			}

			normalized := normalizeSentence(sentence)
			if normalized != "" && normalized == previousSentence && len(words) >= 3 {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC701",
					Message: "Bezprostředně opakovaná věta", Text: strings.TrimSpace(sentence),
				})
			}
			if normalized != "" {
				previousSentence = normalized
			}
		}
	}

	return findings
}

func WriteJSON(w io.Writer, findings []Finding) error {
	if findings == nil {
		findings = []Finding{}
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(findings)
}

func splitSentences(line string) []string {
	parts := sentenceRegexp.Split(line, -1)
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func normalizeSentence(sentence string) string {
	words := wordRegexp.FindAllString(strings.ToLower(sentence), -1)
	return strings.Join(words, " ")
}

func containsTerm(text, term string) bool {
	pattern := `(?i)(^|[^\p{L}\p{N}_])` + regexp.QuoteMeta(term) + `([^\p{L}\p{N}_]|$)`
	return regexp.MustCompile(pattern).FindStringIndex(text) != nil
}
