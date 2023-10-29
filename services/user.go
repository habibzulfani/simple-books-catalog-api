package services

import (
	"Mini-project/database"
	"Mini-project/middlewares"
	"Mini-project/models"
	"Mini-project/utils"
	"fmt"
)

type UserService struct {
	jwtAuth *middlewares.JWTConfig
}

func InitUserService(jwtAuth *middlewares.JWTConfig) UserService {
	return UserService{
		jwtAuth: jwtAuth,
	}
}

func (service *UserService) Register(input models.RegisterInput) (models.User, error) {
	config := &utils.ArgonConfig{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	password, err := utils.CreatePassword(input.Password, config)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
	}

	result := database.DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	err = result.Last(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (service *UserService) Login(input models.LoginInput) (string, error) {
	var user models.User

	err := database.DB.First(&user, "email = ?", input.Email).Error

	if err != nil {
		fmt.Printf("Error decoding hash: %v\n", err)
		return "", err
	}

	fmt.Printf("Input Password: %s\n", input.Password)
	fmt.Printf("User Password: %s\n", user.Password)

	match, err := utils.ComparePassword(input.Password, user.Password)

	isFailed := err != nil || !match

	if isFailed {
		return "", err
	}

	token, err := service.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UserService) GetUser(id string) (models.User, error) {
	var user models.User

	err := database.DB.First(&user, "id = ?", id).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
