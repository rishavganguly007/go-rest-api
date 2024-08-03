package models

import (
	"errors"

	"example.com/go-rest-api/db"
	"example.com/go-rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	u.Id = id
	return err
}

func (u User) ValidateCredentials() error{
	query := "SELECT id, password FROM users WHERE email=?"

	row := db.DB.QueryRow(query, u.Email)
	var retreivedPassword string
	err := row.Scan(&u.Id, &retreivedPassword)

	if err != nil {
		return err
	}
	passwordValid := utils.CheckPasswordHash(u.Password, retreivedPassword)

	if !passwordValid {
		return errors.New("credentials Invalid")
	}

	return nil 
}