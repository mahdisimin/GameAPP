package Handler

import (
	"encoding/json"
	"fmt"
	"gameapp/pkg"
	"gameapp/repository/file"
	"gameapp/service"
	"io"
	"log"
	"net/http"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userReq service.UserRegisterRequest
	var userRes service.UserRegisterResponse
	var fileRepo file.File
	var data = make([]byte, r.ContentLength)
	userService := service.NewUserService(fileRepo)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if dataRaedAll, err := io.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	} else {
		data = dataRaedAll
	}

	if err := json.Unmarshal(data, &userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error on register user : %s", err)

		return
	}

	if uRes, err := userService.Register(userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error on register user : %s", err)

		return
	} else {
		userRes = uRes
	}

	if _, err := fmt.Fprintf(w, `{"Message":"User %s register successfully"}`, userRes.User.Name); err != nil {
		log.Fatal("Error with Writing on Writer")
	}

}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var loginReq service.UserLoginRequest
	var userRes service.UserLoginResponse
	var data = make([]byte, r.ContentLength)
	fileRepo := file.File{}
	userService := service.NewUserService(fileRepo)
	if dataTemp, err := io.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"Message":"Error on read user request : %s"}`, err.Error())
		return
	} else {
		data = dataTemp
	}
	if err := json.Unmarshal(data, &loginReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf(`{"Message":"Error on unmarshal user request : %s"}`, err.Error())
	}

	if userResTemp, err := userService.Login(loginReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error on login user : %s", err)

		return
	} else {
		userRes = userResTemp
	}
	userRespByte, marErr := json.Marshal(userRes)
	if marErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error on marshal user : %s", marErr.Error())

		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(userRespByte)
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	userServ := service.UserService{
		file.File{},
	}
	var getProfileReq service.GetProfileRequest
	var getProfileResp service.GetProfileResponse
	jwtStr := r.Header.Get("Authorization")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	claims, errJWT := pkg.ParseJWT(jwtStr)
	if errJWT != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf(`{"Message":"Error on parse JWT : %s"}`, errJWT.Error())

		return
	}
	userID := claims.UserID
	getProfileReq = service.GetProfileRequest{
		userID,
	}
	if res, err := userServ.GetProfile(getProfileReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf(`{"Message":"Error on get user profile : %s"}`, err.Error())

		return
	} else {
		getProfileResp = res
	}
	getProfileRespJsonBytes, errJ := json.Marshal(getProfileResp)
	if errJ != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf(`{"Message":"Error on marshal user profile : %s"}`, errJ.Error())

		return
	}

	w.Write(getProfileRespJsonBytes)
	//w.WriteHeader(http.StatusOK)

}
