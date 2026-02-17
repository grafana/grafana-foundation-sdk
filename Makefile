COG_VERSION = v0.0.56
COG_DIR     = $(shell go env GOPATH)/bin/cog-$(COG_VERSION)
COG_BIN     = $(COG_DIR)/cli

KIND_REGISTRY_PATH="./workspace/kind-registry"

.PHONY: install-cog
install-cog: echodir $(COG_BIN)

.PHONY: echodir

$(COG_BIN):
	@echo "Installing Cog version $(COG_VERSION)"
	@mkdir -p $(COG_DIR)
	GOBIN=$(COG_DIR) go install github.com/grafana/cog/cmd/cli@$(COG_VERSION)
	@touch $@

.PHONY: clone-kind-registry
clone-kind-registry:
	KIND_REGISTRY_DIR=$(KIND_REGISTRY_PATH) ./scripts/fetch-kind-registry.sh

.PHONY: generate
generate: install-cog clone-kind-registry
	$(COG_BIN) generate --config cog/config.yaml \
		--parameters "output_dir=%l,kind_registry_path=$(KIND_REGISTRY_PATH),release_tag=$(shell cat .release/tag)"

.PHONY: preview
preview: install-cog
	COG_CMD=$(COG_BIN) SKIP_VALIDATION='yes' CLEANUP_WORKSPACE='no' ./scripts/prepare-release.sh
