package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	customer "praktikum/controller"
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

	e.GET("/customers", customer.GetCustomer)
	e.POST("/customers", customer.CreateCustomer)
	e.GET("/customers/:id", customer.GetCustomerById)
	e.PUT("/newcustomer/:id", customer.UpdateCustomerData)
	e.DELETE("/deletecustomer/:id", customer.DeleteCustomer)

	e.Logger.Fatal(e.Start("127.0.0.1:5000"))
}
