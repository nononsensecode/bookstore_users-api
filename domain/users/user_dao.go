package users

import (
	"log"

	"github.com/nononsensecode/bookstore_users-api/app/datasources/mysql/users_db"
	"github.com/nononsensecode/bookstore_users-api/utils/errors"
	"github.com/nononsensecode/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsert = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	userSelect  = "SELECT * FROM users WHERE id = ?"
	queryUpdate = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	queryDelete = "DELETE FROM users WHERE id = ?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(userSelect)
	if err != nil {
		log.Println(err)
		return errors.NewInternalServerError("data cannot be retrieved")
	}
	defer stmt.Close()

	userResult := stmt.QueryRow(user.Id)
	err = userResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		log.Println(err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	users_db.Ping()

	stmt, err := users_db.Client.Prepare(queryInsert)
	if err != nil {
		log.Println(err)
		return errors.NewInternalServerError("data cannot be saved")
	}
	defer stmt.Close()

	saveResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := saveResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdate)
	if err != nil {
		log.Println(err)
		return errors.NewInternalServerError("data cannot be updated")
	}
	defer stmt.Close()

	updateResult, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	_, updateErr = updateResult.RowsAffected()
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDelete)
	if err != nil {
		log.Println(err)
		return errors.NewInternalServerError("data cannot be deleted")
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}
