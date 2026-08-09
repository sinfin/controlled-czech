package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sinfin/controlled-czech/internal/lint"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		writeUsage(stderr)
		return 2
	}
	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		writeUsage(stdout)
		return 0
	}
	if args[0] != "check" {
		fmt.Fprintf(stderr, "Neznámý příkaz: %s\n\n", args[0])
		writeUsage(stderr)
		return 2
	}

	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	fs.SetOutput(stderr)
	format := fs.String("format", "text", "výstup: text nebo json")
	failOnWarning := fs.Bool("fail-on-warning", false, "vrátí exit code 1, pokud existuje varování")
	maxWords := fs.Int("max-words", 25, "doporučený maximální počet slov ve větě")
	if err := fs.Parse(args[1:]); err != nil {
		return 2
	}
	if *format != "text" && *format != "json" {
		fmt.Fprintf(stderr, "Neplatný formát: %s\n", *format)
		return 2
	}
	if fs.NArg() == 0 {
		fmt.Fprintln(stderr, "Chybí cesta k dokumentu nebo adresáři.")
		return 2
	}

	files, err := collectFiles(fs.Args())
	if err != nil {
		fmt.Fprintf(stderr, "Chyba při hledání souborů: %v\n", err)
		return 2
	}
	if len(files) == 0 {
		fmt.Fprintln(stderr, "Nebyl nalezen žádný dokument ke kontrole.")
		return 2
	}

	var findings []lint.Finding
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(stderr, "Nelze přečíst %s: %v\n", path, err)
			return 2
		}
		findings = append(findings, lint.Check(path, string(data), lint.Options{MaxWords: *maxWords})...)
	}

	if *format == "json" {
		if err := lint.WriteJSON(stdout, findings); err != nil {
			fmt.Fprintf(stderr, "Nelze zapsat JSON: %v\n", err)
			return 2
		}
	} else {
		writeText(stdout, findings)
	}

	if *failOnWarning && len(findings) > 0 {
		return 1
	}
	return 0
}

func collectFiles(paths []string) ([]string, error) {
	seen := map[string]bool{}
	var files []string

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			if !seen[path] {
				seen[path] = true
				files = append(files, path)
			}
			continue
		}

		err = filepath.WalkDir(path, func(current string, entry os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				if strings.HasPrefix(entry.Name(), ".") && current != path {
					return filepath.SkipDir
				}
				return nil
			}
			if strings.EqualFold(filepath.Ext(entry.Name()), ".md") && !seen[current] {
				seen[current] = true
				files = append(files, current)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	sort.Strings(files)
	return files, nil
}

func writeText(w io.Writer, findings []lint.Finding) {
	if len(findings) == 0 {
		fmt.Fprintln(w, "Bez nálezů.")
		return
	}
	for _, finding := range findings {
		fmt.Fprintf(w, "%s:%d  %s  %s\n", finding.File, finding.Line, finding.Rule, finding.Message)
		if finding.Text != "" {
			fmt.Fprintf(w, "  %s\n", finding.Text)
		}
	}
	fmt.Fprintf(w, "\n%d varování\n", len(findings))
}

func writeUsage(w io.Writer) {
	fmt.Fprintln(w, "Controlled Czech")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Použití:")
	fmt.Fprintln(w, "  controlled-czech check [volby] <soubor|adresář>...")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Volby příkazu check:")
	fmt.Fprintln(w, "  --format text|json")
	fmt.Fprintln(w, "  --fail-on-warning")
	fmt.Fprintln(w, "  --max-words N")
}
