package controllers

import (
	"log"
	"net/http"
	"toggler/data"
	togglerError "toggler/error"
	"toggler/models"
	"toggler/responses"

	"github.com/gin-gonic/gin"
)

func GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
func GetAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account

		if err := c.BindJSON(&account); err != nil {
			log.Println(err.Error())
			togglerError.SendException(c, togglerError.BadRequest, "Couldn't bind")
			return
		}
		if validationErr := validate.Struct(&account); validationErr != nil {
			log.Println(validationErr.Error())
			togglerError.SendException(c, togglerError.BadRequest, "Invalid data")
			return
		}
		newAccount, newAccountErr := data.AddAccount(&account)
		if newAccountErr != nil {
			log.Println(newAccountErr.Error())
			togglerError.SendException(c, togglerError.InternalError, "Couldn't save feature")
			return
		}

		c.JSON(
			http.StatusCreated,
			responses.AccountResponse{
				Data: *newAccount,
			},
		)
	}
}
func EditAccount() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
func DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
