package data

import (
	"context"
	"errors"
	"log"
	"time"
	"toggler/configs"
	togglerError "toggler/error"
	"toggler/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountCollection *mongo.Collection = configs.GetCollection(configs.DB, "accounts")

func GetAccount(accountId string) (*models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objId, hexErr := primitive.ObjectIDFromHex(accountId)
	if hexErr != nil {
		return nil, errors.New(togglerError.BadRequest)
	}

	var account models.Account
	err := accountCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&account)

	if err != nil {
		return nil, errors.New(togglerError.NotFound)
	}
	return &account, nil
}
func AddAccount(account *models.Account) (*models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newAccount := models.Account{
		Id:        primitive.NewObjectID(),
		Name:      account.Name,
		Features:  []string{},
		Users:     []string{},
		CreatedBy: account.CreatedBy,
		CreatedAt: time.Now().UTC().String(),
	}
	_, err := accountCollection.InsertOne(ctx, newAccount)
	if err != nil {
		log.Println("Couldn't insert account")
		log.Println(err.Error())
		return nil, errors.New(togglerError.InternalError)
	}
	return &newAccount, nil
}
