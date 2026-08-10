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

Doporučená integrace je **skill**, ne kopie celé specifikace do `AGENTS.md`. Trvalé instrukce mají být krátké. Detailní pravidla se mají načíst jen při práci s českým technickým textem.

### Agent se skills

Repo obsahuje hotový skill v [`.agents/skills/controlled-czech/SKILL.md`](.agents/skills/controlled-czech/SKILL.md).

**Codex:** zkopírujte adresář `.agents/skills/controlled-czech/` do projektu. Codex repo-local skills automaticky objeví. Do počátečního kontextu načte pouze název a popis skillu; celý `SKILL.md` načte až při použití. Viz [dokumentace Codex Skills](https://developers.openai.com/codex/skills/).

**Claude Code:** zkopírujte stejný adresář do projektového `.claude/skills/controlled-czech/`, nebo do uživatelského `~/.claude/skills/controlled-czech/`. Claude Code adresář `.agents/skills/` nenačítá. Viz [dokumentace Claude Code Skills](https://code.claude.com/docs/en/skills).

Claude Code standardně nečte `AGENTS.md`. Pokud projekt používá `AGENTS.md`, vytvořte `CLAUDE.md` s řádkem `@AGENTS.md`.

Tento repozitář udržuje skill v `.agents/skills/` a pro Claude Code jej zpřístupňuje symlinkem `.claude/skills/controlled-czech`.

Pokud agent skills objevuje automaticky, pravidla není nutné duplikovat v `AGENTS.md`. Pro explicitní trigger stačí jedna věta:

```markdown
Při psaní českých technických textů použij skill `controlled-czech`.
```

### Agent bez podpory skills

Použijte krátký fallback:

```markdown
Při psaní českých technických textů dodržuj Controlled Czech: https://github.com/sinfin/controlled-czech.
Preferuj informační hustotu před stylem. Jedna věta má vyjadřovat jednu hlavní myšlenku. Neopakuj informace. Používej explicitního aktéra, konzistentní terminologii a normativní slovesa `musí`, `nesmí`, `měl by`, `neměl by`, `může`. Při nejasnosti použij příslušné pravidlo `CCxxx` ze specifikace.
```

Další hotové varianty jsou v [`integrace/`](integrace/).

Tento repozitář používá stejný princip pro vlastní práci AI agentů: kořenový [`AGENTS.md`](AGENTS.md) obsahuje jen trvalé projektové podmínky a odkaz na repo-local skill. [`CLAUDE.md`](CLAUDE.md) stejné instrukce importuje pro Claude Code.

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

## Roadmap

Části původního konceptu, které záměrně nejsou v 0.1, jsou v [`ROADMAP.md`](ROADMAP.md). Patří mezi ně projektová konfigurace, přesnější Markdown lint, editor/LSP, distribuce, evaly a volitelná sémantická kontrola.

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

Kód linteru je dostupný pod licencí MIT. Strojová data v `pravidla/` jsou dostupná pod MIT nebo CC BY 4.0. Specifikace, dokumentace, příklady a šablony jsou dostupné pod licencí Creative Commons Attribution 4.0 International (CC BY 4.0). Podrobnosti jsou v souboru [`LICENSE`](LICENSE).
