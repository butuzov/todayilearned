# --- Required ---------------------------------------------------------------------------
export PATH   := $(PWD)/bin:$(PATH)                    # ./bin to $PATH
export SHELL  := bash                                  # Default Shell


GOPKGS := $(shell go list ./... | grep -vE "(cmd|testdata)" | tr -s '\n' ',' | sed 's/.\{1\}$$//')


define install_go_bin
	@ which $(1) 2>&1 1>/dev/null || GOBIN=$(PWD)/bin go install $(2)
endef


.PHONY: menu
menu:
	@ ./.github/scripts/readme_to_index
	@ ./.github/scripts/update_config_with_menu

.PHONY: static
static: menu
	hugo --logLevel debug
	sleep 1 # cooldown...
	@ # remove readme files from the artifacts
	./.github/scripts/postprocessor_build_remove_readme
	./.github/scripts/postprocessor_build_markdown_links

.PHONY: serve
serve:
	python3 -m http.server 1313 --directory ./build

live: menu
	hugo server

# ~~~ Local Development ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
.PHONY: pre-commit-changes
pre-commit-changes: ## Apply Changes and Run Initial test for `pre-commit`
	pre-commit install
	pre-commit run --all-files
