package controller

import (
	"day7/model"
	"day7/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ControllerUser struct {
	service1 service.ServiceUser
}

func NewControllerUser(service1 service.ServiceUser) *ControllerUser {
	return &ControllerUser{service1}
}

func (co *ControllerUser) GetAllUsersController(c echo.Context) error {
	users, err := co.service1.GetAllUsersService()

	if err != nil {
		log.Println("Get All Users Controller Error", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (co *ControllerUser) GetUserByIdController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println("Failed to convert parameter id")
		return c.JSON(http.StatusBadRequest, errId)
	}

	user, err := co.service1.GetUserByIdService(id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (co *ControllerUser) CreateUserController(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		log.Println("Error Creating User", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	createdUser, errCreatedUser := co.service1.CreateUserService(newUser.Name, newUser.Age)
	if errCreatedUser != nil {
		log.Println("Error Creating User", errCreatedUser)
		return c.JSON(http.StatusBadRequest, errCreatedUser)
	}

	return c.JSON(http.StatusAccepted, createdUser)
}

func (co *ControllerUser) UpdateUserController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println("Failed to convert parameter id")
		return c.JSON(http.StatusBadRequest, errId)
	}

	var updateUser model.User
	if err := c.Bind(&updateUser); err != nil {
		log.Println("Error Updating User", err)
		return c.JSON(http.StatusBadRequest, errId)
	}

	updatedUser, errUpdatedUser := co.service1.UpdateUserService(id, updateUser.Name, updateUser.Age)
	if errUpdatedUser != nil {
		log.Println("Error Updating User", errUpdatedUser)
		return c.JSON(http.StatusBadRequest, errUpdatedUser)
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (co *ControllerUser) DeleteUserController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println("Failed to convert parameter id")
		return c.JSON(http.StatusBadRequest, errId)
	}

	_, err := co.service1.DeleteUserService(id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	response := "Delete Success"

	return c.JSON(http.StatusOK, response)
}
