package controller

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"

	nats "github.com/nats-io/go-nats"
)

var count uint

type Routine struct {
	ID          uint
	ShouldRun   bool
	IsCompleted bool
	Err         error
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
		sendc <- Routine{count, true, false, nil}

		// used to cancel task
		go func() {
			time.Sleep(time.Second * 5)
			sendc <- Routine{count, false, false, nil}
		}()

		// turn this into a goroutine, task to be cancelled
		func(id uint) {
			var pid int
			for {
				msg := <-recv

				if msg.IsCompleted {
					return
				}

				if msg.ID == id {

					// run only once
					if msg.ShouldRun {

						// spawn long running task
						cmd := exec.Command("./bin/task")
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr

						err := cmd.Start()
						if err != nil {
							log.Fatalf("cmd.Run() failed with %s\n", err)
						}

						// capture PID
						pid = cmd.Process.Pid
						log.Println("Running with pid", pid)

						go func() {
							err = cmd.Wait()
							if err != nil {
								log.Println(err)
							} else {
								sendc <- Routine{id, false, true, nil}
							}
						}()
						// ACK for completed JOB sendc <- Routine{id, false}

					} else {
						err := syscall.Kill(int(pid), syscall.SIGKILL)
						if err != nil {
							log.Println(err)
						}
						log.Println(pid)
						log.Println("Exiting", pid)
						return
					}
				} else {
					log.Println("Not referred")
				}

			}

		}(count)

	}
}
