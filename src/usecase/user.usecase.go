package usecase

import (
	"ge-rest-api/src/model"
	"ge-rest-api/src/repository"
	"ge-rest-api/src/utils"
)

type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUseCase struct {
	ur repository.IUserRepository
}

func NewUserUseCase(ur repository.IUserRepository) IUserUseCase {
	return &userUseCase{ur}
}

func (uu *userUseCase) SignUp(user model.User) (model.UserResponse, error) {
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
