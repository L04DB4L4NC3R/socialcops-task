package controller

import (
	"net/http"

	"github.com/rs/cors"
)

func Startup() *http.Handler {
	m := http.NewServeMux()

	// test the server out
	m.HandleFunc("/api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API running"))
	})

	// handle CORS
	corsHandler := cors.Default().Handler(m)
	return &corsHandler
}
