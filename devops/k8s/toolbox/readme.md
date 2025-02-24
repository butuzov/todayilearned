# Toolbox Image

## Dockerfile

```shell
docker build --progress=plain --no-cache -t=butuzov/toolbox:latest --platform=linux/amd64 .
docker push docker.io/butuzov/toolbox
```

```Dockerfile
FROM node:alpine

RUN sed -i -e 's/v[[:digit:]]\..*\//edge\//g' /etc/apk/repositories \
    && apk update \
    && apk --update add --no-cache -t .deps ca-certificates \
    && apk add \
        bash \
        bash-completion \
        curl \
        strace \
        git \
        jq \
        yq \
        openssh \
        openssl \
        procps \
        nmap \
        vim \
        tmux \
    && npm install -g wscat \
    && rm -rf  /var/cache/apk/*  \
    && apk del .deps  \
    && rm -rf  /usr/share/vim/vim81/autoload/* \
    && rm -rf  /usr/share/vim/vim81/compiler/*
```

## Usage via Docker

```shell
docker run -it --rm butuzov/toolbox-node -- wscat
docker run -it --rm --net=host butuzov/toolbox-node -- wscat

# alias
alias tb="docker run -it --rm butuzov/toolbox $1"
tb curl google.com
```

## Usage via Kubernetes

```shell
# Running Forever
kubectl run toolbox --image=butuzov/toolbox:latest --restart=Never -- tail -f /dev/null

# Executing Commands
kubectl exec toolbox -- curl http://service-name

# Alias
alias ktb="exec toolbox -- $1"
tb curl google.com

# Terminating Pod
kuebctl delete pod toolbox
```

### Tools

#### `wscat`

```shell

```
