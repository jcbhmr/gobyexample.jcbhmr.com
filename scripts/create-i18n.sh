#!/usr/bin/env bash
set -Eeuo pipefail
echo "Generating messages.pot"
MDBOOK_OUTPUT='{"xgettext":{}}' mdbook build --dest-dir ./po/
msginit --input ./po/messages.pot --locale "$1" --output-file "./po/$1.po"