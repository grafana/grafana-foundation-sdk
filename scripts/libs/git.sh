#!/usr/bin/env bash

function git_run() {
  local repo_dir="${1}"
  shift

  git -C "${repo_dir}" "$@"
}

function git_has_branch() {
  local repo_dir=${1}
  shift
  local branch=${1}
  shift

  # test local branches
  local in_local
  in_local=$(git_run "${repo_dir}" branch --list "${branch}")
  if [ "${in_local}" != "" ]; then
    echo "0"
    return 0
  fi

  # test remote branches
  local in_remote
  in_remote=$(git_run "${repo_dir}" ls-remote --heads origin "${branch}")
  if [ "${in_remote}" != "" ]; then
    echo "0"
    return 0
  fi

  echo "1"
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

function git_create_orphan_branch() {
  local repo_dir=${1}
  shift
  local branch=${1}
  shift

  git_run "${repo_dir}" checkout --orphan "${branch}"
  git_run "${repo_dir}" rm -rf .
}
