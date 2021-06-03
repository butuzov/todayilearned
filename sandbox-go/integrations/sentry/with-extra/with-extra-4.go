package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

type devNullTransport struct{}

func (t *devNullTransport) Configure(options sentry.ClientOptions) {
	dsn, _ := sentry.NewDsn(options.Dsn)
	fmt.Println()
	fmt.Println("Store Endpoint:", dsn.StoreAPIURL())
	fmt.Println("Headers:", dsn.RequestHeaders())
	fmt.Println()
}

func (t *devNullTransport) SendEvent(event *sentry.Event) {
	fmt.Println("Faked Transport")
}

func (t *devNullTransport) Flush(timeout time.Duration) bool {
	return true
}

type CustomComplexError struct {
	Message  string
	MoreData map[string]string
}

func (e CustomComplexError) Error() string {
	return "CustomComplexError: " + e.Message
}

func (e CustomComplexError) GimmeMoreData() map[string]string {
	return e.MoreData
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Debug: true,
		Dsn:   "https://hello@world.io/1337",

		Transport: &devNullTransport{},
	}); err != nil {
		panic(err)
	}

	sentry.WithScope(func(scope *sentry.Scope) {
		err := CustomComplexError{
			Message: "say what again. SAY WHAT again",
			MoreData: map[string]string{
				"say": "wat1",
			},
		}

		for key, val := range err.GimmeMoreData() {
			scope.SetExtra(key, val)
		}

		sentry.CaptureException(err)
	})
}
