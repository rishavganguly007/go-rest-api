package models

import "time"

type Event struct {
	Id       int
	Name     string 
	Desc     string 
	Location string 
	DateTime time.Time
	UserId int
}

var events = []Event{}

func (e Event) Save() {
	// later add it to db
	events = append(events,e)
}

func GetAllEvents() []Event {
	return events
}
