# Přispívání

Controlled Czech je otevřený návrh. Vítané jsou zejména:

- konkrétní příklady nejasných nebo zbytečně rozvláčných technických formulací,
- návrhy pravidel s jasným důvodem a příklady `před / po`,
- zlepšení heuristik linteru s testy,
- opravy falešně pozitivních varování,
- zkušenosti s používáním v dokumentaci a AI instrukcích.

## Změna specifikace

Návrh změny pravidla by měl obsahovat:

1. problém,
2. navržené znění,
3. alespoň jeden nevhodný příklad,
4. alespoň jeden vhodný příklad,
5. dopad na existující pravidla a linter.

Identifikátor existujícího pravidla se nesmí použít pro jiný význam.

## Změna linteru

Nové chování linteru musí mít test, který před implementací selže a po implementaci projde.

Linter má být konzervativní. Je lepší neoznačit obtížně rozpoznatelné porušení než vytvářet dojem jazykového porozumění, které nástroj nemá.

## Strojová pravidla

Každá fráze v `pravidla/ai-slop.txt` používá formát `CCxxx|fráze`. Identifikátor musí odpovídat pravidlu v `SPEC.md`.

Pořadí diagnostik musí být deterministické. Datová struktura založená na Go `map` proto nesmí přímo určovat pořadí výstupu.
