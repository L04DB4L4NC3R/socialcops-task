package controller

import (
	"net/http"

	nats "github.com/nats-io/go-nats"
	"github.com/rs/cors"
)

func Startup(conn *nats.EncodedConn) *http.Handler {
	m := http.NewServeMux()

	// test the server out
	m.HandleFunc("/api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API running"))
	})

	// test NATS
	m.HandleFunc("/api/v1/nats", testNATS(conn))

	// handle CORS
	corsHandler := cors.Default().Handler(m)
	return &corsHandler
}
