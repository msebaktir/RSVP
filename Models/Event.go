package models

import "google.golang.org/genproto/googleapis/type/datetime"

type Event struct {
	Id          int
	Title       string
	Description string
	StartDate   datetime.DateTime
	LastCall    datetime.DateTime
	Status      int
}
