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

FOUNDATION_SDK_PATH=${FOUNDATION_SDK_PATH:-'./'}
FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'git@github.com:grafana/grafana-foundation-sdk.git'}
GH_CLI_CMD=${GH_CLI_CMD:-"gh"} # Command used to run `gh` (GitHub cli)

#################
### Usage ###
#################

# ./scripts/open-release-pr.sh

#################
### Utilities ###
#################

function gh_run() (
  local repo_dir=${1}
  shift

  cd "$repo_dir"

  $GH_CLI_CMD "$@"
)

############
### Main ###
############

expected_pr_title="Next release"
release_branch='release-preview'

debug "Ensuring that ${FOUNDATION_SDK_REPO} will be used by gh"
gh_run "${FOUNDATION_SDK_PATH}" repo set-default "${FOUNDATION_SDK_REPO}"

pr_exists=$(gh_run "${FOUNDATION_SDK_PATH}" pr list -S "Next release")

if [ "${pr_exists}" == "" ]; then
  info "Opening release Pull Request"
  gh_run "${FOUNDATION_SDK_PATH}" pr create \
    --base main \
    --head "${release_branch}" \
    --title "${expected_pr_title}"\
    --body "Note to maintainers: merging this PR will trigger the creation of a new release with all the modifications included on this branch. See the [release docs](https://github.com/grafana/grafana-foundation-sdk/blob/main/maintainers/releasing.md) for more information."
else
  debug "PR already exists"
fi
