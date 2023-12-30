package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://hello@world.io/1337",
		Debug:            true,
		AttachStacktrace: true,
		// Transport:        &devNullTransport{},
	})

	fmt.Println(err)

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Use(sentryecho.New(sentryecho.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))

	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(200, "ok")
	})

	app.GET("/bar", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})

	app.GET("/fo1", func(ctx echo.Context) error {
		panic("its me again!")
	})

	app.Logger.Fatal(app.Start(":3000"))
}
