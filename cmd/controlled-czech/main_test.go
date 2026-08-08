package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sinfin/controlled-czech/internal/lint"
)

func TestRunCheckJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "doc.md")
	if err := os.WriteFile(path, []byte("Systém musí odpovědět dostatečně rychle."), 0o644); err != nil {
		t.Fatal(err)
	}

	var stdout, stderr strings.Builder
	code := run([]string{"check", "--format", "json", path}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("očekáván exit 0, získán %d, stderr=%s", code, stderr.String())
	}

	var findings []lint.Finding
	if err := json.Unmarshal([]byte(stdout.String()), &findings); err != nil {
		t.Fatalf("neplatný JSON: %v; výstup=%s", err, stdout.String())
	}
	if len(findings) == 0 || findings[0].Rule != "CC603" {
		t.Fatalf("očekáván nález CC603, získáno %#v", findings)
	}
}

func TestRunFailOnWarning(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "doc.md")
	if err := os.WriteFile(path, []byte("Je důležité poznamenat, že služba běží."), 0o644); err != nil {
		t.Fatal(err)
	}

	var stdout, stderr strings.Builder
	code := run([]string{"check", "--fail-on-warning", path}, &stdout, &stderr)
	if code != 1 {
		t.Fatalf("očekáván exit 1, získán %d", code)
	}
}

func TestCollectFilesFromDirectoryOnlyMarkdown(t *testing.T) {
	dir := t.TempDir()
	md := filepath.Join(dir, "a.md")
	txt := filepath.Join(dir, "b.txt")
	if err := os.WriteFile(md, []byte("text"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(txt, []byte("text"), 0o644); err != nil {
		t.Fatal(err)
	}

	files, err := collectFiles([]string{dir})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || files[0] != md {
		t.Fatalf("očekáván pouze %s, získáno %#v", md, files)
	}
}

func TestUnknownCommand(t *testing.T) {
	var stdout, stderr strings.Builder
	code := run([]string{"unknown"}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("očekáván exit 2, získán %d", code)
	}
}
