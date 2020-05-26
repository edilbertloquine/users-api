package services

import (
	"github.com/edilbertloquine/go-microservices/users-api/domain/users"
	"github.com/edilbertloquine/go-microservices/users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func GetUser() {

}

func FindUser() {

}
