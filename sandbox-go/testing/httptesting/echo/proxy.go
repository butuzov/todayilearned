package testecho

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// This is simple proxy middleware
// Example:
//    g := e.Group("/foo")
//  	g.Use(proxyRequests(barUrl))
func proxyRequests(base string) echo.MiddlewareFunc {
	upstream, err := url.Parse(base)
	if err != nil {
		panic(err)
	}

	proxy := middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRandomBalancer([]*middleware.ProxyTarget{
			{
				Name: "direct-debit-settings",
				URL:  upstream,
			},
		}),
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			req.Host = upstream.Host
			c.SetRequest(req)

			return proxy(next)(c)
		}
	}
}
