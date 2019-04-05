package model

import "log"

func CommitDB() {
	res, err := con.Exec("COMMIT")
	if err != nil {
		log.Println("ERROR committing")
		log.Println(err)
	}
	log.Println(res.RowsAffected())
}

func RollbackDB() {
	res, err := con.Exec("ROLLBACK")
	if err != nil {
		log.Println("ERROR rolling back")
		log.Println(err)
	}
	log.Println(res.RowsAffected())
}
