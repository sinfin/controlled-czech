# Instrukce pro AI agenty

Tento repozitář používá Controlled Czech také pro vlastní práci.

Při psaní nebo úpravě českého technického textu použij skill `controlled-czech` v `.agents/skills/controlled-czech/SKILL.md`. `SPEC.md` je kanonický zdroj. Při nejasnosti načti jen příslušné pravidlo `CCxxx`.

Linter v řadě 0.1 musí zůstat deterministický, offline a bez LLM. Nové chování linteru musí mít regresní test. Diagnostika musí odkazovat na odpovídající `CCxxx`.

Do veřejného obsahu nepřidávej interní názvy, procesy ani příklady Sinfinu. Sinfin může být uveden jako iniciátor a správce projektu.

Před dokončením změny spusť:

```bash
gofmt -w cmd internal rules.go
go test ./...
go vet ./...
go build ./cmd/controlled-czech
```

Při změně dokumentace spusť linter také nad změněnými dokumenty.
