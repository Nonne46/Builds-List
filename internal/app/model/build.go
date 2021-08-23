package model

import "time"

type Build struct {
	Id           int
	Name         string
	Description  string
	NameAddr     string
	IsAlive      bool
	Tags         string
	AuthorRepo   string
	ByondVersion string
	Github       string
	BackupDate   time.Time
	Thanks       string
}
