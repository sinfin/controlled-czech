# Assety

Značka Controlled Czech používá hranatou závorku jako znak omezení. Wordmark je `[Controlled Czech]`, krátká značka `[CC]`.

| Soubor | Rozměr | Použití |
| --- | --- | --- |
| `icon.png` | 512 × 512 | ikona pluginu v katalogu a v pickeru |
| `logo-dark.png` | 1600 × 285 | wordmark na tmavém pozadí |
| `logo-light.png` | 1600 × 285 | wordmark na světlém pozadí |
| `og.png` | 1200 × 630 | náhled při sdílení odkazu |
| `screenshot-linter.png` | 1600 × 1000 | výstup `controlled-czech check` |
| `screenshot-pred-a-po.png` | 1600 × 760 | srovnání textu před úpravou a po úpravě |

Barvy: černá `#000000`, bílá `#FFFFFF`, světlé pozadí `#F5F5F5`, šedý text `#6F7172`. V ukázkách značí varování `#FF9A52`, potvrzení `#219D61` a odkaz `#4C84FD`.

## Generování

Zdroje jsou v `zdroje/`. Každý asset je HTML stránka, kterou vykreslí headless Chrome:

```bash
CHROME="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
"$CHROME" --headless --hide-scrollbars --allow-file-access-from-files \
  --screenshot=icon.png --window-size=512,512 zdroje/icon.html
```

Wordmark po vykreslení obřízni a doplň okraj:

```bash
magick logo-dark-raw.png -trim +repage -bordercolor black -border 60 -resize 1600x logo-dark.png
```

Text ve `screenshot-linter.html` odpovídá skutečnému výstupu linteru. Při změně diagnostiky vygeneruj výstup znovu a stránku aktualizuj.

## Písmo

Assety používají písmo ModernGothic, které není součástí repozitáře. Pro vykreslení vlož soubory `ModernGothic-Regular.woff2`, `ModernGothic-Medium.woff2` a `ModernGothic-Bold.woff2` do `zdroje/fonts/`. Bez nich Chrome vykreslí náhradní bezpatkové písmo.

Repozitář záměrně neobsahuje ani soubory písma, ani obtažené obrysy glyfů v SVG. Rastry v tomto adresáři jsou grafika značky, ne distribuce písma.
