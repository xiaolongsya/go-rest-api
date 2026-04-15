package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retriverdPassord string
	err := row.Scan(&u.ID, &retriverdPassord)

	if err != nil {
		return errors.New("invalid credentials")
	}
	passwordValid := utils.CheckPasswordHash(u.Password, retriverdPassord)
	if !passwordValid {
		return errors.New("invalid credentials")
	}
	return nil
}
