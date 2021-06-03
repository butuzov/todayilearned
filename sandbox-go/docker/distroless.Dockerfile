FROM golang:1.16-buster as BUILD

WORKDIR /build
COPY  . /build

RUN go mod download
RUN go test ./... -v

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
    go build \
       -ldflags "-w -s" \
       -o /main

RUN apt-get update  \
    && apt-get install upx-ucl -y \
    && upx --brute /main

# We woundn't use scratch iamge, but distroless base debian.
FROM gcr.io/distroless/static:nonroot
COPY --from=BUILD /main /

ENV APP_PORT 8090
EXPOSE $APP_PORT

WORKDIR /
CMD [ "/main" ]


