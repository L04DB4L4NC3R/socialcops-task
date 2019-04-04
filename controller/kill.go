package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/angadsharma1016/socialcops/model"
	nats "github.com/nats-io/go-nats"
)

func killTask(conn *nats.EncodedConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var killID = r.URL.Query().Get("id")

		// select which task to run
		if killID == "" {
			w.Write([]byte("Usage: /api/v1/process/kill?id=<killID>"))
			return
		}

		id, err := strconv.ParseUint(killID, 10, 64)
		if err != nil {
			log.Println(err)
		}

		// check if task if killable
		killable, err := model.Killable(uint(id))
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		} else if err == sql.ErrNoRows {
			w.Write([]byte("No such process found"))
			return
		}

		// if task is killable, proceed to send interrupt
		if killable {

			// send interrupt signal
			log.Println("Publishing interrupt signal on task ", killID)
			conn.Publish("interrupt", killID)

			w.Write([]byte("Sent kill signal to task" + killID))
			return

		} else {
			w.Write([]byte("The task has already been completed"))
		}

	}
}
