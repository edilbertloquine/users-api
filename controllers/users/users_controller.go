package users

import (
	"net/http"

	"github.com/edilbertloquine/go-microservices/users-api/domain/users"
	"github.com/edilbertloquine/go-microservices/users-api/services"
	"github.com/edilbertloquine/go-microservices/users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }

	// err = json.Unmarshal(bytes, &user)
	// if err != nil {
	// 	log.Printf("Error: %v", err.Error())
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())

		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement this")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement this")
}
