package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/register", userRegister)
	if err := http.ListenAndServe("localhost:8084", mux); err != nil {
		fmt.Printf(err.Error())
	}

}

func userRegister(writer http.ResponseWriter, req *http.Request) {
	fmt.Printf("User Register...\n")
	if req.Method == http.MethodGet {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Ridi Abam qate!!")

	}
}
