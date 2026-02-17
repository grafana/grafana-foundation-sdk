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
source "${__dir}/libs/logs.sh"

workspace="${__dir}/../workspace"
mkdocs_dir="${__dir}/../mkdocs"
build_dir="${workspace}/build"
output_dir="${workspace}/out"

rm -rf "${workspace}"
mkdir -p "${build_dir}"
mkdir -p "${output_dir}"

echo "🪧 Copying docs base"
cp -R ${mkdocs_dir}/* "${build_dir}"

echo "🪧 Copying language docs"
find "./" -maxdepth 1 -mindepth 1 -type d -not -path "./.mkdocs" -not -path "./build" -print | while read -r target; do
    if [ -d "${target}/docs" ]; then
        mkdir -p "${build_dir}/docs/${target}"
        cp -R ${target}/docs/* "${build_dir}/docs/${target}"
    fi
done

mkdocs build -f ${build_dir}/mkdocs.yml -d ${output_dir}

echo "🪧 Minifying HTML"
minhtml --do-not-minify-doctype --ensure-spec-compliant-unquoted-attribute-values --keep-closing-tags --keep-input-type-text-attr --keep-html-and-head-opening-tags --preserve-brace-template-syntax --keep-spaces-between-attributes ${output_dir}/*/*/*/*.html
minhtml --do-not-minify-doctype --ensure-spec-compliant-unquoted-attribute-values --keep-closing-tags --keep-input-type-text-attr --keep-html-and-head-opening-tags --preserve-brace-template-syntax --keep-spaces-between-attributes ${output_dir}/*/*/*/*/*.html
