ARG BUILDPLATFORM=linux/amd64

# Prepare Builder Platrform
FROM --platform=${BUILDPLATFORM} golang:1.26.1-alpine AS builder

ENV CGO_ENABLED=0 
ENV GOARCH=amd64 
ENV GOOS=linux
ENV GOPATH=/root/go

RUN apk add --no-cache curl git bash make tzdata ca-certificates
COPY --from=golangci/golangci-lint:v2.11-alpine /usr/bin/golangci-lint /usr/bin
RUN go install github.com/go-delve/delve/cmd/dlv@v1.26.1

# Build Process
FROM builder AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY  . /build
# Build & Lint & Test
RUN go build -ldflags "-w -s" -o /main
RUN go test ./...
RUN golangci-lint run ./...

# RUN apk update && \
#     apk add upx && \
#     upx --brute /main


# Final scratch image
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo/
COPY --from=build /main /
COPY --from=build /root/go/bin/dlv /bin/

ENV APP_PORT=8090
EXPOSE $APP_PORT

WORKDIR /
# CMD [ "/main" ]
CMD [ "/bin/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/main", "--continue"]