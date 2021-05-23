package users

import (
	"fmt"

	"github.com/nononsensecode/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with id %d is not found", user.Id))
	}

	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	for _, u := range userDB {
		if u.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s is already registered", user.Email))
		}
	}

	current := userDB[user.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user with id %d already exists", user.Id))
	}

	userDB[user.Id] = user

	return nil
}
