#!/usr/bin/env bash

function git_run() {
  local repo_dir="${1}"
  shift

  git -C "${repo_dir}" "$@"
}

function git_has_changes() {
  local repo_dir=${1}
  shift

  local changes
  changes=$(git_run "${repo_dir}" status --porcelain=v1)
  if [ "${changes}" != "" ]; then
    echo "0"
    return 0
  fi

  echo "1"
}
