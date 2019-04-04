package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/angadsharma1016/socialcops/controller"
	"github.com/angadsharma1016/socialcops/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
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
	nc, _ := nats.Connect("nats:4222")
	nec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalln(err)
	}
	defer nec.Close()
	s.NatsCon = nec

	// make routine channels
	sendc := make(chan model.Routine)
	recv := make(chan model.Routine)

	// bind channels to NATS events
	nec.BindSendChan("process", sendc)
	nec.BindRecvQueueChan("process", "queue", recv)

	handler := controller.Startup(nec, &sendc, &recv)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), *handler))
}

func main() {
	var s server
	con := model.ConnectDB()
	defer con.Close()
	s.Init("0.0.0.0:3000")
	s.RunProc()
}
