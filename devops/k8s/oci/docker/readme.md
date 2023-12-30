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
