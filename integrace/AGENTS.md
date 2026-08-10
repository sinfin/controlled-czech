# Controlled Czech pro AGENTS.md

Do `AGENTS.md` nekopírujte celou specifikaci, pokud agent podporuje skills. Dlouhé trvalé instrukce zvyšují spotřebu kontextu i u úloh, které český technický text vůbec nepoužívají.

## Doporučená varianta se skillem

Nainstalujte skill z [`.agents/skills/controlled-czech/`](../.agents/skills/controlled-czech/) do umístění, ze kterého váš agent načítá skills.

- **Codex** načítá projektové skills z `.agents/skills/`.
- **Claude Code** načítá projektové skills z `.claude/skills/` a uživatelské z `~/.claude/skills/`. Adresář `.agents/skills/` nenačítá.

Skill se může aktivovat podle svého popisu bez kopírování pravidel do `AGENTS.md`.

Claude Code standardně nečte `AGENTS.md`. Pokud projekt používá `AGENTS.md`, vytvořte `CLAUDE.md` s řádkem `@AGENTS.md`.

Pokud chcete explicitní trigger, stačí:

```markdown
Při psaní českých technických textů použij skill `controlled-czech`.
```

## Varianta bez podpory skills

Použijte krátký fallback:

```markdown
Při psaní českých technických textů dodržuj Controlled Czech: https://github.com/sinfin/controlled-czech.
Preferuj informační hustotu před stylem. Jedna věta má vyjadřovat jednu hlavní myšlenku. Neopakuj informace. Používej explicitního aktéra, konzistentní terminologii a normativní slovesa `musí`, `nesmí`, `měl by`, `neměl by`, `může`. Při nejasnosti použij příslušné pravidlo `CCxxx` ze specifikace.
```
