# Příklad technické specifikace

## Účel

Služba přijímá soubory a vytváří jejich náhledy.

## Aktéři

- **Klient:** odesílá soubor.
- **API:** přijímá požadavek.
- **Worker:** vytváří náhled.

## Chování

1. Klient odešle soubor.
2. API uloží původní soubor.
3. API vytvoří úlohu.
4. Worker načte úlohu.
5. Worker vytvoří náhled.
6. Worker označí úlohu jako dokončenou.

## Požadavky

- API musí přijmout soubor do 20 MB.
- API musí odpovědět do 500 ms pro p95 bez započtení uploadu dat.
- Worker musí uchovat původní soubor do dokončení úlohy.
- Worker nesmí přepsat původní soubor.
- Worker může zopakovat neúspěšnou úlohu nejvýše třikrát.

## Rozhodnutí

**Rozhodnutí:** Náhled ukládáme jako WebP.

**Důvod:** Cílové klienty formát podporují.

## Otevřené otázky

- Potřebujeme samostatnou retenci pro původní soubor a náhled?
