# `pip`

## Usage

```shell
pip install some-package-name
pip uninstall some-package-name

# intalling local

# Install a package from GitHub
pip install git+https://github.com/psf/requests.git@v2.30.0
# install a package from a distribution file
pip install sampleproject-1.0.tar.gz
# using wheels
pip install sampleproject-1.0-py2.py3-none-any.whl
#
PIP insall Some --find-links=http://webpage
# Regular Install of local project
pip install path/to/project
# Editble Install of local project
pip install -e  path/to/project
# keep quite
pip install sampleproject -q
# install / upgrade
pip install --upgrade sampleproject

# Requirements
pip freeze > requirements.txt
pip install -r requirements.txt
pip install -c constraints.txt


# Latest version
python -m pip install requests
# Specific version
pip install requests==2.5.0
# Minimum version
pip install 'requests>=2.5.0'

# Download wheels
pip download requests
```

## Using Proxy

https://pypicache.readthedocs.io/en/latest/

```shell
python -m pypicache.main /tmp/mypackages
pip install -i http://localhost:8080/simple somepackage
```


```
  GET /

  GET /simple/mypackage

  GET /local/mypackage
          Currently case sensitive

  POST /requirements.txt

  POST /uploadpackage/ - Applies simple logic to parse package name - Can’t overwrite packages

  GET /packages/source/m/mypackage/mypackage-1.0.tar.gz - Checks PyPI if not present locally

  GET /packages/2.7/m/mypackage/mypackage-1.0-py2.7.egg - not implemented
  PUT /packages/2.7/m/mypackage/mypackage-1.0-py2.7.egg - not implemented
```

## Isolated Environments with `pipx`

```
brew install pipx
```

## `~/pip/pip.conf` or `PIP_CONFIG_FILE`

[Posible Locations](https://pip.pypa.io/en/latest/topics/configuration/) or `pip.conf`

## Build Packages

See [Distribute](../../distributing/)

## Special Cases of using Python Libraries.

### `psycopg2` @ m1/m2

This isn't a case anymore, but we leave it out for a history reasons.

```
brew install libpq --build-from-source
brew install openssl@3

export PATH="/opt/homebrew/opt/libpq/bin:$PATH"
export LDFLAGS="-L/opt/homebrew/opt/libpq/lib"
export LDFLAGS="-L/opt/homebrew/opt/openssl@3/lib ${LDFLAGS}"
export CPPFLAGS="-I/usr/local/opt/openssl/include"

pip install psycopg2
```
