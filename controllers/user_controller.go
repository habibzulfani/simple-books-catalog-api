package controllers

import (
	"Mini-project/middlewares"
	"Mini-project/models"
	"Mini-project/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController(jwtAuth *middlewares.JWTConfig) UserController {
	return UserController{
		service: services.InitUserService(jwtAuth),
	}
}

func (controller *UserController) Register(c echo.Context) error {
	var userInput models.RegisterInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Failed",
			Message: "Invalid Request",
		})
	}

	user, err := controller.service.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response{
		Status:  "Success",
		Message: "User Registered",
		Data:    user,
	})

}

func (controller *UserController) Login(c echo.Context) error {
	var userInput models.LoginInput

	err := c.Bind(&userInput)
	fmt.Printf("Email from request: %s\n", userInput.Email)
	fmt.Printf("Password from request: %s\n", userInput.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Failed",
			Message: "Invalid Request",
		})
	}

	token, err := controller.service.Login(userInput)

	if err != nil {
		fmt.Printf("Login Error: %v\n", err) // Log the error
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Failed",
			Message: "Invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "User Registered",
		Data:    token,
	})
}

func (controller *UserController) GetUser(c echo.Context) error {
	claims, err := middlewares.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response{
			Status:  "Failed",
			Message: "Invalid token",
		})
	}

	user, err := controller.service.GetUser((strconv.Itoa(claims.ID)))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Status:  "Failed",
			Message: "User not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "User data",
		Data:    user,
	})

}
