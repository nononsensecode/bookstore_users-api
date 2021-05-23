package users

import (
	"strings"

	"github.com/nononsensecode/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

func (user *User) Validate() *errors.RestErr {
	if user.Id <= 0 {
		return errors.NewBadRequestError("user id cannot be less than or equal to zero")
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
