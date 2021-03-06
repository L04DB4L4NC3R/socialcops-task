package bigproc

import (
	"database/sql"
	"log"
)

func savetoSQL(con *sql.DB, dataArr []Data) {
	tx, err := con.Begin()
	if err != nil {
		log.Println("ERROR creating transaction")
	}
	for _, data := range dataArr {
		res, err := tx.Exec("INSERT INTO BIGTASK VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.Field1, data.Field2, data.Field3, data.Field4, data.Field5, data.Field6, data.Field7, data.Field8, data.Field9, data.Field10, data.Field11, data.Field12, data.Field13, data.Field14, data.Field15, data.Field16, data.Field17, data.Field18, data.Field19, data.Field20, data.Field21, data.Field22, data.Field23, data.Field24, data.Field25)
		if err != nil {
			log.Println(err)
			if err = tx.Rollback(); err != nil {
				log.Println("ERROR while rolling back")
				log.Println(err)
			}
		}
		log.Println("Saved an object to DB")
		log.Println(res.RowsAffected())
	}
	if err = tx.Commit(); err != nil {
		log.Println("Error committing transaction")
		log.Println(err)
	}
}
