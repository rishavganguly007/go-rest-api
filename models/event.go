package models

import (
	"time"

	"example.com/go-rest-api/db"
)

type Event struct {
	Id       int64 `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Location string `json:"location"`
	DateTime time.Time `json:"date"`
	UserId int64	`json:"user_id"`
}

var events = []Event{}

func (e Event) Save() error {
	// later add it to db
	query := `
	INSERT INTO events(id, name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Desc, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Desc, &event.DateTime, &event.Location, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var event Event
	err = row.Scan(&event.Id, &event.Name, &event.Desc, &event.DateTime, &event.Location, &event.UserId)
	if err != nil {
		return nil, err
	}
	
	return &event, nil
}

func (event Event) Update(id int64) (error) {
	query := `
	UPDATE events
	SET name = ?, desc = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(event.Name, event.Desc, event.Location, event.DateTime)
	if err != nil {
		return err
	}
	
	return nil
}

func (event Event) Delete(id int64) (error) {
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Id)
	if err != nil {
		return err
	}
	
	return nil
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registraions(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func (e Event) CancelRegister(userId int64) error {
	query := "DELETE FROM registraions WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id, userId)
	if err != nil {
		return err
	}
	
	return nil
}
