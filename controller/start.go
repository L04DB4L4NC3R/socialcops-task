package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/angadsharma1016/socialcops/model"
	nats "github.com/nats-io/go-nats"
)

/**
* @api {get} /api/v1/process/start start the long running task
* @apiName start the long running task
* @apiGroup all
*
*@apiParamExample {string} request-example
*
*curl http://<domain:port>/api/v1/process/start
*
*@apiParamExample {json} response-example
*
*{
*	"taskID": 1,
*	"timestamp": "string Unix timestamp",
*	"taskName": "bigproc",
*	"isCompleted":false,
*	"wasInterrupted": false
*}
*
 */
func startTask(conn *nats.EncodedConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// proc ID
		taskID++

		// run main process goroutine
		go func(id uint) {
			log.Println("Task ID:", id)
			var pid int

			/* long running task */
			log.Println("staging long running process >>>>>>>>>>>>>>>>>>>>")

			// spawn long running task
			cmd := exec.Command("./bin/main", "bigproc", string(taskID))
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
					go model.CompleteRoutine(id)
					// commit
					go model.Commit()
				}
			}()
			/* End long running task */

			conn.Subscribe("interrupt", func(message *nats.Msg) {

				strKid := string(message.Data)
				log.Println("Received interrupt on task", strings.Replace(strKid, "\"", "", -1))
				killID, err := strconv.Atoi(strings.Replace(strKid, "\"", "", -1))
				if err != nil {
					log.Println("ERROR WHILE UNMARSHALLING")
					log.Println(err)
				}
				log.Println(killID)
				if id == uint(killID) {
					log.Println("interrupt hit >>>>>>>>>>>>>>>>>>>>")

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
					// rollback
					go model.Rollback()

					return

				} else {
					// particular goroutine not referred
					log.Println("Not referred. This routine is", id)
					log.Println("message ID", killID)
				}
			})

			// run indefinately until stopped
			select {}
		}(taskID)

		// save activation details
		rt := model.RoutineInfo{
			ID:          taskID,
			IsCompleted: false,
			Timestamp:   time.Now().String(),
			Task:        "bigproc",
		}
		if err := rt.Save(); err != nil {
			log.Println(err)
			w.Write([]byte("Error occurred while saving to SQL"))
			return
		}
		json.NewEncoder(w).Encode(rt)
	}
}
