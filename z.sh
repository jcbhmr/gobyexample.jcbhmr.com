#!/usr/bin/env bash
set -Eeuo pipefail
cd "$1"
rm -rf src
mkdir -p src
echo "# Summary" > src/SUMMARY.md
echo "[Go by Example](README.md)" >> src/SUMMARY.md
echo "---" >> src/SUMMARY.md
echo "# Go by Example" > src/README.md
n=0
old_ifs=$IFS
new_ifs=$'\n'
IFS=$new_ifs
for title in $(cat examples.txt); do
    IFS=$old_ifs
    n=$((n+1))
    slug=$(echo "$title" | awk -F'->' '{ print $1 }' | tr '[:upper:]' '[:lower:]' | sed 's/ /-/g' | sed 's|/|-|g' | sed "s/'//g" | sed 's/--/-/g' | sed 's/--/-/g')
    echo "- [$title]($slug.md)" >> src/SUMMARY.md
    go_code=$(cat examples/"$slug"/*.go)
    sh_code=$(cat examples/"$slug"/*.sh)
    echo "# $title" > src/"$slug".md
    echo '```go' >> src/"$slug".md
    echo "$go_code" >> src/"$slug".md
    echo '```' >> src/"$slug".md
    echo '```sh' >> src/"$slug".md
    echo "$sh_code" >> src/"$slug".md
    echo '```' >> src/"$slug".md
    IFS=$new_ifs
done