package customer

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var LastId = 1
var CustomerList = make(map[int]Customer)

func GetCustomer(c echo.Context) error {
	var result []Customer

	dataDummyCustomer := Customer{
		Id:    1,
		Name:  "huda",
		Email: "huda@gmail.com",
	}

	result = append(result, dataDummyCustomer)
	for key, _ := range CustomerList {
		tempCustomer := CustomerList[key]

		result = append(result, tempCustomer)
	}

	return c.JSON(http.StatusOK, result)
}

func CreateCustomer(c echo.Context) error {
	req := new(Customer)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := Customer{
		Id:    LastId,
		Name:  req.Name,
		Email: req.Email,
	}

	CustomerList[LastId] = user
	LastId++

	return c.JSON(http.StatusCreated, user)
}
