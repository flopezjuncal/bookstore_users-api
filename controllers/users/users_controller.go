package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/flopezjuncal/bookstore_users-api/domain/users"
	"github.com/flopezjuncal/bookstore_users-api/services"
	"github.com/flopezjuncal/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		fmt.Println("\nERROR!!!")
		c.JSON(restErr.Status, restErr)
		return
	}

	/*
		//LO COMENTADO EQUIVALE AL ShouldBindJSON(...)...
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			//TODO: Handle Error
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(bytes, &user)
		if err != nil {
			fmt.Println(err)
			//TODO: Handle JSON error
			//return
		}
	*/

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		fmt.Println("\nERROR!!!")
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println("SALIO TODO JOYA")
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "IMPLEMENT ME!")
}
