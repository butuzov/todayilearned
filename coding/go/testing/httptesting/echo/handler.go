package testecho

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, verifyTokenWithKey)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}

			return token, nil
		},
	}))

	g.POST("", func(ctx echo.Context) error {
		user := ctx.Get("user").(*jwt.Token)

		if claims, ok := user.Claims.(jwt.MapClaims); ok {
			return ctx.String(http.StatusOK, fmt.Sprintf("Welcome, %s\n", claims["Name"]))
		}

		return ctx.String(http.StatusOK, fmt.Sprintf("Welcome"))
	})

	return e
}
