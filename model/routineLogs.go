package model

import "log"

func (rt *RoutineInfo) Save() error {
	log.Println(rt)
	insert, err := con.Query("INSERT INTO ROUTINES VALUES (?, ?, ?, ?, ?)", rt.ID, rt.Timestamp, rt.Task, rt.IsCompleted, rt.WasInterrupted)

	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return nil
}

func ReadRoutineInfo(c chan RoutineInfoReturn) {
	rows, err := con.Query("SELECT id, timestamp, task_name, iscompleted, wasinterrupted FROM ROUTINES")

	var rts []RoutineInfo
	for rows.Next() {
		var rt RoutineInfo
		if err = rows.Scan(&rt.ID, &rt.Timestamp, &rt.Task, &rt.IsCompleted, &rt.WasInterrupted); err != nil {
			c <- RoutineInfoReturn{nil, err}
			return
		}
		rts = append(rts, rt)

	}

	defer rows.Close()
	c <- RoutineInfoReturn{rts, nil}
	return
}

func CompleteRoutine(taskID uint) {
	row, err := con.Query("UPDATE ROUTINES SET iscompleted = true WHERE id = ?", taskID)
	defer row.Close()
	// if there is an error inserting, handle it
	if err != nil {
		log.Println(err)
		return
	}
	// be careful deferring Queries if you are using transactions
	return
}

func InterruptRoutine(taskID uint) {
	row, err := con.Query("UPDATE ROUTINES SET wasinterrupted = true WHERE id = ?", taskID)
	defer row.Close()
	// if there is an error inserting, handle it
	if err != nil {
		log.Println(err)
		return
	}
	// be careful deferring Queries if you are using transactions
	return
}

//Killable Is the process killable
func Killable(taskID uint) (killable bool, err error) {
	var completed, interrupted bool

	row := con.QueryRow("SELECT iscompleted, wasinterrupted FROM ROUTINES WHERE id = ?", taskID)

	if err = row.Scan(&completed, &interrupted); err != nil {
		return false, err
	}

	return !completed && !interrupted, nil
}
