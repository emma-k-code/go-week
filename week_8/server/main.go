package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type UserData struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Note   string `json:"note"`
}

type AdminData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user/get", userData)
	e.GET("/admin/get", adminData)

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}

func userData(c echo.Context) error {
	data := UserData{
		UserID: "111",
		Name:   "Test User",
		Note:   "測試用會員",
	}
	return c.JSON(http.StatusOK, data)
}

func adminData(c echo.Context) error {
	data := AdminData{
		ID:   "333",
		Name: "Admin",
		Note: "測試用管理員",
	}
	return c.JSON(http.StatusOK, data)
}
