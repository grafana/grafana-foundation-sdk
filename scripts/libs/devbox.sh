#!/usr/bin/env bash

function devbox_run() {
  local language="${1}"
  shift
  local pwd="${1}"
  shift

  devbox run -q -c ".config/$language" -- "cd $pwd ; $@"
}
