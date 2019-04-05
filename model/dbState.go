package model

import "log"

func Commit() {
	res, err := con.Exec("COMMIT")
	if err != nil {
		log.Println("ERROR committing")
		log.Println(err)
	}
	log.Println(res.RowsAffected())
}

func Rollback() {
	res, err := con.Exec("ROLLBACK")
	if err != nil {
		log.Println("ERROR rolling back")
		log.Println(err)
	}
	log.Println(res.RowsAffected())
}
