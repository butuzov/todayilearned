FROM golang:1.17-buster as BUILD

WORKDIR /build

COPY  go.mod  .
RUN go mod download

COPY  . /build
RUN go test ./... -v

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
    go build -trimpath \
       -ldflags "-w -s" \
       -o /main

#  optionaly upx compression.
RUN apt-get update  \
    && apt-get install upx-ucl -y \
    && upx --brute /main

# We woundn't use scratch iamge, but distroless base debian.
FROM gcr.io/distroless/static:nonroot
COPY --from=BUILD /main /

WORKDIR /
CMD [ "/main" ]
