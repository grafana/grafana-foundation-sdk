COG_VERSION = v0.0.45
COG_DIR     = $(shell go env GOPATH)/bin/cog-$(COG_VERSION)
COG_BIN     = $(COG_DIR)/cli

GRAFANA_VERSION="v12.2.1"
KIND_REGISTRY_VERSION="next" # It points to existing schemas from the main Grafana branch.
KIND_REGISTRY_PATH="../kind-registry" # You have to clone the repository if you want to do local testing.

.PHONY: install-cog
install-cog: echodir $(COG_BIN)

.PHONY: echodir

$(COG_BIN):
	@echo "Installing Cog version $(COG_VERSION)"
	@mkdir -p $(COG_DIR)
	GOBIN=$(COG_DIR) go install github.com/grafana/cog/cmd/cli@$(COG_VERSION)
	@touch $@

.PHONY: update-cog
update-app-sdk: ## Update the Grafana App SDK dependency in go.mod
	@pwd
	go get github.com/grafana/cog@$(COG_VERSION)
	go mod tidy

.PHONY: generate
generate: install-cog 
	bash -c 'source ./scripts/versions.sh && \
		cog generate --config .cog/config.yaml \
		--parameters "output_dir=%l,kind_registry_path=$(KIND_REGISTRY_PATH),kind_registry_version=$(KIND_REGISTRY_VERSION),grafana_version=$(GRAFANA_VERSION),release_branch=main,all_grafana_versions=$$ALL_GRAFANA_VERSIONS,cog_version=$$COG_VERSION"'
