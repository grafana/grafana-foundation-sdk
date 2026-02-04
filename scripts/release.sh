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

DRY_RUN=${DRY_RUN:-"yes"} # Some kind of fail-safe to ensure that we're only pushing something when we mean it.

FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'git@github.com:grafana/grafana-foundation-sdk.git'}

CLEANUP_WORKSPACE=${CLEANUP_WORKSPACE:-"yes"} # Should the workspace be deleted after the script runs?
WORKSPACE_PATH=${WORKSPACE_PATH:-'./workspace'}

#################
### Usage ###
#################

# LOG_LEVEL=7 ./scripts/release.sh

#################
### Utilities ###
#################

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

function clone_foundation_sdk() {
  local clone_into_dir="${1}"
  shift

  git clone "${FOUNDATION_SDK_REPO}" "${clone_into_dir}"
}

############
### Main ###
############

foundation_sdk_path="${WORKSPACE_PATH}/foundation-sdk"
release_file_marker=".release/tag"

if [ ! -f ".release/tag" ]; then
    notice "Release marker not found, aborting."
    exit 0
fi

if [ "${DRY_RUN}" == "no" ]; then
  warning "Dry-run is OFF."
else
  notice "Dry-run is ON."
fi

debug "workspace path: ${WORKSPACE_PATH}"

# Just in case there are leftovers from a previous run.
rm -rf "${WORKSPACE_PATH}"

info "Cloning grafana-foundation-sdk into ${foundation_sdk_path}"
clone_foundation_sdk "${foundation_sdk_path}"

next_tag=$(cat "${foundation_sdk_path}/${release_file_marker}")

git_run "${foundation_sdk_path}" rm "${release_file_marker}"
git_run "${foundation_sdk_path}" commit -m "Remove release marker for ${next_tag}"

run_when_safe git_run "${foundation_sdk_path}" tag "${next_tag}"

run_when_safe git_run "${foundation_sdk_path}" push origin main
run_when_safe git_run "${foundation_sdk_path}" push origin --tags
