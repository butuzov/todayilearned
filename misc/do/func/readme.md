# Functions

Serverless functions on demand from [digitalocean](https://www.digitalocean.com/products/functions)

## Reading

- https://gist.github.com/5ran6/1a77247b6f24c6c831ca35c4a3fbe487

## Project Structure

Typical Project Structure for Go (+extras)

```
./lib                                 # Library Follder
./lib/go.mod                          # (extra go.mod)
./packages/project/function           # `fucntion` from `project`
./packages/project/function/go.mod    # (extra go.mod)
./packages/                           # pacakges root (mandatory)
./packages/go.mod                     # (extra go.mod)
./project.yml                         # project.yml
./go.mod                              # (extra go.mod)
./go.work                             # (extra go.mod)
```

```yaml
parameters: {}
environment:
  PORT: ${PORT}

packages:
  - name: foobar
    environment: {}
    parameters: {}
    actions:
      - name: accounts
        runtime: "go:default"
        environment: {}
        parameters: {}
        parameters: {}
  - name: barbuz
    environment: {}
    parameters: {}
    actions:
      - name: kyc
        runtime: "go:default"
        environment: {}
        parameters: {}
```

## Commands

```shell
# Example fo CI dry run for deployment.
export CI_NAME="lobbyx-ci-$(tr -dc A-Za-z0-9 < /dev/urandom | head -c 13; echo)"
doctl sls install
doctl sls namespaces create --label $CI_NAME --region=ams
doctl sls connect $CI_NAME
doctl sls status
doctl sls deploy .
doctl sls namespaces delete $CI_NAME -f
```

```shell
doctl sls watch . # run in continues watch and deploy mode
```