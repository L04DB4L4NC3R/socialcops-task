package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/angadsharma1016/socialcops/model"
)

/**
* @api {get} /api/v1/process/view view all the long running tasks
* @apiName view the long running task
* @apiGroup all
*@apiParamExample {string} request-example
*
*curl http://<domain:port>/api/v1/process/view
*
*@apiParamExample {json} response-example
*

*[
*	{
*		"taskID": 1,
*		"timestamp": "string Unix timestamp",
*		"taskName": "bigproc",
*		"isCompleted":false,
*		"wasInterrupted": false
*	},
*	{
*		"taskID": 2,
*		"timestamp": "string Unix timestamp",
*		"taskName": "bigproc",
*		"isCompleted":false,
*		"wasInterrupted": false
*	},
*	{	"taskID": 3,
*		"timestamp": "string Unix timestamp",
*		"taskName": "bigproc",
*		"isCompleted":false,
*		"wasInterrupted": false
*	}
*]
 */
func viewTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := make(chan model.RoutineInfoReturn)
		go model.ReadRoutineInfo(c)
		msg := <-c

		if msg.Err != nil {
			log.Println(msg.Err)
			w.Write([]byte("Error occurred while reading tasks"))
			return
		}

		json.NewEncoder(w).Encode(msg.RoutineInfo)
		return
	}
}
