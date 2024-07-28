package models

import "example.com/go-rest-api/db"

type User struct {
	Id       int64
	Email    string `binding:"required`
	Password string `binding:"required`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	u.Id = id
	return err
}