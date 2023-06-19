package usecase

import (
	"ge-rest-api/src/model"
	"ge-rest-api/src/repository"
	"ge-rest-api/src/utils"
	"ge-rest-api/src/validator"
)

type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUseCase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUseCase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUseCase {
	return &userUseCase{ur, uv}
}

func (uu *userUseCase) SignUp(user model.User) (model.UserResponse, error) {

	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	hashed, err := utils.HashPassword(user.Password)

	if err != nil {
		return model.UserResponse{}, err
	}

	newUser := model.User{Email: user.Email, Password: string(hashed)}

	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	res := model.UserResponse{Email: newUser.Email, ID: newUser.ID}

	return res, nil
}

func (uu *userUseCase) Login(user model.User) (string, error) {

	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := model.User{}

	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(storedUser.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
