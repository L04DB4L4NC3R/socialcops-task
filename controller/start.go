package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/angadsharma1016/socialcops/model"
)

func startTask(sendc *chan model.Routine, recv *chan model.Routine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var taskName = r.URL.Query().Get("task")

		// select which task to run
		if taskName == "" {
			w.Write([]byte("Usage: /api/v1/process?task=<taskName>"))
			return
		}

		// proc ID
		taskID++

		// send activation signal
		*sendc <- model.Routine{taskID, true, false}

		// run main process goroutine
		go func(id uint) {
			var pid int

			for {

				// wait for signal
				msg := <-*recv

				// if completed, send ACK and exit
				if msg.IsCompleted {
					log.Println("The task has been completed. Task ID:", id)
					return
				}

				if msg.ID == id {

					// run only once
					if msg.ShouldRun {

						// spawn long running task
						cmd := exec.Command("./bin/" + taskName)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr

						err := cmd.Start()
						if err != nil {
							log.Fatalf("cmd.Run() failed with %s\n", err)
						}

						// capture PID
						pid = cmd.Process.Pid
						log.Println("Running with pid", pid)

						// execute task as a fork
						go func() {
							err = cmd.Wait()
							if err != nil {
								log.Println(err)
							} else {

								// completion signal
								*sendc <- model.Routine{id, false, true}
							}
						}()

					} else {

						// on interrupt, SIGKILL syscall for the process pid
						log.Println("Interrupt on pid:", pid)
						log.Println("Interrupt on task ID:", id)
						err := syscall.Kill(int(pid), syscall.SIGKILL)
						if err != nil {
							log.Println(err)
						}
						log.Println("Killing process with pid:", pid)
						log.Println("Killing task with task ID:", id)
						return
					}
				} else {

					// particular goroutine not referred
					log.Println("Not referred")
					continue
				}

			}

		}(taskID)

		json.NewEncoder(w).Encode(model.RoutineInfo{
			ID:          taskID,
			IsCompleted: false,
			Timestamp:   time.Now(),
			Task:        taskName,
		})
	}
}
