FROM golang:1.16-alpine as BUILD


WORKDIR /build
COPY  . /build

RUN go mod download
# RUN GODEBUG=gocachehash=1 go test ./... -v
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
    go build \
       -ldflags "-w -s" \
       -o /main

RUN apk update && \
    apk add upx && \
    upx --brute /main


FROM golang:1.15-alpine as EXTRA
RUN apk update && apk add ca-certificates tzdata


# Final scratch image
FROM scratch

COPY --from=EXTRA /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=EXTRA /usr/share/zoneinfo /usr/share/zoneinfo/
COPY --from=BUILD /main /

ENV APP_PORT 8090
EXPOSE $APP_PORT

WORKDIR /
CMD [ "/main" ]
