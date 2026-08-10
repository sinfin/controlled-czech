# Instrukce pro AI agenty

Tento repozitář používá Controlled Czech také pro vlastní práci.

Při psaní nebo úpravě českého technického textu použij skill `controlled-czech` v `.agents/skills/controlled-czech/SKILL.md`. `SPEC.md` je kanonický zdroj. Při nejasnosti načti jen příslušné pravidlo `CCxxx`.

Linter v řadě 0.1 musí zůstat deterministický, offline a bez LLM. Nové chování linteru musí mít regresní test. Diagnostika musí odkazovat na odpovídající `CCxxx`.

Do veřejného obsahu nepřidávej interní názvy, procesy ani příklady Sinfinu. Sinfin může být uveden jako iniciátor a správce projektu.

Skill je uložen jen v `.agents/skills/controlled-czech/`. Manifesty pluginů na něj odkazují polem `skills`. Skill nikdy needituj na dvou místech.

Při vydání verze aktualizuj `version` v `.claude-plugin/plugin.json`, `.claude-plugin/marketplace.json` a `.codex-plugin/plugin.json`. Uživatelé Claude Code dostanou aktualizaci pluginu jen při změně `version`.

Po změně manifestů nebo skillu spusť `claude plugin validate .`.

Před dokončením změny spusť:

```bash
gofmt -w cmd internal rules.go
go test ./...
go vet ./...
go build ./cmd/controlled-czech
```

Při změně dokumentace spusť linter také nad změněnými dokumenty.
