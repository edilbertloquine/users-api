package mysql_utils

import (
	"strings"

	"github.com/edilbertloquine/go-microservices/users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "sql: no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing mysql database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
