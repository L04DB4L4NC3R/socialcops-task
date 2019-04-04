package controller

import (
	"net/http"

	nats "github.com/nats-io/go-nats"
	"github.com/rs/cors"
)

var taskID uint

func Startup(conn *nats.EncodedConn) *http.Handler {
	m := http.NewServeMux()

	// test service
	m.HandleFunc("/api/v1/test", testNATS(conn))
	m.HandleFunc("/api/v1/process/start", startTask(conn))
	m.HandleFunc("/api/v1/process/kill", killTask(conn))
	m.HandleFunc("/api/v1/process/view", viewTasks())

	// handle CORS
	corsHandler := cors.Default().Handler(m)
	return &corsHandler
}
