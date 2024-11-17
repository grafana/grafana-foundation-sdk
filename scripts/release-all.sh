#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch the error in case mysqldump fails (but gzip succeeds) in `mysqldump |gzip`
set -o pipefail

bold=$(tput bold)
normal=$(tput sgr0)

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/libs/logs.sh"
source "${__dir}/versions.sh"

#################
### Usage ###
#################

# LOG_LEVEL=7 ./scripts/release-version.sh v10.2.x

############
### Main ###
############

for version in ${ALL_GRAFANA_VERSIONS//;/ } ; do
  info "🪧 Releasing ${bold}${version}${normal}"
  $__dir/release-version.sh "${version}"
done
