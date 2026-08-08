# Instrukce pro AI agenty

Tento repozitář používá Controlled Czech také pro vlastní vývoj a dokumentaci.

## Český technický text

Při psaní nebo úpravě českého technického textu dodržuj [`SPEC.md`](SPEC.md).

- Preferuj informační hustotu před stylistickou pestrostí.
- Jedna věta má vyjadřovat jednu hlavní myšlenku.
- Používej explicitního aktéra a konzistentní terminologii.
- Neopakuj informaci jinými slovy.
- Odstraň metatext a přechodové fráze bez informační funkce.
- Nepoužívej doslovné překlady anglických idiomů, pokud existuje jednodušší české vyjádření.
- Rozlišuj fakt, požadavek, rozhodnutí, důvod, předpoklad, omezení, otevřenou otázku a příklad.
- Normativní požadavky formuluj pomocí `musí`, `nesmí`, `měl by`, `neměl by` a `může`.
- Pokud odpovídáš na rozhodovací otázku, uveď doporučení před argumentací.

`SPEC.md` je kanonický zdroj. README ani příklady nesmí měnit význam pravidel.

## Linter

Linter v řadě 0.1 musí být deterministický, offline a bez LLM.

Nové chování linteru musí mít regresní test. Stabilní identifikátor `CCxxx` musí odpovídat pravidlu v `SPEC.md`.

## Rozsah projektu

Do veřejného obsahu nepřidávej interní názvy, procesy ani příklady Sinfinu. Sinfin může být uveden jako iniciátor a správce projektu.

## Ověření změn

Před dokončením změny spusť:

```bash
gofmt -w cmd internal rules.go
go test ./...
go vet ./...
go build ./cmd/controlled-czech
```

Při změně dokumentace spusť linter také nad změněnými dokumenty. Nálezy uvnitř záměrně nevhodných příkladů posuď podle kontextu.
