package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/angadsharma1016/socialcops/controller"
)

type server struct {
	Host string
	Port string
}

func (s *server) Init(addr string) {
	str := strings.Split(addr, ":")
	s.Host = str[0]
	s.Port = str[1]
	log.Println("Server ready at", addr)
}

func (s server) Serve(h *http.Handler) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), *h))
}

func main() {
	var s server
	s.Init("0.0.0.0:3000")
	s.Serve(controller.Startup())
}
