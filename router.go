package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{ // If we need Auth
	// 	KeyLookup: "query:api-key",
	// 	Validator: func(key string, c echo.Context) (bool, error) {
	// 		return key == "KEY", nil
	// 	},
	// }))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))
	e.Use()

	e.GET("/iin_check/:iin", IinVerification)

	return e
}
