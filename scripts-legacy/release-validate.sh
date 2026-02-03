#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch the error in case mysqldump fails (but gzip succeeds) in `mysqldump | gzip`
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/libs/logs.sh"

release_path=${1:-"./"}

find "${release_path}" -maxdepth 1 -mindepth 1 -type d -not -path '*/.*' -print | while read -r target; do
  target=${target#"$release_path/"}

  if [ -e "${__dir}/validate/${target}.sh" ]; then
    info "  → checking ${target}"
    ${__dir}/validate/${target}.sh "${release_path}"
  else
    warning "  → skipping ${target}: no validator file found"
  fi
done