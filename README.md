# Controlled Czech

**Otevřená specifikace pro srozumitelnou, jednoznačnou a strojově kontrolovatelnou technickou češtinu.**

Controlled Czech je sada pravidel pro texty, u kterých je důležitější přesnost a informační hustota než stylistická pestrost. Hodí se pro technickou dokumentaci, specifikace, zadání, rozhodnutí, runbooky a instrukce pro AI agenty.

Aktuální stav: **0.1 Draft**.

## Proč Controlled Czech

Technický text má předat informaci s co nejmenším úsilím na straně čtenáře.

Přirozená čeština umožňuje synonyma, implicitní význam, dlouhá souvětí a stylistické variace. V běžném jazyce jsou užitečné. Ve specifikacích mohou snižovat přesnost a zpomalovat čtení.

Generativní AI přidala další problém. Jazykové modely často vytvářejí text, který je gramaticky správný a na první pohled dobře strukturovaný, ale je zbytečně dlouhý. Důležitá informace může být skrytá mezi úvodními větami, opakováním, přechodovými frázemi, překladovými anglicismy a vysvětlováním samozřejmého.

Pro tento jev používáme neformální označení **AI slop**. Controlled Czech se ho snaží omezit bez ztráty hodnotné informace.

Controlled Czech upřednostňuje:

- informaci před stylem,
- konkrétní tvrzení před obecným úvodem,
- krátkou větu před složitým souvětím,
- explicitní význam před implikací,
- konzistentní termín před stylistickou variací,
- strukturu před prózou,
- měřitelnou hodnotu před neurčitým přídavným jménem,
- odpověď před argumentací,
- odstranění textu před jeho stylistickým rozšířením.

Cílem není psát primitivní češtinou. Cílem je, aby čtenář našel důležitou informaci rychle a pochopil ji stejně jako autor.

## Rychlý příklad

Místo:

> V rámci nasazení nové verze bychom měli ideálně zajistit, aby se migrace provedla až poté, co budou dokončeny všechny běžící úlohy, protože jinak může dojít k nekonzistenci dat.

Pište:

> Před nasazením dokončete všechny běžící úlohy. Potom spusťte migraci. Pokud migraci spustíte dříve, data mohou být nekonzistentní.

## Pro co se hodí

Controlled Czech je vhodný zejména pro:

- technickou dokumentaci,
- specifikace a zadání,
- koncepty a návrhy systémů,
- ADR a jiná technická rozhodnutí,
- issue a pull request popisy,
- runbooky a provozní postupy,
- API dokumentaci,
- produktové požadavky,
- interní knowledge base,
- instrukce v `AGENTS.md`, `CLAUDE.md` a system promptech,
- text předávaný mezi AI agenty,
- AI generované analýzy, které má rychle číst člověk.

Společným znakem je požadavek na vysoký poměr **informace / text**.

## Pro co se nehodí

Controlled Czech není obecný návod na dobrou češtinu. Není určen pro texty, kde jsou styl, emoce nebo rytmus součástí sdělení.

Typicky se nehodí pro:

- beletrii,
- marketingový copywriting,
- osobní komunikaci,
- publicistiku,
- eseje,
- storytelling,
- dialog,
- brand voice.

Používejte ho pro technický obsah, ne automaticky pro všechnu firemní komunikaci.

## Použití s AI agenty

Nejjednodušší integrace je odkaz na specifikaci a několik nejdůležitějších pravidel v `AGENTS.md`.

```markdown
### Česká technická dokumentace

Při psaní české technické dokumentace dodržuj Controlled Czech:
https://github.com/sinfin/controlled-czech

- Jedna věta má vyjadřovat jednu hlavní myšlenku.
- Preferuj krátké věty a explicitního aktéra.
- Neopakuj stejnou informaci jinými slovy.
- Odstraň úvody, závěry a přechodové fráze, které nepřidávají informaci.
- Nepoužívej anglické idiomy doslovně přeložené do češtiny.
- Rozlišuj fakt, požadavek, rozhodnutí, předpoklad a otevřenou otázku.
- Pro normativní požadavky používej `musí`, `nesmí`, `měl by`, `neměl by` a `může`.
- Pokud lze informaci vyjádřit hodnotou, tabulkou nebo krátkým seznamem, neobaluj ji zbytečnou prózou.
```

Hotové varianty jsou v [`integrace/`](integrace/).

## Specifikace

Kanonický dokument je [`SPEC.md`](SPEC.md). Každé pravidlo má stabilní identifikátor ve tvaru `CCxxx`.

Controlled Czech rozlišuje:

- **normativní pravidla** — určují význam Controlled Czech,
- **doporučení** — zlepšují čitelnost, ale mohou mít oprávněnou výjimku,
- **heuristiky linteru** — automatické kontroly, které záměrně pokrývají jen část specifikace.

## Linter

Součástí projektu je malý deterministický linter. Nepoužívá LLM a text automaticky nepřepisuje.

```bash
go run ./cmd/controlled-czech check README.md
```

Kontrola více cest:

```bash
go run ./cmd/controlled-czech check README.md SPEC.md priklady/
```

JSON výstup:

```bash
go run ./cmd/controlled-czech check --format json dokumentace/
```

CI režim:

```bash
go run ./cmd/controlled-czech check --fail-on-warning dokumentace/
```

Po vydání tagované verze bude možné použít také:

```bash
go install github.com/sinfin/controlled-czech/cmd/controlled-czech@latest
```

## Inspirace

Controlled Czech není překlad ani česká verze ASD-STE100. Vznikl nezávisle a používá vlastní pravidla pro češtinu.

Projekt se inspiruje zejména:

- [ASD-STE100 Simplified Technical English](https://www.asd-ste100.org/) — controlled natural language pro technickou dokumentaci; aktuální Issue 9 vyšlo v lednu 2025,
- [ISO 24495-1:2023 Plain language](https://www.iso.org/standard/78907.html) — obecné principy srozumitelného jazyka použitelné také pro technické psaní a controlled languages,
- [RFC 2119](https://www.rfc-editor.org/info/rfc2119/) a [RFC 8174](https://www.rfc-editor.org/info/rfc8174/) — inspirace pro jednoznačné rozlišování úrovně požadavků.

## Původ projektu

Controlled Czech inicioval a spravuje [Sinfin](https://sinfin.digital/).

Projekt vznikl z praktické potřeby psát technické koncepty a dokumentaci tak, aby byly rychle čitelné člověkem a současně méně nejednoznačné pro AI systémy.

## Licence

Kód linteru je dostupný pod licencí MIT. Specifikace, dokumentace, pravidla, příklady a šablony jsou dostupné pod licencí Creative Commons Attribution 4.0 International (CC BY 4.0). Podrobnosti jsou v souboru [`LICENSE`](LICENSE).
