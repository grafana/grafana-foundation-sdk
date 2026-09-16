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

# These environment variables can be used to alter the behavior of the release script.

COG_CMD=${COG_CMD:-"cog"} # Command used to run `cog`
CODEGEN_PIPELINE_CONFIG=${CODEGEN_PIPELINE_CONFIG:-".cog/config.yaml"} # Codegen pipeline config file to use.

FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'git@github.com:grafana/grafana-foundation-sdk.git'}

WORKSPACE_PATH=${WORKSPACE_PATH:-'./workspace'}
DRY_RUN=${DRY_RUN:-"yes"} # Some kind of fail-safe to ensure that we're only pushing something when we mean it.

#################
### Usage ###
#################

# LOG_LEVEL=7 ./scripts/update-schemas.sh

#################
### Utilities ###
#################

function clone_foundation_sdk() {
  local clone_into_dir="${1}"
  shift

  git clone "${FOUNDATION_SDK_REPO}" "${clone_into_dir}"
  git_run "${clone_into_dir}" fetch origin --tags
}

function run_when_safe() {
  local command=${1}
  shift

  if [ "${DRY_RUN}" == "no" ]; then
    ${command} "$@"
  else
    warning "Dry run enabled: skipping execution of \"${command} $*\""
    info "Run this script with DRY_RUN=no to disable dry-run mode."
  fi
}

############
### Main ###
############

foundation_sdk_path="${WORKSPACE_PATH}/foundation-sdk"
update_branch="schema-updates-$(date +'%F')"

if [ "${DRY_RUN}" == "no" ]; then
  warning "Dry-run is OFF."
else
  notice "Dry-run is ON."
fi

debug "workspace path: ${WORKSPACE_PATH}"
debug "update branch: ${update_branch}"

# Just in case there are leftovers from a previous run.
rm -rf "${WORKSPACE_PATH}"

info "Cloning grafana-foundation-sdk into ${foundation_sdk_path}"
clone_foundation_sdk "${foundation_sdk_path}"

info "Identifying schema updates"
current_work_dir="$(pwd)"
cd "${foundation_sdk_path}" && ${COG_CMD} config update-inputs -vvv --config "${CODEGEN_PIPELINE_CONFIG}" --write > updates-table.md && cd "${current_work_dir}"

pwd

# The file is empty
if [ ! -s "${foundation_sdk_path}/updates-table.md" ]; then
    info "No updates found."
    exit 0
fi

info "Updates found"

info "Creating new schemas update branch"
git_run "${foundation_sdk_path}" checkout -b "${update_branch}"

debug "Adding changes to git staging area"
git_run "${foundation_sdk_path}" add .cog

debug "Publish new '${update_branch}' branch"
run_when_safe git_run "${foundation_sdk_path}" push origin "${update_branch}"
