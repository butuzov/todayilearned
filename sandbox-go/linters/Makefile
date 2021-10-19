
.DEFAULT_GOAL := help

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

define github_url
    https://github.com/$(GITHUB)/releases/download/v$(VERSION)/$(ARCHIVE)
endef


# ~~ [ golangci-lint ] ~~~ https://github.com/golangci/golangci-lint ~~~~~~~~~~~~~~~~~~~~~

VERSION ?= 1.42.1

GOLANGCI := $(shell command -v golangci-lint || echo "bin/golangci-lint")
golangci-lint: bin/golangci-lint ## Installs golangci-lint (linter)

install: bin/golangci-lint

bin/golangci-lint: GITHUB  := golangci/golangci-lint
bin/golangci-lint: ARCHIVE := golangci-lint-$(VERSION)-$(OSTYPE)-amd64.tar.gz
bin/golangci-lint: bin
	@ printf "Install golangci-linter $(VERSION)... "
	@ curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - $(shell printf golangci-lint-$(VERSION)-$(OSTYPE)-amd64/golangci-lint | tr A-Z a-z ) > $@-$(VERSION) && chmod +x $@-$(VERSION)
	@ echo "done."


# ~~~ Updates ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

update-gocritic: ## Go Critic
	@./scripts/updates_for_gocritic

update-gocritic...: ##
	@./scripts/updates_for_gocritic -verbose

update-linters: ## Supproted Linters
	@./scripts/update_linters_list

update-linters...: ##
	@./scripts/update_linters_list -verbose

config: ## Updates Configs
	@./scripts/update_default_config

	@ echo "runnig yq merge"
	@ ./scripts/exclude_yaml > .golangci.ongoing.excludes.yaml
	@ [[ -f .golangci.ongoing.yaml ]] && unlink .golangci.ongoing.yaml || echo "no .golangci.ongoing.yaml"
	@ cp .golangci.default.yaml .golangci.ongoing.yaml
	@ yq ea -i 'select(fileIndex == 0) * select(fileIndex == 1)' .golangci.ongoing.yaml .golangci.ongoing.excludes.yaml
	@ yq  ea -i 'select(fileIndex == 0) * select(fileIndex == 1)' .golangci.ongoing.yaml .golangci.ongoing.fixes.yaml
	@ yq  ea -i 'del(.linters-settings.errcheck.exclude)' .golangci.ongoing.yaml
	@ yq  ea -i 'del(.linters-settings.custom.example)' .golangci.ongoing.yaml

	# gocritic noise
	@ yq  ea -i 'del(.linters-settings.gocritic.enabled-checks[0])' .golangci.ongoing.yaml
	@ yq  ea -i '.linters-settings.gocritic.enabled-checks += "ruleguard"' .golangci.ongoing.yaml
	@ yq  ea -i '.linters-settings.gocritic.enabled-checks += "nestingreduce"' .golangci.ongoing.yaml
	@ yq  ea -i '.linters-settings.gocritic.enabled-checks += "truncatecmp"' .golangci.ongoing.yaml
	@ yq  ea -i '.linters-settings.gocritic.enabled-checks += "unnamedresult"' .golangci.ongoing.yaml


	# echo "useless"
	@ yq ea -i 'del(.issues.new-from-rev)' .golangci.ongoing.yaml
	@ yq ea -i 'del(.issues.new-from-patch)' .golangci.ongoing.yaml
	# echo "incompattible rules"
	@ yq ea -i 'del(.linters)' .golangci.ongoing.yaml
	# exhaisive struct settings.
	@ yq ea -i 'del(.linters-settings.exhaustivestruct.struct-patterns)' .golangci.ongoing.yaml


# ~~~ Checks ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# GOLANGCI_BIN := golangci-lint-$(VERSION)
GOLANGCI_BIN ?= golangci-lint


edit:
	code ~/.dotfiles/.golangci.yaml

check-gocritic: ## Go-Critic
	$(GOLANGCI_BIN) run -v -c .golangci.gocritic.yaml testdata-gocritic

check-revive: ## revive
	$(GOLANGCI_BIN) run -c .golangci.revive.yaml testdata-revive


CONFIG ?= .golangci.ongoing.yaml
check: clean ##
	$(GOLANGCI_BIN) version
	$(GOLANGCI_BIN) run -c $(CONFIG) -v testdata --enable-all --sort-results --issues-exit-code 0

linter: clean
	$(GOLANGCI_BIN) version
	$(GOLANGCI_BIN) run -c $(CONFIG) -v testdata/$(linter).go --disable-all -E $(linter)


test: clean ## Runs tests for ongoing examples
	CONFIG=$(PWD)/.golangci.ongoing.yaml $(MAKE) check > check.log
	@ ./scripts/test_ongoing

test...: clean ## Runs tests for ongoing examples
	CONFIG=$(PWD)/.golangci.ongoing.yaml $(MAKE) check > check.log
	@ ./scripts/test_ongoing -verbose

mine: clean ## Runs tests for ongoing examples
	CONFIG=~/.dotfiles/.golangci.yaml $(MAKE) check > check.log
	@ ./scripts/test_ongoing

mine...: clean ## Runs tests for ongoing examples
	CONFIG=~/.dotfiles/.golangci.yaml $(MAKE) check > check.log
	@  ./scripts/test_ongoing -verbose

clean:
	golangci-lint cache clean

dep-gawk:
	@ if [ -z "$(shell command -v gawk)" ]; then  \
		if [ -x /usr/local/bin/brew ]; then $(MAKE) _brew_gawk_install; exit 0; fi; \
		if [ -x /usr/bin/apt-get ]; then $(MAKE) _ubuntu_gawk_install; exit 0; fi; \
		if [ -x /usr/bin/yum ]; then  $(MAKE) _centos_gawk_install; exit 0; fi; \
		if [ -x /sbin/apk ]; then  $(MAKE) _alpine_gawk_install; exit 0; fi; \
		echo  "GNU Awk Required.";\
		exit 1; \
	fi

_brew_gawk_install:
	@ echo "Instaling gawk using brew... "
	@ brew install gawk --quiet
	@ echo "done"

_ubuntu_gawk_install:
	@ echo "Instaling gawk using apt-get... "
	@ apt-get -q install gawk -y
	@ echo "done"

_alpine_gawk_install:
	@ echo "Instaling gawk using yum... "
	@ apk add --update --no-cache gawk
	@ echo "done"

_centos_gawk_install:
	@ echo "Instaling gawk using yum... "
	@ yum install -q -y gawk;
	@ echo "done"


help: dep-gawk
	@echo "==============================================================================="
	@echo "=  Help Page: github.com/butuzov/mira                                         ="
	@echo "==============================================================================="
	@cat $(MAKEFILE_LIST) | \
		grep -E '^# ~~~ .*? [~]+$$|^[a-zA-Z0-9_-]+:.*?## .*$$' | \
		gawk '{if ( $$1=="#" ) { \
			match($$0, /^# ~~~ (.+?) [~]+$$/, a);\
			{print "\n", a[1], ""}\
		} else { \
			match($$0, /^([a-zA-Z/_-]+):.*?## (.*)$$/, a); \
			{printf "  - \033[32m%-20s\033[0m %s\n",   a[1], a[2]} \
 		}}'
	@echo ""

