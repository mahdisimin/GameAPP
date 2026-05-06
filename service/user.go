package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gameapp/entity"
)

type UserService struct {
	Repo UserRepository
}

type UserRepository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	RegisterUser(entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type UserRegisterRequest struct {
	Name        string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password"`
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

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type UserLoginResponse struct {
	UserID int `json:"user_id"`
}

func (u UserService) Register(userReq UserRegisterRequest) (UserRegisterResponse, error) {
	// TODO - verify phone number by verification code

	// TODO - validate phone number

	if len(userReq.Password) < 8 {
		return UserRegisterResponse{}, errors.New("password must be at least 8 characters")
	}

	if len(userReq.Name) < 3 {
		return UserRegisterResponse{}, fmt.Errorf("name %s is too short", userReq.Name)
	}

	if isPhoneNumberExists, err := u.Repo.IsPhoneNumberExist(userReq.PhoneNumber); isPhoneNumberExists || err != nil {
		if isPhoneNumberExists {
			return UserRegisterResponse{}, fmt.Errorf("phoneNumber %s already exists", userReq.PhoneNumber)
		}
		return UserRegisterResponse{}, fmt.Errorf("check phone number uniqueness failed: %w", err)

	}

	hashPassword := hashTextFunc(userReq.Password)

	user := entity.User{
		ID:          0,
		PhoneNumber: userReq.PhoneNumber,
		Avatar:      userReq.Avatar,
		Name:        userReq.Name,
		Password:    hashPassword,
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

func (u UserService) Login(userReq UserLoginRequest) (UserLoginResponse, error) {
	var user entity.User

	isExists, err := u.Repo.IsPhoneNumberExist(userReq.PhoneNumber)
	if err != nil || !isExists {
		if err != nil {
			return UserLoginResponse{}, fmt.Errorf("check phone number failed: %s", err.Error())
		}
		return UserLoginResponse{}, errors.New("phone number does not exist")
	}
	hashPassword := hashTextFunc(userReq.Password)
	if usertemp, err := u.Repo.GetUserByPhoneNumber(userReq.PhoneNumber); err != nil {
		return UserLoginResponse{}, fmt.Errorf("check phone number failed: %s", err.Error())
	} else {
		user = usertemp
	}
	if user.Password != hashPassword {
		return UserLoginResponse{}, errors.New("password does not match")
	}

	return UserLoginResponse{
		int(user.ID),
	}, nil
}

func hashTextFunc(text string) string {
	plainPass := text
	hashPass := md5.Sum([]byte(plainPass))
	hashText := hex.EncodeToString(hashPass[:])
	return hashText
}
