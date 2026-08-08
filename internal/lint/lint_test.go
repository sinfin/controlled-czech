package lint

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLongSentence(t *testing.T) {
	text := "Tato věta obsahuje opravdu mnoho slov a pokračuje dál tak dlouho, že překročí doporučený limit pro jednu technickou větu a současně zůstává jedinou větou, kterou má linter označit jako příliš dlouhou pro snadné čtení dokumentace."
	findings := Check("doc.md", text, Options{MaxWords: 25})
	assertHasRule(t, findings, "CC102")
}

func TestVagueExpression(t *testing.T) {
	findings := Check("doc.md", "Systém musí odpovědět dostatečně rychle.", Options{})
	assertHasRule(t, findings, "CC603")
}

func TestAISlopPhrase(t *testing.T) {
	findings := Check("doc.md", "Je důležité poznamenat, že služba používá PostgreSQL.", Options{})
	assertHasRule(t, findings, "CC705")
}

func TestLikelyMissingActor(t *testing.T) {
	findings := Check("doc.md", "Po dokončení se vytvoří report.", Options{})
	assertHasRule(t, findings, "CC201")
}

func TestDiscouragedTerminology(t *testing.T) {
	findings := Check("doc.md", "Agent načte repo a zkontroluje změny.", Options{})
	assertHasRule(t, findings, "CC301")
}

func TestAdjacentDuplicateSentence(t *testing.T) {
	findings := Check("doc.md", "Systém uloží záznam. Systém uloží záznam.", Options{})
	assertHasRule(t, findings, "CC701")
}

func TestFencedCodeIsIgnored(t *testing.T) {
	text := "```text\nJe důležité poznamenat, že se vytvoří report z repo.\n```"
	findings := Check("doc.md", text, Options{})
	if len(findings) != 0 {
		t.Fatalf("očekáváno 0 nálezů v code fence, získáno: %#v", findings)
	}
}

func TestJSONOutputEmptyArray(t *testing.T) {
	var b strings.Builder
	if err := WriteJSON(&b, nil); err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(b.String()) != "[]" {
		t.Fatalf("očekáváno [], získáno %q", strings.TrimSpace(b.String()))
	}
}

func TestJSONOutput(t *testing.T) {
	findings := []Finding{{File: "doc.md", Line: 3, Rule: "CC603", Message: "Neurčitý výraz", Text: "ideálně"}}
	var b strings.Builder
	if err := WriteJSON(&b, findings); err != nil {
		t.Fatal(err)
	}
	var decoded []Finding
	if err := json.Unmarshal([]byte(b.String()), &decoded); err != nil {
		t.Fatalf("výstup není validní JSON: %v", err)
	}
	if len(decoded) != 1 || decoded[0].Rule != "CC603" {
		t.Fatalf("neočekávaný JSON výstup: %#v", decoded)
	}
}

func assertHasRule(t *testing.T, findings []Finding, rule string) {
	t.Helper()
	for _, finding := range findings {
		if finding.Rule == rule {
			return
		}
	}
	t.Fatalf("pravidlo %s nebylo nalezeno v %#v", rule, findings)
}
