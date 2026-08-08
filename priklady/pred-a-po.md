# Před a po

## Dlouhé souvětí

**Před:**

> V rámci nasazení nové verze bychom měli ideálně zajistit, aby se migrace provedla až poté, co budou dokončeny všechny běžící úlohy, protože jinak může dojít k nekonzistenci dat.

**Po:**

> Před nasazením dokončete všechny běžící úlohy. Potom spusťte migraci. Pokud migraci spustíte dříve, data mohou být nekonzistentní.

## Nejasný aktér

**Před:**

> Po dokončení se vytvoří report a odešle se administrátorovi.

**Po:**

> Worker vytvoří report. Worker odešle report administrátorovi.

## AI slop

**Před:**

> Je důležité poznamenat, že v dnešním rychle se měnícím prostředí existuje několik faktorů, které je potřeba zvážit. Pokud jde o databázi, PostgreSQL představuje velmi relevantní možnost.

**Po:**

> **Doporučení:** Použij PostgreSQL.

## Rozhodnutí a důvod

**Před:**

> PostgreSQL je asi nejlepší, protože ho tým zná a tak bude jednodušší ho provozovat.

**Po:**

> **Rozhodnutí:** Použijeme PostgreSQL.  
> **Důvod:** Tým PostgreSQL již provozuje.
