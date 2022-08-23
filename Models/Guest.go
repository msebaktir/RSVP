package models

type Guest struct {
	Id              int
	Name            string
	Email           string
	Phone           string
	Status          int
	HowManyPeople   int
	HowManyChildren int
	Customer        Customer
	Event           Event
	Gift            Gift
}
