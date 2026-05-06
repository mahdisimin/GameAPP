package file

import (
	"encoding/json"
	"fmt"
	"gameapp/entity"
	"os"
	"strings"
)

const (
	baseRoute = "C:\\simin\\Projects\\gameapp"
)

type File struct {
}

func (f File) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (f File) IsPhoneNumberExist(phoneNumber string) (bool, error) {
	isExists := false

	allUsers, err := f.GetAllUsers()
	if err != nil {
		return false, fmt.Errorf("failed on getAllUsers: %v", err.Error())
	}
	for _, user := range allUsers {
		if user.PhoneNumber == phoneNumber {
			isExists = true
			return isExists, nil
		}
	}
	return isExists, nil
}

func (f File) RegisterUser(user entity.User) (entity.User, error) {
	userId := uint8(f.getLastID())
	user.ID = userId
	dataByte := make([]byte, 1024)
	userJson, _ := json.Marshal(user)
	dataByte = append(userJson, '\n')
	fmt.Println(string(userJson))
	file, err := os.OpenFile(baseRoute+"\\Users.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return entity.User{}, err
	}
	if _, err := file.Write(dataByte); err != nil {
		return entity.User{}, fmt.Errorf("error on write to file: %w", err)
	}
	if err := file.Close(); err != nil {
		return entity.User{}, fmt.Errorf("error on close file: %w", err)
	}
	return user, nil

}

func (f File) getLastID() int {
	dataSlice, err := f.GetAllUsers()
	if err != nil {
		return -1
	}
	userCount := len(dataSlice)
	lastID := userCount
	return lastID
}

func (f File) GetAllUsers() ([]entity.User, error) {
	var userSlice []entity.User
	dataByte := make([]byte, 1024)
	file := &os.File{}
	if fileTemp, err := os.OpenFile(baseRoute+"\\Users.txt", os.O_RDONLY|os.O_CREATE|os.O_RDWR, 0777); err != nil {
		return []entity.User{}, err
	} else {
		file = fileTemp
	}
	dataLength, _ := file.Read(dataByte)
	dataByte = dataByte[:dataLength]
	userString := string(dataByte)
	userString = strings.Trim(userString, "\r\n")
	userStringSlice := strings.Split(userString, "\n")
	for _, userString := range userStringSlice {
		user := entity.User{}
		json.Unmarshal([]byte(userString), &user)
		userSlice = append(userSlice, user)
	}
	//dataString := string(dataByte)
	//userStringSlice := strings.Split(dataString, "\n")
	file.Close()
	return userSlice, nil
}
