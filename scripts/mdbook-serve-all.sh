#!/usr/bin/env bash
set -Eeuo pipefail
mdbook watch & # English (default)
mdbook watch --dest-dir ./book/zh ./zh & # Chinese
mdbook watch --dest-dir ./book/fr ./fr & # French
mdbook watch --dest-dir ./book/it ./it & # Italian
mdbook watch --dest-dir ./book/jp ./jp & # Japanese
mdbook watch --dest-dir ./book/ko ./ko & # Korean
mdbook watch --dest-dir ./book/ru ./ru & # Russian
mdbook watch --dest-dir ./book/uk ./uk & # Ukrainian
mdbook watch --dest-dir ./book/pt-BR ./pt-BR & # Brazilian Portuguese
mdbook watch --dest-dir ./book/vi ./vi & # Vietnamese
mdbook watch --dest-dir ./book/my ./my & # Burmese
wait -n || c=$?
kill $(jobs -p)
exit $c