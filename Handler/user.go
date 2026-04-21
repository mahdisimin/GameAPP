package Handler

import (
	"encoding/json"
	"gameapp/repository/file"
	"gameapp/service"
	"io"
	"log"
	"net/http"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userReq service.UserRegisterRequest
	var fileRepo file.File
	var data = make([]byte, r.ContentLength)
	userService := service.NewUserService(fileRepo)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if dataRaedAll, err := io.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data = dataRaedAll
	}

	if err := json.Unmarshal(data, &userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if _, err := w.Write([]byte(`{"Message" : "New User Register"}`)); err != nil {
		log.Fatal("UserRegisterHandler err:", err)
	}

	if _, err := userService.Register(userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}
