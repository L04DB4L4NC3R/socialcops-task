package controller

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/angadsharma1016/socialcops/model"
	nats "github.com/nats-io/go-nats"
)

func testNATS(con *nats.EncodedConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Query().Get("task") == "" {
			w.Write([]byte("Usage: /api/v1/process?task=<taskName>"))
			return
		}

		taskID++
		log.Println(taskID)
		sendc := make(chan model.Routine)
		recv := make(chan model.Routine)

		// bind channels to NATS events
		con.BindSendChan("foo", sendc)
		con.BindRecvQueueChan("foo", "queue", recv)

		// send initiation message
		sendc <- model.Routine{taskID, true, false}

		// used to cancel task
		go func() {
			time.Sleep(time.Second * 5)
			sendc <- model.Routine{taskID, false, false}
		}()

		// turn this into a goroutine, task to be cancelled
		go func(id uint) {
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
						cmd := exec.Command("./bin/" + r.URL.Query().Get("task"))
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
								sendc <- model.Routine{id, false, true}
							}
						}()
						// ACK for completed JOB sendc <- model.Routine{id, false}

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

		}(taskID)

	}
}
