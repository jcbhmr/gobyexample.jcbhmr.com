#!/usr/bin/env bash
set -Eeuo pipefail
mdbook build # English (default)
mdbook build --dest-dir ./book/zh ./zh # Chinese
mdbook build --dest-dir ./book/fr ./fr # French
mdbook build --dest-dir ./book/it ./it # Italian
mdbook build --dest-dir ./book/ja ./ja # Japanese
mdbook build --dest-dir ./book/ko ./ko # Korean
mdbook build --dest-dir ./book/ru ./ru # Russian
mdbook build --dest-dir ./book/uk ./uk # Ukrainian
mdbook build --dest-dir ./book/pt-BR ./pt-BR # Brazilian Portuguese
mdbook build --dest-dir ./book/vi ./vi # Vietnamese
mdbook build --dest-dir ./book/my ./my # Burmese
