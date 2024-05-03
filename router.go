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
	// e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{ // If we need Simple Auth
	// 	KeyLookup: "query:api-key",
	// 	Validator: func(key string, c echo.Context) (bool, error) {
	// 		return key == "KEY", nil
	// 	},
	// }))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	e.GET("/iin_check/:iin", IinVerification)
	e.POST("/people/info", AddCitizen)
	e.GET("/people/info/iin/:iin", GetCitizenByIin)

	// In requirements written this "/people/info/phone" but I think here should be "/people/info" or "/people/info/name/{}"
	// 3.2. Получение ранее сохраненных данных о человеке по части имени:
	// "/people/info/phone/{часть_имени_или_фамилии}" - метод GET, ответ на запрос в виде JSON (массив структур данных):

	// because we need to get citizen by Name not by Phone and using param for that is not a good idea
	// I think we need to use query params for that
	// e.GET("/people/info/phone/:name", GetCitizenByIin)
	e.GET("/people/info", GetCitizens)

	return e
}
