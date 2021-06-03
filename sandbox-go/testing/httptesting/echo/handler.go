package testecho

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type handler struct {
	users map[string]string
	echo  http.Handler
}

func Handler(db map[string]string) http.Handler {
	e := echo.New()

	e.GET("/public", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "this is public endpoint")
	})

	e.POST("/authorize", func(ctx echo.Context) error {
		user := ctx.FormValue("user")
		pass := ctx.FormValue("pass")

		for u, p := range db {
			if user == u && pass == p {
				// create token
				token := jwt.New(jwt.SigningMethodHS256)

				// set our claims
				token.Claims = &TokenClaim{
					&jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
					},
					"root",
					u,
				}
				// generate encoded token and send it as response
				t, err := token.SignedString(secret)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}

				return ctx.JSON(http.StatusOK, map[string]string{
					"token": t,
				})
			}
		}

		return echo.ErrUnauthorized
	})

	g := e.Group("/restricted", middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &TokenClaim{},
		SigningKey: secret,
	}))

	g.POST("", func(ctx echo.Context) error {
		var (
			user   = ctx.Get("user").(*jwt.Token)
			claims = user.Claims.(*TokenClaim)
		)

		return ctx.String(http.StatusOK, fmt.Sprintf("Welcome, %s\n", claims.Name))
	})

	return e
}
