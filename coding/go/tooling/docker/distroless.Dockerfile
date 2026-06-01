ARG BUILDPLATFORM=linux/amd64

# --- Builder ------------------------------------------------------------------
FROM --platform=${BUILDPLATFORM} golang:1.26.1-bookworm AS builder

ENV CGO_ENABLED=0 
ENV GOARCH=amd64 
ENV GOOS=linux
ENV GOPATH=/root/go

RUN go install github.com/go-delve/delve/cmd/dlv@v1.26.1
COPY --from=golangci/golangci-lint:v2.11 /usr/bin/golangci-lint /usr/bin

# --- Build --------------------------------------------------------------------
FROM builder AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY  . /build

# --- Code Quality -- Build & Lint & Test --------------------------------------
RUN go test ./...
RUN golangci-lint run ./...
RUN go build -o /main

# --- Packing it ---------------------------------------------------------------
# RUN echo "deb http://deb.debian.org/debian bookworm-backports main" > /etc/apt/sources.list.d/backports.list && \
#     apt-get update && apt-get install -y -t bookworm-backports upx-ucl
# RUN upx --brute /main


# --- Prod Ready Image ---------------------------------------------------------
FROM gcr.io/distroless/static:nonroot
COPY --from=build /main /
COPY --from=build /root/go/bin/dlv /bin/

ENV APP_PORT=8090
EXPOSE $APP_PORT

WORKDIR /
CMD [ "/main" ]
CMD [ "/bin/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/main", "--continue"]


