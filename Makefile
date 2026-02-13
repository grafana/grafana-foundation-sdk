DEPS_BIN_DIR = bin
COMPOSER_VERSION = 2.9.5
COMPOSER_BIN = $(DEPS_BIN_DIR)/composer
COG_VERSION = v0.0.54
COG_DIR     = $(shell go env GOPATH)/bin/cog-$(COG_VERSION)
COG_BIN     = $(COG_DIR)/cli

KIND_REGISTRY_PATH="./kind-registry"

# Within devbox
ifneq "$(DEVBOX_CONFIG_DIR)" ""
    RUN_DEVBOX:=
else # Normal shell
    RUN_DEVBOX:=devbox run
endif

.PHONY: install-cog
install-cog: echodir $(COG_BIN)

.PHONY: echodir

$(COG_BIN):
	@echo "Installing Cog version $(COG_VERSION)"
	@mkdir -p $(COG_DIR)
	GOBIN=$(COG_DIR) go install github.com/grafana/cog/cmd/cli@$(COG_VERSION)
	@touch $@

$(COMPOSER_BIN):
	@mkdir -p $(DEPS_BIN_DIR)
	curl -o "$(COMPOSER_BIN)" https://getcomposer.org/download/$(COMPOSER_VERSION)/composer.phar
	@chmod u+x "$(COMPOSER_BIN)"

.PHONY: deps
deps: php-deps python-deps

.PHONY: php-deps
php-deps:
	echo "Installing PHP dependencies"
	$(RUN_DEVBOX) composer install

.PHONY: python-deps
python-deps:
	echo "Installing Python dependencies"
	$(RUN_DEVBOX) pip install -r requirements.txt

.PHONY: clone-kind-registry
clone-kind-registry:
	./scripts/fetch-kind-registry.sh

.PHONY: generate
generate: install-cog clone-kind-registry
	$(COG_BIN) generate --config .cog/config.yaml \
		--parameters "output_dir=%l,kind_registry_path=$(KIND_REGISTRY_PATH)"
