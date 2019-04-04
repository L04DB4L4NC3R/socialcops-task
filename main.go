package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/angadsharma1016/socialcops/controller"
	nats "github.com/nats-io/go-nats"
)

type server struct {
	Host    string
	Port    string
	NatsCon *nats.EncodedConn
}

func (s *server) Init(addr string) {
	str := strings.Split(addr, ":")
	s.Host = str[0]
	s.Port = str[1]
	log.Println("Server ready at", addr)
}

func (s server) RunProc() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalln(err)
	}
	defer nec.Close()
	s.NatsCon = nec

	// make routine channels
	sendc := make(chan controller.Routine)
	recv := make(chan controller.Routine)

	// bind channels to NATS events
	nec.BindSendChan("process", sendc)
	nec.BindRecvQueueChan("process", "queue", recv)

	handler := controller.Startup(nec, &sendc, &recv)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), *handler))
}

func main() {
	var s server
	s.Init("0.0.0.0:3000")
	s.RunProc()
}
