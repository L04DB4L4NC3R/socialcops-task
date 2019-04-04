package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/angadsharma1016/socialcops/model"
)

func killTask(sendc *chan model.Routine) http.HandlerFunc {
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

		// send activation signal
		*sendc <- model.Routine{uint(id), false, false}

		w.Write([]byte("Sent kill signal to task" + killID))

	}
}
