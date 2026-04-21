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

type UserRegisterRequest struct {
	Name        string
	PhoneNumber string
	Avatar      string
}

type UserRegisterResponse struct {
	User     entity.User
	metadata string
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (u UserService) Register(userReq UserRegisterRequest) (UserRegisterResponse, error) {
	// TODO - verify phone number by verification code

	// TODO - validate phone number

	if isPhoneNumberExists, err := u.Repo.IsPhoneNumberExist(userReq.PhoneNumber); isPhoneNumberExists || err != nil {
		if isPhoneNumberExists {
			return UserRegisterResponse{}, fmt.Errorf("phoneNumber %s already exists", userReq.PhoneNumber)
		}
		return UserRegisterResponse{}, fmt.Errorf("check phone number uniqueness failed: %w", err)

	}

	user := entity.User{
		ID:          0,
		PhoneNumber: userReq.PhoneNumber,
		Avatar:      userReq.Avatar,
		Name:        userReq.Name,
	}

	createdUser, err := u.Repo.RegisterUser(user)
	if err != nil {
		return UserRegisterResponse{}, fmt.Errorf("register user failed: %w", err)
	}

	responseUser := UserRegisterResponse{
		User:     createdUser,
		metadata: "",
	}
	return responseUser, nil
}
