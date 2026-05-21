package models

import (
	"errors"
	"example/com/rest-api/db"
	"example/com/rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.Id = userId
	return nil
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	isPasswordValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
