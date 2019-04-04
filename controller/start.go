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
			w.Write([]byte("Usage: /api/v1/process/start?task=<taskName>"))
			return
		}

		// proc ID
		taskID++

		// send activation signal
		*sendc <- model.Routine{taskID, true, false, false}

		// run main process goroutine
		go func(id uint) {
			log.Println("Task ID:", id)
			var pid int

			for {

				msg := <-*recv

				if msg.ID == id {

					// if completed, send ACK and exit
					if msg.IsCompleted {
						log.Println("The task has been completed. Task ID:", id)
						go model.CompleteRoutine(id)
						return
					}

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
								w.Write([]byte("Completed task" + string(id)))
								// completion signal
								*sendc <- model.Routine{id, false, true, false}
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
						// update the task as interrupted
						go model.InterruptRoutine(id)
						return
					}
				} else {

					// particular goroutine not referred
					log.Println("Not referred. This routine is", id)
					log.Println("message ID", msg.ID)
					continue
				}

			}

		}(taskID)

		rt := model.RoutineInfo{
			ID:          taskID,
			IsCompleted: false,
			Timestamp:   time.Now().String(),
			Task:        taskName,
		}
		if err := rt.Save(); err != nil {
			log.Println(err)
			w.Write([]byte("Error occurred while saving to SQL"))
			return
		}
		json.NewEncoder(w).Encode(rt)
	}
}
