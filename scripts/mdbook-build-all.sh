#!/usr/bin/env bash
set -Eeuo pipefail
mdbook build # English (default)
# mdbook build ./zh # Chinese
mdbook build ./fr # French
mdbook build ./it # Italian
# mdbook build ./ja # Japanese
# mdbook build ./ko # Korean
# mdbook build ./ru # Russian
# mdbook build ./uk # Ukrainian
# mdbook build ./pt-BR # Brazilian Portuguese
# mdbook build ./vi # Vietnamese
# mdbook build ./my # Burmese
