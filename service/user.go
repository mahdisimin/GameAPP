package service

import (
	"fmt"
	"gameapp/entity"
)

type UserService struct {
	repo UserRepository
}

type UserRepository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	RegisterUser(userReq RegisterRequest) (CreatedUser entity.User, err error)
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
	Avatar      string
}

type RegisterResponse struct {
	user     entity.User
	metadata string
}

func (u UserService) Register(userReq RegisterRequest) (RegisterResponse, error) {

	// TODO - verify phone number by verification code

	// TODO - validate phone number

	if isPhoneNumberExists, err := u.repo.IsPhoneNumberExist(userReq.PhoneNumber); isPhoneNumberExists || err != nil {
		if isPhoneNumberExists {
			return RegisterResponse{}, fmt.Errorf("phoneNumber %s already exists", userReq.PhoneNumber)
		}
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("check phone number uniqueness failed: %w", err)

		}
	}

	createdUser, err := u.repo.RegisterUser(userReq)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("register user failed: %w", err)
	}

	responseUser := RegisterResponse{
		user:     createdUser,
		metadata: "",
	}
	return responseUser, nil
}
