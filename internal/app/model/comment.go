package model

import (
	"errors"
	"time"
)

type Comment struct {
	Id       int
	IdPage   int
	Username string
	Comment  string
	Time     time.Time
}

func (c *Comment) CheckUsername() {
	if len(c.Username) <= 0 {
		c.Username = "Анон-сама"
	}
}

func (c *Comment) CheckMessage() error {
	if len(c.Comment) <= 0 {
		return errors.New("Comment too short")
	}
	return nil
}
