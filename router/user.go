package router

import (
	"gameapp/Handler"
	"log"
	"net/http"
)

func ResloveRouter() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/register", Handler.UserRegisterHandler)
	mux.HandleFunc("/HealthCheck", Handler.HealthCheckHandler)

	log.Println("Server is Listening on port 8088")
	if err := http.ListenAndServe(":8088", mux); err != nil {
		return err
	}
	return nil
}
