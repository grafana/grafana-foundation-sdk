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
source "${__dir}/libs/git.sh"

# These environment variables can be used to alter the behavior of the script.

UPDATES_BRANCH=${UPDATES_BRANCH:-"schema-updates"}
UPDATES_TABLE_FILE=${UPDATES_TABLE_FILE:-"workspace/foundation-sdk/updates-table.md"}

FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'git@github.com:grafana/grafana-foundation-sdk.git'}
GH_CLI_CMD=${GH_CLI_CMD:-"gh"} # Command used to run `gh` (GitHub cli)

#################
### Usage ###
#################

# ./scripts/open-update-schemas-pr.sh

############
### Main ###
############

debug "Ensuring that ${FOUNDATION_SDK_REPO} will be used by gh"
$GH_CLI_CMD repo set-default "${FOUNDATION_SDK_REPO}"

echo -e "This PR contains the following updates:\n" > body.md
cat $UPDATES_TABLE_FILE >> body.md

info "Opening schema updates Pull Request"
$GH_CLI_CMD pr create \
  --base main \
  --head "${UPDATES_BRANCH}" \
  --title "Schema updates" \
  --body "$(cat body.md)"
