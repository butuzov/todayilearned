
## Dockerizing Go Apps

Note on dockerizing apps.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Len %d\n", len(body))
}
```

```Dockerfile
# builder go
FROM golang:1.12.7-alpine3.10 as BUILD_GO
WORKDIR /app
# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor
COPY ./ ./
RUN go build -o main . 

# builder
FROM alpine:3.10 as BUILD
RUN apk update && apk add ca-certificates tzdata

# Final scratch image
FROM scratch

COPY --from=BUILD /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILD /usr/share/zoneinfo /usr/share/zoneinfo/

COPY --from=BUILD_GO /app/main /
CMD [ "/main" ]
```
