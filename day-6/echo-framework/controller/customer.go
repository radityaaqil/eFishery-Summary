package customer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var allCustomers = []Customer{
	{
		Id:    1,
		Name:  "huda",
		Email: "huda@gmail.com",
	},
}

var LastId = allCustomers[0].Id + 1

// GET
func GetCustomer(c echo.Context) error {
	log.Println("Get All Customers Request")
	return c.JSON(http.StatusOK, allCustomers)
}

// POST
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

	allCustomers = append(allCustomers, user)
	LastId++

	log.Println("New Customer Entry")

	return c.JSON(http.StatusCreated, user)
}

// GET BY ID
func GetCustomerById(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	log.Println("Get Customer By Id")

	return c.JSON(http.StatusFound, allCustomers[idInt-1])
}

// PATCH
func UpdateCustomerData(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	req := new(Customer)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	newData := Customer{
		Id:    allCustomers[idInt-1].Id,
		Name:  req.Name,
		Email: req.Email,
	}

	allCustomers[idInt-1] = newData

	log.Println("Update Customer Data Request")

	return c.JSON(http.StatusAccepted, allCustomers[idInt-1])
}

// DELETE
func DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	var designatedIdx int

	for i := 0; i < len(allCustomers); i++ {
		if idInt == allCustomers[i].Id {
			designatedIdx = i
			break
		}
	}

	allCustomers = append(allCustomers[:designatedIdx], allCustomers[designatedIdx+1:]...)

	log.Println("Delete Customer Data Request")

	return c.JSON(http.StatusAccepted, allCustomers)
}
