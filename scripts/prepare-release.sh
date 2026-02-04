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

COG_CMD=${COG_CMD:-"cog"} # Command used to run `cog`
GH_CLI_CMD=${GH_CLI_CMD:-"gh"} # Command used to run `gh` (GitHub cli)

CODEGEN_PIPELINE_CONFIG=${CODEGEN_PIPELINE_CONFIG:-"${__dir}/../.cog/post-12/config.yaml"} # Codegen pipeline config file to use.

KIND_REGISTRY_PATH=${KIND_REGISTRY_PATH:-'../kind-registry'} # Path to the kind-registry

KIND_REGISTRY_REPO=${KIND_REGISTRY_REPO:-'https://github.com/grafana/kind-registry.git'}
FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'git@github.com:grafana/grafana-foundation-sdk.git'}

SKIP_VALIDATION=${SKIP_VALIDATION:-"no"}
CLEANUP_WORKSPACE=${CLEANUP_WORKSPACE:-"yes"} # Should the workspace be deleted after the script runs?
WORKSPACE_PATH=${WORKSPACE_PATH:-'./workspace'}

#################
### Usage ###
#################

# LOG_LEVEL=7 ./scripts/prepare-release.sh

#################
### Utilities ###
#################

function clone_kind_registry() {
  local clone_into_dir="${1}"
  shift

  git clone --depth 1 "${KIND_REGISTRY_REPO}" "${clone_into_dir}"
}

function clone_foundation_sdk() {
  local clone_into_dir="${1}"
  shift

  git clone "${FOUNDATION_SDK_REPO}" "${clone_into_dir}"
}

function run_codegen() {
  local extra_parameters="${1}"
  shift

  $COG_CMD generate \
    --config "${CODEGEN_PIPELINE_CONFIG}" \
    --parameters "${extra_parameters}"
}

function gh_run() (
  local repo_dir=${1}
  shift

  cd "$repo_dir"

  $GH_CLI_CMD "$@"
)

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

function next_version() {
  local base="$1"
  shift
  local step="$1"
  shift

  local RE='v[^0-9]*\([0-9]*\)[.]\([0-9]*\)[.]\([0-9]*\)\([0-9A-Za-z-]*\)'

  local MAJOR=`echo $base | sed -e "s#$RE#\1#"`
  local MINOR=`echo $base | sed -e "s#$RE#\2#"`
  local PATCH=`echo $base | sed -e "s#$RE#\3#"`

  case "$step" in
    major)
      let MAJOR+=1
      let MINOR=0
      let PATCH=0
      ;;
    minor)
      let MINOR+=1
      let PATCH=0
      ;;
    patch)
      let PATCH+=1
      ;;
  esac

  echo "v$MAJOR.$MINOR.$PATCH"
}

############
### Main ###
############

function cleanup() {
  debug "Cleaning up workspace"
  rm -rf "${WORKSPACE_PATH}"
}
if [ "${CLEANUP_WORKSPACE}" == "yes" ]; then
  trap cleanup EXIT # run the cleanup() function on exit
fi

codegen_output_path="${WORKSPACE_PATH}/codegen"
foundation_sdk_path="${WORKSPACE_PATH}/foundation-sdk"
release_branch='release-preview'
release_file_marker="${foundation_sdk_path}/.release/tag"

if [ -f ".release/tag" ]; then
    notice "Release marker found, aborting."
    exit 0
fi

if [ "${DRY_RUN}" == "no" ]; then
  warning "Dry-run is OFF."
else
  notice "Dry-run is ON."
fi

if [ "${SKIP_VALIDATION}" == "yes" ]; then
  warning "Code validation is OFF."
else
  debug "Code validation is ON."
fi

debug "Release branch: ${release_branch}"
debug "kind-registry path: ${KIND_REGISTRY_PATH}"
debug "workspace path: ${WORKSPACE_PATH}"

# Just in case there are leftovers from a previous run.
rm -rf "${WORKSPACE_PATH}"

if [ ! -d "${KIND_REGISTRY_PATH}" ]; then
  info "Cloning kind-registry into ${KIND_REGISTRY_PATH}"
  clone_kind_registry "${KIND_REGISTRY_PATH}"
fi

info "Cloning grafana-foundation-sdk into ${foundation_sdk_path}"
clone_foundation_sdk "${foundation_sdk_path}"

info "Pulling kind-registry@main"
git_run "${KIND_REGISTRY_PATH}" checkout main
git_run "${KIND_REGISTRY_PATH}" pull --ff-only origin main

release_branch_exists=$(git_has_branch "${foundation_sdk_path}" "${release_branch}")
if [ "$release_branch_exists" != "0" ]; then
  debug "No existing release branch: next version will be determined from the latest tag"
  latest_tag=$(git_run "${foundation_sdk_path}" describe --tags --abbrev=0 2>/dev/null || echo 'v0.0.0')
  next_tag=$(next_version "${latest_tag}" patch)
else
  debug "Existing release branch found: next version will be read from it"
  git_run "${foundation_sdk_path}" fetch origin "${release_branch}"
  git_run "${foundation_sdk_path}" checkout "${release_branch}"
  next_tag=$(cat "${release_file_marker}")
  git_run "${foundation_sdk_path}" checkout -

  debug "Deleting the local release branch"
  git_run "${foundation_sdk_path}" branch -D "${release_branch}"
fi

debug "Next version: ${next_tag}"

info "Creating new release branch"
git_run "${foundation_sdk_path}" checkout -b "${release_branch}"

release_dir=$(dirname "$release_file_marker")
if [ ! -d "${release_dir}" ]; then
  mkdir "${release_dir}"
fi

echo "${next_tag}" > "${release_file_marker}"
debug "Next version: ${next_tag}"

info "Running cog"
run_codegen "output_dir=${codegen_output_path}/%l,kind_registry_path=${KIND_REGISTRY_PATH},release_tag=${next_tag}"

debug "Copying generated content to grafana-foundation-sdk"
find "${codegen_output_path}" -maxdepth 1 -mindepth 1 -print | while read -r target; do
  target=${target#"$codegen_output_path/"}

  # By removing the language folder before copying the generated output, we make
  # sure that files that might have been generated by a previous release but
  # aren't in the current workspace are pruned.
  rm -rf "${foundation_sdk_path:?}/${target}"
  cp -R "${codegen_output_path}/${target}" "${foundation_sdk_path}"
done

if [ "${SKIP_VALIDATION}" == "yes" ]; then
  warning "Skipping release validation"
else
  info "Validating release"
  "$__dir/release-validate.sh" "${foundation_sdk_path}"
fi

debug "Adding changes to git staging area"
git_run "${foundation_sdk_path}" add .

has_changes=$(git_has_changes "${foundation_sdk_path}")
if [ "${has_changes}" != "0" ]; then
  warning "No changes detected: aborting."
  exit 0
fi

git_run "${foundation_sdk_path}" commit -m "Prepare ${next_tag} release"

info "Pushing release branch ${release_branch}"
run_when_safe git_run "${foundation_sdk_path}" push origin "+${release_branch}"

debug "Ensuring that ${FOUNDATION_SDK_REPO} will be used by gh"
run_when_safe gh_run "${foundation_sdk_path}" repo set-default "${FOUNDATION_SDK_REPO}"

if [ "$release_branch_exists" != "0" ]; then
  info "Opening release Pull Request"
  run_when_safe gh_run "${foundation_sdk_path}" pr create \
    --base main \
    --head "${release_branch}" \
    --title "Next release" \
    --body "Next release."
fi

if [ "${DRY_RUN}" != "no" ]; then
  notice "Review the changes on the ${release_branch} branch in ${foundation_sdk_path} and re-run this script with DRY_RUN=no to disable dry-run mode."
  notice "Tip: git diff main..${release_branch}"
fi
