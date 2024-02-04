<!-- menu: Docker -->
# Docker

```shell
# Image
docker image list
docker image ls -a # Also show untagged images
docker image list --digests
docker image list full_image_name --digests
docker image list --filter danling=true  # dangling
docker image list --filter before=ubuntu # pulled before ubuntu image
docker image list --filter since=ubuntu  # pulled since ubuntu image
# ImageS
docker images --format "{{.Repository}}:{{.Tag}}"
docker images --format "table {{.ID}}\t{{.Repository}}:{{.Tag}}\t{{.Size}}"
docker images --digests

# exceing into
docker exec 7e20c58dcd17 ls /
# in the shell
docker run -it name sh

# format (but this not work same everywhere)
docker system info --format 'Default logging driver: {{.LoggingDriver}}'
docker ps --format '{{json .}}'
docker ps --format '{{json .}}' | jq
docker ps --format '{{ .ID}} {{ .Names}} {{ .Networks}} {{ .Ports}}'
docker ps --format 'table {{ .ID}}\t | \t{{ .Names}}\t | \t{{ .Networks}}\t | \t{{ .Ports}}'
docker inspect --format '{{join .Config.Cmd " , "}}'
docker inspect --format '{{json .config}}' sdk2i
# go templates
docker inspect --format='{{range .NetworkSettings.Networks}}{{println .IPAddress }}{{println .IPPrefixLen}}{{ .MacAddress}}{{end}}'
```

## Usefull

### See How Long a Container Ran For

```shell
> docker container inspect  82d07bd09e7 | grep dAt
            "StartedAt": "2024-01-12T03:57:21.246442922Z",
            "FinishedAt": "0001-01-01T00:00:00Z",
> docker container inspect  82d07bd09e7 | jq '.[].State.StartedAt,.[].State.FinishedAt'
```

### Healthchecks

#### HEALTHCHECK

https://docs.docker.com/engine/reference/builder/#healthcheck

#### Docker-Compose: `curl`

Such curl command will generate more readable logs and reduce its memory fingerprint.

```shell
docker container inspect a76b884d1b2a | jq '.[].State.Health.Log[] | [.Output,.Start,.End] '
docker container inspect a76b884d1b2a | jq '.[].State.Health.Log[] | {Output,Start,End} '
```
```yaml
x-healthchecks: &healthchecks
  start_period: 20s
  interval: 10s
  timeout: 1s
  retries: 60

services:
  service_name:
    ...
    healthcheck:
      <<: *healthchecks
      test: curl -s --fail localhost:port > /dev/null
```

## Tutorials

* [ENTRYPOINT vs CMD]()
* [Faster Multi-Platform Builds](https://www.docker.com/blog/faster-multi-platform-builds-dockerfile-cross-compilation-guide/)

## KnowHow

### Distroless Images

- https://github.com/GoogleContainerTools/distroless
- [debugging distroless images](https://edu.chainguard.dev/chainguard/chainguard-images/debugging-distroless-images/)

## Creating Images With

### Docker

```shell
# Builds
docker build -t custom-nginx:v1.
docker build -t foo/bar .

docker buildx build -t docker.io/foo/bar:latest .
```

### `github actions`

https://github.com/docker/build-push-action

```yaml
name: Verify and Push

on: [push, pull_request]

jobs:
  verify:
      - name: Log in to Docker Hub
        uses: docker/login-action@v3
        with:
          username: butuzov
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v2

      - name: Set up container image build
        uses: docker/setup-buildx-action@v3

      - name: Build and push container image
        uses: docker/build-push-action@v5
        with:
          push: true
          tags: butuzov/image:latest
          platforms: linux/amd64,linux/arm64

      #  # or Do it via goreleaser.
      # - uses: goreleaser/goreleaser-action@v5
      #   with:
      #     args: release --rm-dist
      #   env:
      #     GITHUB_TOKEN: "${{ secrets.GITHUB_TOKEN }}"
```
