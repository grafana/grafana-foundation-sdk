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
source "${__dir}/../libs/logs.sh"

build_dir="${__dir}/../../build"
docs_build_dir="${build_dir}/docs"
mkdocs_dir="${__dir}/../../.mkdocs"

rm -rf "${build_dir}"
mkdir -p "${build_dir}"
mkdir -p "${docs_build_dir}"

${__dir}/_pull_versions.sh "${build_dir}"

echo "🪧 Building main website"
mkdocs build -f ${mkdocs_dir}/mkdocs.yml -d ${docs_build_dir}

echo '[]' > "${docs_build_dir}/versions.json"

for short_version in ${ALL_GRAFANA_VERSIONS//;/ } ; do
    full_version="${short_version}+cog-${COG_VERSION}"

    log_group_start "Building documentation for version ${full_version}"
    echo "🪧 Building documentation for version ${full_version}"

    cat <<< $(jq ". += [{\"version\": \"${full_version}\", \"title\": \"${short_version}\"}]" "${docs_build_dir}/versions.json") > "${docs_build_dir}/versions.json"

    SOURCE_VERSION_FOLDER="${build_dir}/versions/${full_version}" mkdocs build -f ${mkdocs_dir}/mkdocs-version.yml -d ${docs_build_dir}/${full_version}

    echo "🪧 Minifying HTML"
    minhtml --do-not-minify-doctype --ensure-spec-compliant-unquoted-attribute-values --keep-closing-tags --keep-input-type-text-attr --keep-html-and-head-opening-tags --preserve-brace-template-syntax --keep-spaces-between-attributes ${docs_build_dir}/${full_version}/*/*/*/*.html
    minhtml --do-not-minify-doctype --ensure-spec-compliant-unquoted-attribute-values --keep-closing-tags --keep-input-type-text-attr --keep-html-and-head-opening-tags --preserve-brace-template-syntax --keep-spaces-between-attributes ${docs_build_dir}/${full_version}/*/*/*/*/*.html

    log_group_end
done
