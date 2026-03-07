package service

import (
	"fmt"
	"gameapp/entity"
)

type UserService struct {
	Repo UserRepository
}

type UserRepository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	RegisterUser(entity.User) (entity.User, error)
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
	Avatar      string
}

type RegisterResponse struct {
	User     entity.User
	metadata string
}

func (u UserService) Register(userReq RegisterRequest) (RegisterResponse, error) {
	// TODO - verify phone number by verification code

	// TODO - validate phone number

	if isPhoneNumberExists, err := u.Repo.IsPhoneNumberExist(userReq.PhoneNumber); isPhoneNumberExists || err != nil {
		if isPhoneNumberExists {
			return RegisterResponse{}, fmt.Errorf("phoneNumber %s already exists", userReq.PhoneNumber)
		} else if err != nil {
			return RegisterResponse{}, fmt.Errorf("check phone number uniqueness failed: %w", err)

		}
	}
	user := entity.User{
		ID:          0,
		PhoneNumber: userReq.PhoneNumber,
		Avatar:      userReq.Avatar,
		Name:        userReq.Name,
	}

	createdUser, err := u.Repo.RegisterUser(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("register user failed: %w", err)
	}

	responseUser := RegisterResponse{
		User:     createdUser,
		metadata: "",
	}
	return responseUser, nil
}
