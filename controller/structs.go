package controller

import "time"

var taskID uint

type Routine struct {
	ID          uint
	ShouldRun   bool
	IsCompleted bool
}

type RoutineInfo struct {
	ID          uint      `json:"taskID"`
	Timestamp   time.Time `json:"timestamp"`
	Task        string    `json:"taskName"`
	IsCompleted bool      `json:"isCompleted"`
}
