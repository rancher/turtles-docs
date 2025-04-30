# Define variables for reuse
TMP_DIR = tmp
BUILD_DIR = build
LOG_DIR = $(TMP_DIR)
SITE_DIR = $(BUILD_DIR)/site

# Playbooks
GH_PAGES_PLAYBOOK = turtles-gh-pages-playbook.yml
REMOTE_PLAYBOOK = playbook-remote.yml
DEV_PLAYBOOK = turtles-dev-playbook.yml

# Logs
GH_PAGES_LOG = $(LOG_DIR)/gh-pages-build.log
REMOTE_LOG = $(LOG_DIR)/remote-build.log
DEV_LOG = $(LOG_DIR)/dev-build.log

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: build
build: ## Build the site using a specified playbook. Usage: make build PLAYBOOK=<playbook> LOG=<log_file>
	@if [ -z "$(PLAYBOOK)" ] || [ -z "$(LOG)" ]; then \
		echo "Error: PLAYBOOK and LOG variables must be set."; \
		exit 1; \
	fi
	mkdir -p $(TMP_DIR)
	npx antora --version
	npx antora --stacktrace --log-format=pretty --log-level=info \
		$(PLAYBOOK) \
		2>&1 | tee $(LOG)

.PHONY: gh-pages
gh-pages: environment ## Build the site locally using the GitHub Pages playbook.
	$(MAKE) build PLAYBOOK=$(GH_PAGES_PLAYBOOK) LOG=$(GH_PAGES_LOG)

.PHONY: remote
remote: environment ## Build the site remotely using the remote playbook.
	$(MAKE) build PLAYBOOK=$(REMOTE_PLAYBOOK) LOG=$(REMOTE_LOG)

##@ Cleanup

.PHONY: clean
clean: ## Remove the build and temporary directories.
	rm -rf $(BUILD_DIR) $(TMP_DIR)

##@ Environment

.PHONY: environment
environment: ## Install and update npm dependencies.
	npm install && npm update

##@ Preview

.PHONY: preview
preview: ## Preview the site locally with http-server.
	npx http-server $(SITE_DIR) -c-1

##@ Development

.PHONY: dev
dev: environment ## Build the site using the development playbook.
	$(MAKE) build PLAYBOOK=$(DEV_PLAYBOOK) LOG=$(DEV_LOG)

.PHONY: watch
watch: environment ## Watch for changes, rebuild, and preview with hot reload.
	npm install -g nodemon
	npm install -g concurrently
	concurrently \
		"nodemon --watch content --watch docs --ext adoc,yml --exec 'make dev'" \
		"make preview"

##@ Verification

.PHONY: verify-broken-links
verify-broken-links: ## Verify broken GitHub links in .adoc files.
	go run tools/verifybrokenlinks/main.go -docs-dir=docs -max-parallel=10

##@ CI

.PHONY: ci
ci: environment gh-pages dev verify-broken-links ## Run the build and verification for continuous integration.
