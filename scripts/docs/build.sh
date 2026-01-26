#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch the error in case mysqldump fails (but gzip succeeds) in `mysqldump |gzip`
set -o pipefail

RELEASE_TAG=${1:-}

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/../versions.sh"
source "${__dir}/../libs/logs.sh"

build_dir="${__dir}/../../build"
docs_build_dir="${build_dir}/docs"
mkdocs_dir="${__dir}/../../.mkdocs"

rm -rf "${build_dir}"
mkdir -p "${build_dir}"
mkdir -p "${docs_build_dir}"

"${__dir}"/_pull_versions.sh "${build_dir}" "${RELEASE_TAG}"

echo "🪧 Building main website"
mkdocs build -f "${mkdocs_dir}"/mkdocs.yml -d "${docs_build_dir}"

echo '[]' > "${docs_build_dir}/versions.json"

if [ -n "${RELEASE_TAG}" ]; then
  build_documentation "${RELEASE_TAG}" "latest"
fi

for short_version in ${ALL_GRAFANA_VERSIONS//;/ } ; do
    full_version="${short_version}+cog-${COG_VERSION}"
    build_documentation "$full_version" "$short_version"
done

build_documentation() {
      log_group_start "Building documentation for version $1"
      echo "🪧 Building documentation for version $1"
  
      cat <<< $(jq ". += [{\"version\": \"$1\", \"title\": \"$2\"}]" "${docs_build_dir}/versions.json") > "${docs_build_dir}/versions.json"
  
      SOURCE_VERSION_FOLDER="${build_dir}/versions/$1" mkdocs build -f "${mkdocs_dir}"/mkdocs-version.yml -d "${docs_build_dir}"/"$1"
  
      echo "🪧 Minifying HTML"
      minhtml --do-not-minify-doctype --ensure-spec-compliant-unquoted-attribute-values --keep-closing-tags --keep-input-type-text-attr --keep-html-and-head-opening-tags --preserve-brace-template-syntax --keep-spaces-between-attributes "${docs_build_dir}"/"$1"/*/*/*/*.html
      minhtml --do-not-minify-doctype --ensure-spec-compliant-unquoted-attribute-values --keep-closing-tags --keep-input-type-text-attr --keep-html-and-head-opening-tags --preserve-brace-template-syntax --keep-spaces-between-attributes "${docs_build_dir}"/"$1"/*/*/*/*/*.html
  
      log_group_end
}
