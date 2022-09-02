package router

import (
	"day9/controller"
	"day9/repository"
	"day9/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *echo.Echo {
	//Initialize Echo
	e := echo.New()

	//User
	UserRepository := repository.NewRepositoryUser(db)
	UserService := service.NewUserService(*UserRepository)
	UserController := controller.NewControllerUser(*UserService)

	e.GET("/users", UserController.GetAllUsersController)
	e.GET("/user/:id", UserController.GetUserByIdController)
	e.POST("/newuser", UserController.CreateUserController)
	e.PUT("/updateuser/:id", UserController.UpdateUserController)
	e.DELETE("/deleteuser/:id", UserController.DeleteUserController)

	return e
}
