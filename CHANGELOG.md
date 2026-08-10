# Changelog

## 0.1.2 — Draft

- nové pravidlo CC204: preferuj `je` před opisem slovesa být (`představuje`, `slouží jako`, `funguje jako`),
- linter kontroluje CC204 podle seznamu `pravidla/opisy-slovesa-byt.txt`,
- distribuce jako plugin pro Claude Code: `.claude-plugin/plugin.json` a marketplace `.claude-plugin/marketplace.json`,
- distribuce jako plugin pro Codex a ChatGPT: `.codex-plugin/plugin.json` a katalog `.agents/plugins/marketplace.json`,
- oba manifesty načítají skill z `.agents/skills/`, takže skill nemá duplikát,
- postup instalace pluginu v `README.md`,
- assety značky v `assets/`: ikona `[CC]`, wordmark, OG obrázek a dva screenshoty; zdrojové HTML v `assets/zdroje/`.

## 0.1.1 — Draft

- oprava YAML frontmatteru skillu; neuvozený `description` s dvojtečkou nešel načíst striktním YAML parserem,
- popis skillu nově uvádí účel skillu, ne jen situace použití,
- skill odkazuje na kanonický `SPEC.md` také přes URL pro projekty bez lokální kopie specifikace,
- skill doporučuje spustit linter nad výsledným dokumentem,
- podpora Claude Code: `CLAUDE.md` importuje `AGENTS.md` a skill je dostupný přes `.claude/skills/controlled-czech`,
- postup integrace pro Claude Code v `README.md` a `integrace/AGENTS.md`.

## 0.1.0 — Draft

- první návrh specifikace Controlled Czech,
- pravidla CC1xx až CC7xx,
- normativní slovník požadavků,
- pravidla proti informačnímu šumu a AI slop,
- příklady a šablony,
- tokenově úsporný repo-local skill pro AI agenty,
- fallback integrace pro `AGENTS.md` a system prompt,
- vlastní kořenový `AGENTS.md`, který aplikuje Controlled Czech na tento repozitář,
- první deterministický CLI linter,
- roadmapa pro projektové profily, přesnější Markdown lint, distribuci a volitelnou sémantickou kontrolu.
