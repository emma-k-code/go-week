package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	Kernel "./kernel"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/weather", getWeather)

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}

func getWeather(c echo.Context) error {
	city := c.QueryParam("city")

	data := Kernel.GetWeather(city)

	return c.JSON(http.StatusOK, data)
}
