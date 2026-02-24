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
source "${__dir}/../libs/logs.sh"

declare -a examples=(
    "alert-rule"
    "custom-panel"
    "custom-query"
    "grafana-agent-overview"
    "linux-node-overview"
    "red-method"
)

for i in "${examples[@]}"
do
    info "Running $i"
    PYTHONPATH=$PYTHONPATH:"${__dir}/../../python" python "${__dir}/../../examples/python/$i/main.py"
done
