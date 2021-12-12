package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Debug:            true,
		AttachStacktrace: true,
		Environment:      "production",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
	defer sentry.Recover()

	func1()
}
