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

#################
### Usage ###
#################

# ./scripts/release-mode.sh

############
### Main ###
############

release_marker=".release/tag"

# Make sure the tags are up-to-date
git_run "${FOUNDATION_SDK_PATH}" fetch origin --tags 2> /dev/null

# No release file means no release was ever made: nothing to release.
if [ ! -f "${release_marker}" ]; then
  echo "prepare"
  exit 0
fi

latest_tag=$(git_run "${FOUNDATION_SDK_PATH}" describe --tags --match 'v*.*.*' --abbrev=0 2>/dev/null || echo 'v0.0.0')
current_marker=$(cat "${release_marker}")

debug "latest: $latest_tag"
debug "current_marker: $current_marker"

# The release marker and latest tags are equal: we just merged the
# release PR and already performed the release. Let's prepare a new one.
if [ "${latest_tag}" == "${current_marker}" ]; then
  echo "prepare"
  exit 0
fi

# The release marker and latest tags are different: we just merged the
# release PR and should perform the release now.
echo "release"
exit 0
