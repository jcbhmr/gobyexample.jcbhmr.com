#!/usr/bin/env bash
set -Eeuo pipefail
echo "Building source language book"
mdbook build
for po in ./po/*.po; do
    lang=$(basename "$po" .po)
    echo "Building translated book for $lang"
    MDBOOK_BOOK__LANGUAGE="$lang" \
        MDBOOK_OUTPUT__HTML__SITE_URL="/$lang/" \
        MDBOOK_OUTPUT__HTML__REDIRECT='{}' \
        mdbook build --dest-dir "./book/$lang/"
done
echo "Generating report"
i18n-report ./book/report.html ./po/*.po || true
