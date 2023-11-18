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
	return func(c *gin.Context) {
		accountId := c.Param("accountId")

		account, accountErr := data.GetAccount(accountId)

		if accountErr != nil {
			togglerError.SendException(c, togglerError.NotFound, "Account not found")
			return
		}

		c.JSON(
			http.StatusOK,
			responses.AccountResponse{
				Data: *account,
			},
		)
	}
}
func GetAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accounts []models.Account

		accounts, accountsErr := data.GetAccounts()
		if accountsErr != nil {
			togglerError.SendException(c, accountsErr.Error(), "Couldn't load accounts")
			return
		}

		c.JSON(
			http.StatusOK,
			responses.AccountsResponse{
				Data: accounts,
			},
		)
	}
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
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		account, accountErr := data.GetAccount(accountId)

		if accountErr != nil {
			togglerError.SendException(c, togglerError.NotFound, "Account not found")
			return
		}

		var updates models.Account
		if err := c.BindJSON(&updates); err != nil {
			log.Println(err.Error())
			togglerError.SendException(c, togglerError.BadRequest, "Couldn't bind")
			return
		}
		newAccount, newAccountErr := data.EditAccount(account, &updates)
		if newAccountErr != nil {
			log.Println(newAccountErr.Error())
			togglerError.SendException(c, togglerError.InternalError, "Couldn't save feature")
			return
		}
		c.JSON(
			http.StatusOK,
			responses.AccountResponse{
				Data: *newAccount,
			},
		)
	}
}
func DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId := c.Param("accountId")

		account, accountErr := data.GetAccount(accountId)

		if accountErr != nil {
			togglerError.SendException(c, accountErr.Error(), "Account not found")
			return
		}
		deleteErr := data.DeleteAccount(account, true)
		if deleteErr != nil {
			log.Println(deleteErr.Error())
			togglerError.SendException(c, togglerError.InternalError, "Couldn't delete account")
			return
		}
		c.Status(http.StatusNoContent)
		c.Done()
	}
}
