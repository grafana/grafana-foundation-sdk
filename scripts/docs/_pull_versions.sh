#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch the error in case mysqldump fails (but gzip succeeds) in `mysqldump |gzip`
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/../versions.sh"

tmp_dir=${1:-"./"}
last_tag=${2:-""}

FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'https://github.com/grafana/grafana-foundation-sdk.git'}
SUPPORTED_LANGUAGES=("go" "php" "java" "python" "typescript")

# Cleanup potential leftovers from previous runs.
rm -rf "${tmp_dir}/tmp-foundation-sdk"
rm -rf "${tmp_dir}/versions"

mkdir -p "${tmp_dir}/versions"
git clone "${FOUNDATION_SDK_REPO}" "${tmp_dir}/tmp-foundation-sdk"

if [ -n "$last_tag" ]; then
  echo "🪧 Pulling documentation for Foundation SDK version ${last_tag} (using current checkout)"

  # Use the current checkout directory (which contains freshly generated files from CI)
  repo_root="${__dir}/../.."

  for lang in "${SUPPORTED_LANGUAGES[@]}"; do
      if [ -d "${repo_root}/${lang}/docs" ]; then
          echo "  Copying ${lang}/docs to versions/${last_tag}/${lang}"
          mkdir -p "${tmp_dir}/versions/${last_tag}/${lang}"
          cp -R "${repo_root}/${lang}"/docs/* "${tmp_dir}/versions/${last_tag}/${lang}"
      fi
  done
fi

for version in ${ALL_GRAFANA_VERSIONS//;/ } ; do
    full_version="${version}+cog-${COG_VERSION}"

    echo "🪧 Pulling documentation for Foundation SDK version ${full_version}"

    git -C "${tmp_dir}/tmp-foundation-sdk" fetch origin "${full_version}"
    git -C "${tmp_dir}/tmp-foundation-sdk" checkout "${full_version}"
    git -C "${tmp_dir}/tmp-foundation-sdk" pull --ff-only origin "${full_version}"

    find "${tmp_dir}/tmp-foundation-sdk" -maxdepth 1 -mindepth 1 -type d -print | while read -r target; do
        language_name=$(basename "$target")
        if [ -d "${target}/docs" ]; then
            mkdir -p "${tmp_dir}/versions/${full_version}/${language_name}"
            cp -R "${target}"/docs/* "${tmp_dir}/versions/${full_version}/${language_name}"
        fi
    done
done
