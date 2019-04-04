package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/angadsharma1016/socialcops/model"
)

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
