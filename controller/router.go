package controller

import (
	"net/http"

	"github.com/angadsharma1016/socialcops/model"
	nats "github.com/nats-io/go-nats"
	"github.com/rs/cors"
)

var taskID uint

func Startup(conn *nats.EncodedConn, sendc *chan model.Routine, recv *chan model.Routine) *http.Handler {
	m := http.NewServeMux()

	// test service
	m.HandleFunc("/api/v1/test", testNATS(conn))
	m.HandleFunc("/api/v1/process", startTask(sendc, recv))
	m.HandleFunc("/api/v1/process/kill", killTask(sendc))

	// handle CORS
	corsHandler := cors.Default().Handler(m)
	return &corsHandler
}
