#!/usr/bin/env bash

set -euo pipefail

KIND_REGISTRY_DIR="./kind-registry"
REPO_URL="https://github.com/grafana/kind-registry.git"

if [ -d $KIND_REGISTRY_DIR/.git ]; then
  echo "Updating kind-registry"
  git -C "$KIND_REGISTRY_DIR" fetch --all --quiet
  git -C "$KIND_REGISTRY_DIR" reset --hard origin/main --quiet
else
  echo "Cloning kind-registry..."
  git clone --depth=1 "$REPO_URL" "$KIND_REGISTRY_DIR"
fi

