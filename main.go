package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/angadsharma1016/socialcops/bigproc"
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
	handler := controller.Startup(nec)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), *handler))
}

func main() {
	db := model.ConnectDB()
	defer db.Close()

	// if flagged to bigproc, execute accordingly else run normally
	if len(os.Args) > 1 {
		if os.Args[1] == "bigproc" {
			bigproc.Do(db)
		} else {
			fmt.Println("usage: ./bin/main OR ./bin/main bigproc")
		}
		return
	}
	var s server
	s.Init("0.0.0.0:3000")
	s.RunProc()
}
