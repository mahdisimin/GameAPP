package file

import (
	"encoding/json"
	"fmt"
	"gameapp/entity"
	"os"
	"strings"
)

func (f File) IsPhoneNumberExist(phoneNumber string) (bool, error) {
	isExists := false

	return isExists, nil
}
func (f File) RegisterUser(user entity.User) (entity.User, error) {
	userId := uint8(f.getLastID())
	user.ID = userId
	f.baseRoute = baseRoute
	dataByte := make([]byte, 1024)
	userJson, _ := json.Marshal(user)
	dataByte = append(userJson, '\n')
	fmt.Println(string(userJson))
	file, err := os.OpenFile(f.baseRoute+"\\Users.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
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
	dataSlice, err := f.getAllUsers()
	if err != nil {
		return -1
	}
	userCount := len(dataSlice)
	lastID := userCount
	return lastID
}

func (f File) getAllUsers() ([]string, error) {
	dataByte := make([]byte, 1024)
	file := &os.File{}
	f.baseRoute = baseRoute
	if fileTemp, err := os.OpenFile(f.baseRoute+"\\Users.txt", os.O_RDONLY|os.O_CREATE|os.O_RDWR, 0777); err != nil {
		return []string{}, err
	} else {
		file = fileTemp
	}
	dataLength, _ := file.Read(dataByte)
	dataByte = dataByte[:dataLength]
	dataString := string(dataByte)
	dataSlice := strings.Split(dataString, "\n")
_:
	file.Close()
	return dataSlice, nil
}
