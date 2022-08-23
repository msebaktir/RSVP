package models

type Customer struct {
	ID       int
	UserName string
	Password string
	Status   int
	Event    Event
}
