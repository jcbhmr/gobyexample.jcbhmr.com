#!/usr/bin/env bash
set -Eeuo pipefail
echo "Generating messages.pot"
MDBOOK_OUTPUT='{"xgettext":{}}' mdbook build --dest-dir ./po/
for po in ./po/*.po; do
    lang=$(basename "$po" .po)
    echo "Merging messages.pot with $lang.po"
    msgmerge --update "$po" ./po/messages.pot
done