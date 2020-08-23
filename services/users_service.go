package services

import (
	"github.com/ramski29/bookstore_users-api/domain/user"
)

func CreateUser(user user.User) (*user.User, error) {
	return &user, nil
}
