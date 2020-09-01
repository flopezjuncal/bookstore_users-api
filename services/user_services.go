package services

import (
	"github.com/flopezjuncal/bookstore_users-api/domain/users"
	"github.com/flopezjuncal/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	//WE ASSUME WE HAVE A VALID ID
	//if userId <= 0 {
	//	return nil, errors.NewBadRequestError("invalid user id")
	//}

	auxUser := &users.User{ID: userId}
	if err := auxUser.Get(); err != nil {
		return nil, err
	}
	return auxUser, nil
}
