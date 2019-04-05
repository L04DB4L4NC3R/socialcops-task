package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	con *sql.DB
	DB  **sql.DB
)

func ConnectDB() *sql.DB {
	str := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", str)
	if err != nil {
		log.Println(err)
		log.Println("Error connecting to DB")
		os.Exit(1)
		return nil
	}
	log.Println("Connected to DB")
	con = db
	DB = &db
	return con
}

type Routine struct {
	ID             uint `json:"id"`
	ShouldRun      bool `json:"shouldRun"`
	IsCompleted    bool `json:"isCompleted"`
	WasInterrupted bool `json:"wasInterrupted"`
}

type RoutineInfo struct {
	ID             uint   `json:"taskID"`
	Timestamp      string `json:"timestamp"`
	Task           string `json:"taskName"`
	IsCompleted    bool   `json:"isCompleted"`
	WasInterrupted bool   `json:"wasInterrupted"`
}

type RoutineInfoReturn struct {
	RoutineInfo []RoutineInfo `json:"routineInfoLogs"`
	Err         error         `json:"err"`
}
