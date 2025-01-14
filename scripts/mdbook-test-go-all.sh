#!/usr/bin/env bash
set -Eeuo pipefail
./scripts/mdbook-test-go.go # English (default)
./scripts/mdbook-test-go.go --dest-dir ./book/zh ./zh # Chinese
./scripts/mdbook-test-go.go --dest-dir ./book/fr ./fr # French
./scripts/mdbook-test-go.go --dest-dir ./book/it ./it # Italian
./scripts/mdbook-test-go.go --dest-dir ./book/jp ./jp # Japanese
./scripts/mdbook-test-go.go --dest-dir ./book/ko ./ko # Korean
./scripts/mdbook-test-go.go --dest-dir ./book/ru ./ru # Russian
./scripts/mdbook-test-go.go --dest-dir ./book/uk ./uk # Ukrainian
./scripts/mdbook-test-go.go --dest-dir ./book/pt-BR ./pt-BR # Brazilian Portuguese
./scripts/mdbook-test-go.go --dest-dir ./book/vi ./vi # Vietnamese
./scripts/mdbook-test-go.go --dest-dir ./book/my ./my # Burmese
