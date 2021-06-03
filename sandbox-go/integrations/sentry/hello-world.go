package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
)

var (
	errIsOne   = errors.New("This is One Error")
	errIsTwo   = errors.New("This is Two Error")
	errIsThree = fmt.Errorf("This is Three Error: %w", errIsTwo)
	errIsFour  = errors.New("This is Four Error")
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Debug:            true,
		AttachStacktrace: true,
		Environment:      "development",
		IgnoreErrors: []string{
			errIsTwo.Error(),
		},
		ServerName: "service_name",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtra("character.name", "Mighty Fighter")
	})

	defer sentry.Flush(2 * time.Second)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		wg.Done()
	})

	sentry.CaptureException(errIsFour)

	wg.Wait()
}
