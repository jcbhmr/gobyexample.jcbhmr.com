#!/usr/bin/env bash
set -Eeuo pipefail
mdbook watch &
for po in ./po/*.po; do
    lang=$(basename "$po" .po)
    MDBOOK_BOOK__LANGUAGE="$lang" \
        MDBOOK_OUTPUT__HTML__SITE_URL="/$lang/" \
        MDBOOK_OUTPUT__HTML__REDIRECT='{}' \
        mdbook watch --dest-dir "./book/$lang/" &
done
watch "i18n-report ./book/report.html ./po/*.po || true" &>/dev/null &
go run github.com/eliben/static-server -silent ./book/ &
wait -n