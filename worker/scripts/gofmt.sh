#!/usr/bin/env bash
set -eou pipefail

d="${1}"
p="$(fgrep 'module' "${d}/go.mod" | awk '{print $2}')"

go fmt "${d}/..."
gfs="$(find "${d}" -type f -iname "*.go" ! -path "${d}/vendor/*" ! -path "${d}/tools/vendor/*")"

gfs="$(for gf in $gfs ; do
  fgrep -L 'DO NOT EDIT' $gf || true
done)"

gfs="$(for gf in $gfs ; do
  gofumpt -w "${gf}" && fgrep -lw 'import' $gf || true
done)"

for gf in $gfs ; do
   goimports -w "${gf}" && gci write -s standard -s default -s "Prefix(${p})" "${gf}" 2> /dev/null
done
