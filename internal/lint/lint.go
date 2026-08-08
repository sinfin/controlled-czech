package lint

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
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

var (
	wordRegexp     = regexp.MustCompile(`[\p{L}\p{N}]+`)
	sentenceRegexp = regexp.MustCompile(`[.!?]+`)
	markdownPrefix = regexp.MustCompile(`^\s*(#{1,6}\s+|[-*+]\s+|\d+[.)]\s+|>\s*)`)
)

var vagueTerms = mustLoadLines("pravidla/neurcite-vyrazy.txt")

var aiSlopPhrases = mustLoadLines("pravidla/ai-slop.txt")

var actorlessPhrases = []string{
	"se vytvoří", "se odešle", "se provede", "se uloží", "se zpracuje",
	"se nastaví", "se přidá", "se odstraní", "bude vytvořen", "bude vytvořena",
	"bude odeslán", "bude odeslána",
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

func mustLoadTerminology(path string) map[string]string {
	data, err := controlledczech.Rules.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var source map[string][]string
	if err := json.Unmarshal(data, &source); err != nil {
		panic(err)
	}
	result := make(map[string]string)
	for preferred, discouraged := range source {
		for _, term := range discouraged {
			result[term] = preferred
		}
	}
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

	scanner := bufio.NewScanner(strings.NewReader(text))
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		raw := scanner.Text()
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
			if containsTerm(lower, phrase) {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC705",
					Message: fmt.Sprintf("Metatextová nebo AI-slop fráze: %q", phrase), Text: trimmed,
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

		for discouraged, preferred := range discouragedTerms {
			if containsTerm(lower, discouraged) {
				findings = append(findings, Finding{
					File: file, Line: lineNo, Rule: "CC301",
					Message: fmt.Sprintf("Nepreferovaný termín %q; preferuj %q", discouraged, preferred), Text: trimmed,
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
