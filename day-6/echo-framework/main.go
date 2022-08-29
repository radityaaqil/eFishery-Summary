package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	customer "echo-framework/controller"
)

func main() {
	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message":       "Success",
		}

		return e.JSON(http.StatusOK, result)
	})

	e.GET("", customer.GetCustomer)
	e.POST("", customer.CreateCustomer)

	e.Logger.Fatal(e.Start("127.0.0.1:5000"))
}
