package model

import "time"

type Comment struct {
	Id       int
	IdPage   int
	Username string
	Comment  string
	Time     time.Time
}
