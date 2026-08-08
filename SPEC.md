# Controlled Czech 0.1 Draft

Tento dokument je kanonická specifikace Controlled Czech.

Controlled Czech je řízená podoba češtiny určená pro technické a informačně husté texty. Neřeší veškerou českou gramatiku ani stylistiku. Omezuje vybrané jazykové volby, které často zvyšují nejednoznačnost, délku nebo náklady na čtení.

## 1. Úrovně pravidel

Každé pravidlo má identifikátor `CCxxx`.

- **Musí** označuje závazné pravidlo Controlled Czech.
- **Měl by** označuje doporučení. Odchylka může být oprávněná.
- **Může** označuje povolenou možnost.

Automatický linter nemusí umět ověřit všechna pravidla. Heuristické kontroly nesmí být vydávány za úplné jazykové porozumění.

## 2. Normativní slovník požadavků

V normativním textu mají následující výrazy stabilní význam:

| Výraz | Význam |
| --- | --- |
| `musí` | závazný požadavek |
| `nesmí` | závazný zákaz |
| `měl by` | doporučené chování; odchylka vyžaduje zdůvodnění |
| `neměl by` | nedoporučené chování; odchylka může být oprávněná |
| `může` | povolená možnost |

Tyto výrazy se píší běžnými malými písmeny. Controlled Czech se zde inspiruje principem BCP 14, ale nepřebírá jeho pravidlo o verzálkách.

## 3. CC1xx — Věty

### CC101 — Jedna věta vyjadřuje jednu hlavní myšlenku

Věta nesmí spojovat několik samostatných požadavků, rozhodnutí nebo kroků jen kvůli plynulosti textu.

**Nevhodné:**

> Server přijme požadavek, vytvoří úlohu, kterou worker zpracuje, a pokud zpracování selže, odešle upozornění.

**Vhodné:**

> Server přijme požadavek. Server vytvoří úlohu. Worker úlohu zpracuje. Pokud zpracování selže, worker odešle upozornění.

### CC102 — Preferuj krátké věty

Věta by měla mít nejvýše 25 slov, pokud delší věta nepřináší jasnou výhodu. Délka je heuristika, ne absolutní měřítko kvality.

### CC103 — Chronologii piš v pořadí akcí

Popis kroků by měl odpovídat pořadí, ve kterém nastanou.

**Nevhodné:** `Před spuštěním migrace dokončete po jejím naplánování běžící úlohy.`

**Vhodné:** `Naplánujte migraci. Dokončete běžící úlohy. Potom spusťte migraci.`

### CC104 — Podmínku formuluj explicitně

Pro běžnou podmínku preferuj konstrukci `Pokud ..., ...`.

**Nevhodné:** `Při případném selhání buildu vznikne incident.`

**Vhodné:** `Pokud build selže, systém vytvoří incident.`

## 4. CC2xx — Aktér a akce

### CC201 — Pojmenuj aktéra

Pokud je aktér důležitý pro pochopení nebo odpovědnost, musí být z věty zřejmý.

**Nevhodné:** `Po dokončení se vytvoří report.`

**Vhodné:** `Worker po dokončení vytvoří report.`

### CC202 — Preferuj aktivní rod

Aktivní rod obvykle lépe ukazuje odpovědnost za akci.

**Nevhodné:** `Ticket bude vytvořen systémem.`

**Vhodné:** `Systém vytvoří ticket.`

### CC203 — Akci pojmenuj konkrétním slovesem

Preferuj `uloží`, `odešle`, `vytvoří`, `zastaví` před obecnými konstrukcemi typu `provede zpracování` nebo `zajistí realizaci`.

## 5. CC3xx — Terminologie

### CC301 — Jeden pojem má jeden preferovaný název

Dokument musí používat konzistentní terminologii, pokud změna názvu nenese význam.

**Nevhodné:** `repozitář`, `repo`, `repository` pro stejný objekt.

**Vhodné:** `repozitář`.

### CC302 — Jeden termín nemá mít více významů

Stejné slovo by v jednom dokumentu nemělo označovat různé doménové objekty.

### CC303 — Odborný termín definuj při prvním důležitém použití

Definice může být krátká.

> **Tenant:** samostatný zákazník systému s izolovanými daty.

### CC304 — Anglický odborný termín používej, pokud je přesnější nebo běžnější

Controlled Czech nezakazuje anglické technické termíny. Zakazuje nekonzistentní nebo samoúčelné střídání názvů.

## 6. CC4xx — Požadavky

### CC401 — Závazný požadavek formuluj pomocí `musí`

**Nevhodné:** `Auditní záznam je uložen po dobu 90 dní.`

**Vhodné:** `Systém musí uchovat auditní záznam 90 dní.`

### CC402 — Závazný zákaz formuluj pomocí `nesmí`

**Nevhodné:** `Není povoleno zapisovat do chráněné větve.`

**Vhodné:** `Agent nesmí zapisovat do chráněné větve.`

### CC403 — Doporučení formuluj pomocí `měl by` nebo `neměl by`

Doporučení nesmí být zaměnitelné se závazným požadavkem.

### CC404 — Povolenou možnost formuluj pomocí `může`

`Může` nesmí současně znamenat pravděpodobnost a povolení v jednom požadavku.

## 7. CC5xx — Typ informace

### CC501 — Rozlišuj současný stav a požadovaný stav

Fakt o existujícím systému nesmí být formulován tak, aby vypadal jako budoucí požadavek.

Pokud by mohl být typ informace nejasný, použij označení `Fakt` nebo `Požadavek`.

> **Fakt:** Služba dnes ukládá soubory 30 dní.  
> **Požadavek:** Služba musí ukládat soubory 90 dní.

### CC502 — Označ rozhodnutí, pokud by mohlo být zaměněno s návrhem

Doporučený zápis:

> **Rozhodnutí:** Data ukládáme do PostgreSQL.

### CC503 — Odděl rozhodnutí od důvodu

**Vhodné:**

> **Rozhodnutí:** Použijeme PostgreSQL.  
> **Důvod:** Tým PostgreSQL již provozuje.

### CC504 — Označ předpoklad

Předpoklad není fakt.

> **Předpoklad:** Jeden tenant má nejvýše 100 aktivních repozitářů.

### CC505 — Označ otevřenou otázku

Neuzavřená otázka nesmí vypadat jako rozhodnutí.

> **Otevřená otázka:** Potřebujeme samostatnou databázi pro auditní data?

### CC506 — Pro složitější chování můžeš použít strukturu `Když / Pokud / Akce / Výsledek`

Tato struktura je vhodná zejména pro stavové a integrační scénáře.

### CC507 — Označ omezení, pokud ovlivňuje návrh nebo použití

Omezení popisuje hranici, kterou návrh musí respektovat, ale samo nemusí být požadavkem na chování systému.

> **Omezení:** Služba nesmí vyžadovat přístup k veřejnému internetu.

### CC508 — Příklad označ jako příklad

Příklad vysvětluje pravidlo nebo chování. Nesmí nahrazovat normativní požadavek ani rozhodnutí.

> **Příklad:** Pro soubor o velikosti 10 MB může klient použít upload jedním požadavkem.

## 8. CC6xx — Nejednoznačnost

### CC601 — Nahraď neurčitou hodnotu měřitelnou hodnotou, pokud je pro požadavek důležitá

**Nevhodné:** `API musí odpovědět dostatečně rychle.`

**Vhodné:** `API musí odpovědět do 500 ms pro p95.`

### CC602 — Nepoužívej neurčitý odkaz, pokud může mít více antecedentů

**Nevhodné:** `Worker jej potom odešle jemu.`

**Vhodné:** `Worker odešle report administrátorovi.`

### CC603 — Omez vágní výrazy

Výrazy jako `nějaký`, `případně`, `v zásadě`, `ideálně`, `dostatečný`, `rozumný`, `relevantní`, `standardní`, `běžný` nebo `časem` musí být posouzeny podle kontextu. Linter je může označit jako varování.

### CC604 — Negaci formuluj přímo

**Nevhodné:** `Není povoleno, aby uživatel odstranil auditní záznam.`

**Vhodné:** `Uživatel nesmí odstranit auditní záznam.`

### CC605 — Překladový anglicismus nepoužívej, pokud má čeština jednodušší vyjádření

Controlled Czech připouští anglické odborné termíny, ale nedoporučuje doslovné překlady anglických idiomů a diskurzních frází.

**Nevhodné:** `Na konci dne je potřeba rozhodnout.`

**Vhodné:** `Musíme rozhodnout.`

## 9. CC7xx — Informační hustota

### CC701 — Neopakuj stejnou informaci

**Nevhodné:**

> Systém musí uchovat auditní záznam. Je tedy nutné zajistit, aby auditní záznam zůstal uložen.

**Vhodné:**

> Systém musí uchovat auditní záznam.

### CC702 — Nepoužívej prázdný úvod

Úvod musí přidat informaci nebo kontext nutný pro pochopení.

**Nevhodné:** `V dnešním rychle se měnícím světě softwarového vývoje je důležité mít spolehlivý systém pro správu incidentů.`

**Vhodné:** `Systém spravuje provozní incidenty.`

### CC703 — Nevysvětluj samozřejmý důsledek hodnoty

**Nevhodné:** `Timeout je 30 sekund. To znamená, že systém čeká maximálně 30 sekund.`

**Vhodné:** `Timeout: 30 s`.

### CC704 — Nepoužívej metatext bez informační funkce

**Nevhodné:** `Podívejme se nyní podrobněji na jednotlivé oblasti.`

Pokud hned následuje nadpis, větu odstraň.

### CC705 — Omez falešné přechody

Fráze typu `je důležité poznamenat`, `za zmínku stojí`, `s ohledem na výše uvedené` nebo `v tomto kontextu` musí přidat význam. Pokud je lze odstranit beze změny informace, odstraň je.

### CC706 — Neopakuj otázku před odpovědí

Začni odpovědí, pokud opakování otázky nepřidává důležitý kontext.

### CC707 — Preferuj odpověď před argumentací

Pokud text směřuje k doporučení nebo rozhodnutí, uveď je před podpůrnými důvody.

**Vhodné:**

> **Doporučení:** Použij PostgreSQL.
>
> Důvody:
> - tým ho již provozuje,
> - požadované datové typy podporuje bez další služby.

### CC708 — Preferuj strukturovanou hodnotu před větou, pokud věta nepřidává význam

**Vhodné:**

> **Timeout:** 30 s  
> **Retence:** 90 dní  
> **Retries:** 3

## 10. Strojová kontrola

Linter v0.1 záměrně nekontroluje kompletní gramatiku ani význam dokumentu.

Může spolehlivě nebo užitečně heuristicky kontrolovat zejména:

- CC102 — příliš dlouhé věty,
- CC201 — vybrané neosobní konstrukce s pravděpodobně chybějícím aktérem,
- CC301 — známé nepreferované termíny,
- CC603 — známé vágní výrazy,
- CC701 — bezprostředně duplicitní věty,
- CC702/CC704/CC705 — známé AI-slop a metatextové fráze.

Varování linteru není důkaz porušení specifikace. Autor musí rozhodnout podle kontextu.

## 11. Stabilita pravidel

Identifikátory pravidel se po zveřejnění nepoužijí pro jiný význam. Zrušené číslo zůstane rezervované.

Dokud je projekt ve verzi `0.x`, může se formulace i rozsah pravidel měnit.
