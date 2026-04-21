package Handler

import (
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("ok")); err != nil {
		log.Println(err)
	}
}
