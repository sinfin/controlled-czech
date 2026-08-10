---
name: controlled-czech
description: "Pravidla Controlled Czech pro srozumitelnou a informačně hustou technickou češtinu. Použij při psaní, úpravě nebo revizi českých technických textů: dokumentace, specifikace, koncepty, ADR, issue, PR popisy, runbooky, produktové požadavky a informačně husté AI výstupy. Nepoužívej pro marketing, beletrii, osobní komunikaci ani brand voice."
---

# Controlled Czech

Cíl: maximum informace při minimálním čtecím úsilí.

1. Uveď odpověď, doporučení nebo rozhodnutí před argumentací.
2. Jedna věta vyjadřuje jednu hlavní myšlenku.
3. Pojmenuj aktéra a preferuj aktivní rod.
4. Používej jeden termín pro jeden pojem.
5. Neopakuj informaci jinými slovy.
6. Odstraň úvody, shrnutí, metatext a přechody bez nové informace.
7. Nahraď vágní výraz měřitelnou hodnotou, pokud je hodnota podstatná.
8. Nepřekládej doslovně anglické idiomy, pokud existuje přirozené české vyjádření.
9. Rozlišuj fakt, požadavek, rozhodnutí, důvod, předpoklad, omezení, otevřenou otázku a příklad.
10. Požadavky formuluj pomocí `musí`, `nesmí`, `měl by`, `neměl by` a `může`.

Kanonická specifikace s pravidly `CCxxx`:

1. `SPEC.md` v kořenu projektu, pokud existuje,
2. jinak <https://github.com/sinfin/controlled-czech/blob/main/SPEC.md>.

Při nejasnosti načti jen příslušné pravidlo `CCxxx`, ne celou specifikaci.

Pokud je v projektu dostupný linter `controlled-czech`, spusť nad výsledným dokumentem `controlled-czech check <cesta>`. V repozitáři Controlled Czech použij `go run ./cmd/controlled-czech check <cesta>`.
