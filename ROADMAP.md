# Roadmap

Tento dokument zachycuje části původního konceptu, které záměrně nejsou součástí verze 0.1.

Roadmap není závazný harmonogram. `SPEC.md` zůstává jediným kanonickým zdrojem aktuálních pravidel.

## 0.2 — Přesnost a projektové profily

### Projektová konfigurace

Umožnit projektu rozšířit Controlled Czech bez forku specifikace.

Konfigurace by měla podporovat zejména:

- preferovanou a nepreferovanou terminologii projektu,
- povolené doménové výrazy,
- vlastní vágní výrazy a fráze,
- limit délky věty,
- zapnutí nebo vypnutí vybraných heuristik,
- cesty a soubory zahrnuté nebo vyloučené z kontroly.

Výchozí chování musí zůstat použitelné bez konfiguračního souboru.

### Markdown-aware lint

Současný linter používá úmyslně malou textovou heuristiku. Další verze by měla lépe rozlišovat strukturu Markdownu.

Priorita:

- inline code,
- odkazy a URL,
- správné párování fenced code bloků,
- tabulky,
- citace a příklady,
- explicitní potlačení konkrétního nálezu.

Cílem je umožnit lintovat i celý repozitář Controlled Czech bez varování ze záměrně nevhodných příkladů.

### Čeština a tvarosloví

Rozšířit deterministické kontroly tam, kde to nezvyšuje počet falešných nálezů.

Možné oblasti:

- tvary vágních výrazů,
- tvary preferovaných a nepreferovaných termínů,
- přesnější hranice slov,
- lepší dělení vět u zkratek, desetinných čísel a technických identifikátorů.

Linter nesmí předstírat jazykové porozumění, které nemá.

### Úplnost tvrzení a postupů

Původní koncept používá kontrolní otázky `Kdo? / Co? / Kdy? / Za jakých podmínek?`.

Budoucí sémantická kontrola může upozornit, pokud technické tvrzení nebo krok postupu postrádá údaj důležitý pro provedení nebo odpovědnost. Kontrola nesmí vyžadovat všechny čtyři údaje tam, kde nejsou pro význam potřebné.

### Procedurální text

Upřesnit rozdíl mezi normativním požadavkem a přímým pracovním pokynem.

Například `Systém musí uložit záznam.` je požadavek na systém, zatímco `Spusť migraci.` je krok postupu. Controlled Czech má podporovat oba typy textu bez nucení imperativu do normativního slovníku `musí / nesmí / může`.

## Distribuce a integrace

### Binární vydání

Publikovat sestavení pro běžné platformy bez požadavku na lokální Go toolchain.

Zvážit:

- GitHub Releases,
- Homebrew,
- jednoduchou instalaci pro Linux a Windows,
- oficiální GitHub Action,
- pre-commit integraci.

### Editor a LSP

Po stabilizaci diagnostik zvážit editorovou integraci, která zobrazí `CCxxx` přímo u textu.

Preferovat společné LSP nebo jinou editorově neutrální vrstvu před samostatnou implementací stejné logiky pro každý editor. VS Code rozšíření může být tenký klient nad touto vrstvou.

### Skill jako instalovatelný balíček

Repo-local skill zůstává nejjednodušší variantou pro jednotlivé projekty.

Pro širší distribuci zvážit balíček, který obsahuje skill a potřebné reference bez kopírování celé specifikace do trvalého promptu. Pro ChatGPT a Codex může být vhodným distribučním formátem plugin.

## 0.3 — Volitelná sémantická kontrola

Deterministický linter musí zůstat výchozí a offline.

Volitelná sémantická vrstva může kontrolovat pravidla, která nelze spolehlivě určit regulárními výrazy nebo jednoduchým parserem. Může používat LLM pouze po explicitním zapnutí.

Kandidáti:

- více hlavních myšlenek v jedné větě,
- chybějící nebo nejasný aktér,
- významově opakovaná informace,
- záměna faktu, požadavku, rozhodnutí a předpokladu,
- skrytá podmínka,
- nejasný antecedent,
- doporučení schované za dlouhou argumentací,
- kontrola struktury `Když / Pokud / Akce / Výsledek`,
- kontrola úplnosti `Kdo? / Co? / Kdy? / Za jakých podmínek?` podle kontextu.

Sémantická kontrola musí u každého nálezu uvést odpovídající `CCxxx`. Nesmí tiše měnit text.

## Korpus a evaly

Před rozšiřováním heuristik vytvořit malý veřejný korpus českých technických textů.

Korpus by měl obsahovat:

- lidsky psanou dokumentaci,
- AI generované texty,
- vhodné i nevhodné příklady,
- očekávané nálezy podle `CCxxx`,
- příklady, které nesmí být označeny.

Korpus umožní měřit falešné pozitivní a falešné negativní nálezy. Souhrnné skóre čitelnosti nebo informační hustoty má smysl až tehdy, pokud bude jeho význam vysvětlitelný a ověřený na datech.

## Návrhy oprav

Linter může v budoucnu nabídnout explicitní návrh jednodušší formulace.

Návrh musí být oddělený od samotného nálezu. Automatické přepisování dokumentu bez kontroly uživatele není cílem projektu.

## Rozšíření šablon

Podle zkušeností z reálných projektů lze doplnit volitelné části šablon, například:

- datový model,
- rozhraní,
- příklady,
- bezpečnostní omezení,
- provozní dopady.

Šablony nemají nutit každý dokument k vyplnění všech sekcí.

## Mimo scope

Controlled Czech zůstává projektem pro češtinu. Není cílem vytvářet obecný framework pro další jazyky.

Projekt nemá vyžadovat:

- SaaS službu,
- proprietární backend,
- povinné LLM,
- kompletní kontrolu české gramatiky nebo pravopisu,
- používání Controlled Czech pro marketing, beletrii nebo osobní komunikaci.
