<!-- menu: Distributing  -->

# Distributing

## Pakcages

### pypi.org

This is a part of [deadlinks](http://github.com/butuzov/deadlinks/) deploy makefile

```Makefile
# ~~~ Deployments ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

build-prod: venv clean ## Build disto (source & wheel) - Production
	@ $(PYTHON) setup.py sdist bdist_wheel > /dev/null 2>&1

build-dev: venv clean ## Build disto (source & wheel) - Development
	@ \
	DEADLINKS_BRANCH=$(BRANCH) \
	DEADLINKS_COMMIT=$(COMMIT) \
	DEADLINKS_TAGGED=$(TAGGED) \
	$(PYTHON) setup.py sdist bdist_wheel > /dev/null 2>&1

clean: ## Cleanup Build artifacts
	@echo "Cleanup Temporary Files"
	@rm -rf ${DIST}
	@rm -rf ${BUILD}
	@rm -f deadlinks/__develop__.py

deploy:
	@ $(PYTHON) -m pip install --upgrade wheel twine -q

deploy-test: deploy build-dev ## PyPi Deploy (test.pypi.org)
	twine upload --repository-url https://test.pypi.org/legacy/ dist/*;\

deploy-prod: deploy build-prod ## PyPi Deploy (pypi.org)
	twine upload --repository-url https://upload.pypi.org/legacy/ dist/*;\

pre-depeloy-check: venv clean deps ## Install Development Version
	$(PYTHON) -m pip uninstall deadlinks -y
		DEADLINKS_BRANCH=$(BRANCH) \
		DEADLINKS_COMMIT=$(COMMIT) \
		DEADLINKS_TAGGED=$(TAGGED) \
		$(PYTHON) setup.py develop -q 2>&1 1> /dev/null
```

### PIP

```
pip install --help

Usage:
  pip install [options] <requirement specifier> [package-index-options] ...
  pip install [options] -r <requirements file> [package-index-options] ...
  pip install [options] [-e] <vcs project url> ...
  pip install [options] [-e] <local project path> ...
  pip install [options] <archive url/path> ...
```

### Jupyter

You always can distribute Python code as Jupyter files to run the in Jupyter/Jupyterlab.

### Poetry

TODO: Write docs

```shell
poetry new poetry-demo
poetry build
poetry publish
```


## Containers

### `Dockerfile`

Most common way to build container image is to use Docker command (You also can build containers using `Bazel`, `podman` etc...)

{{% list "Dockerfile,hello.py,.dockerignore,requirements.txt" %}}
