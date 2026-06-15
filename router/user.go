package router

import (
	"gameapp/Handler"
	"log"
	"net/http"

	muxlib "github.com/gorilla/mux"
)

func ResloveRouter() error {
	//mux := http.NewServeMux()
	mux := muxlib.NewRouter()

	mux.HandleFunc("/user/register", Handler.UserRegisterHandler)
	mux.HandleFunc("/user/login", Handler.UserLoginHandler)
	mux.HandleFunc("/HealthCheck", Handler.HealthCheckHandler)
	mux.HandleFunc("/user/getProfile", Handler.GetProfileHandler)

	log.Println("Server is Listening on port 8088")
	if err := http.ListenAndServe(":8088", mux); err != nil {
		return err
	}
	return nil
}
