package model

import "log"

func (rt *RoutineInfo) Save() error {
	log.Println(rt)
	insert, err := con.Query("INSERT INTO ROUTINES VALUES (?, ?, ?, ?)", rt.ID, rt.Timestamp, rt.Task, rt.IsCompleted)

	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return nil
}

func (rt *RoutineInfo) Read(c chan RoutineInfoReturn) {
	rows, err := con.Query("SELECT id, timestamp, task_name, iscompleted FROM ROUTINES")

	var rts []RoutineInfo
	for rows.Next() {
		var rt RoutineInfo
		if err = rows.Scan(&rt.ID, &rt.Timestamp, &rt.Task, &rt.IsCompleted); err != nil {
			c <- RoutineInfoReturn{nil, err}
			return
		}
		rts = append(rts, rt)

	}

	defer rows.Close()
	c <- RoutineInfoReturn{rts, nil}
	return
}
