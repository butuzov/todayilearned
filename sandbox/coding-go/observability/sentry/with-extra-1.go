package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

func prettyPrint(v interface{}) string {
	pp, _ := json.MarshalIndent(v, "", "  ")
	return string(pp)
}

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

		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			// Solution 1 (use beforeSend, which will be applied to
			// all events and is usually application specific):
			if ex, ok := hint.OriginalException.(CustomComplexError); ok {
				for key, val := range ex.GimmeMoreData() {
					event.Extra[key] = val
				}
			}

			fmt.Printf("%s\n\n", prettyPrint(event.Extra))

			return event
		},
		Transport: &devNullTransport{},
	}); err != nil {
		panic(err)
	}

	sentry.CaptureException(CustomComplexError{
		Message: "say what again. SAY WHAT again",
		MoreData: map[string]string{
			"say": "wat1",
		},
	})
}
