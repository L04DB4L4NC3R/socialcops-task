package controller

import (
	"log"
	"net/http"
	"time"

	nats "github.com/nats-io/go-nats"
)

var count uint

type Routine struct {
	ID        uint
	ShouldRun bool
}

func testNATS(con *nats.EncodedConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count++
		log.Println(count)
		sendc := make(chan Routine)
		recv := make(chan Routine)

		// bind channels to NATS events
		con.BindSendChan("foo", sendc)
		con.BindRecvQueueChan("foo", "queue", recv)

		// send initiation message
		sendc <- Routine{count, true}

		// used to cancel task
		go func() {
			time.Sleep(time.Second * 5)
			sendc <- Routine{count, false}
		}()

		// turn this into a goroutine, task to be cancelled
		go func(id uint) {

			for {
				msg := <-recv
				if msg.ID == id {
					if msg.ShouldRun {

						log.Println("Long running task")

					} else {
						log.Println("Exiting")
						return
					}
				} else {
					log.Println("Not referred")
				}

			}

		}(count)

	}
}
