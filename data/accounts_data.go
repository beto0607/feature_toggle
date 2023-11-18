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
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	objId, hexErr := primitive.ObjectIDFromHex(accountId)
	if hexErr != nil {
		return nil, errors.New(togglerError.BadRequest)
	}

	var account models.Account
	err := accountCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&account)

	if err != nil {
		return nil, errors.New(togglerError.NotFound)
	}
	return &account, nil
}

func GetAccounts() ([]models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), listTimeout)
	defer cancel()
	cur, err := accountCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(togglerError.InternalError)
	}
	//reading from the db in an optimal way
	defer cur.Close(ctx)
	var accounts []models.Account
	accounts = []models.Account{}

	for cur.Next(ctx) {
		var singleAccount models.Account
		if err = cur.Decode(&singleAccount); err != nil {
			log.Println("Couldn't parse account")
			log.Println(err.Error())
			return nil, errors.New(togglerError.InternalError)
		}

		accounts = append(accounts, singleAccount)
	}

	return accounts, nil
}

func AddAccount(account *models.Account) (*models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	newAccount := models.Account{
		Name:      account.Name,
		Features:  []primitive.ObjectID{},
		Users:     []primitive.ObjectID{},
		CreatedBy: account.CreatedBy,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	result, err := accountCollection.InsertOne(ctx, newAccount)
	if err != nil {
		log.Println("Couldn't insert account")
		log.Println(err.Error())
		return nil, errors.New(togglerError.InternalError)
	}
	newAccount.Id = result.InsertedID.(primitive.ObjectID)
	return &newAccount, nil
}

func EditAccount(account *models.Account, updates *models.Account) (*models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	account.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	account.Name = updates.Name
	account.Users = updates.Users
	account.Features = updates.Features

	_, err := accountCollection.UpdateByID(ctx, account.Id, bson.M{"$set": account})
	if err != nil {
		log.Println("Couldn't update account")
		log.Println(err.Error())
		return nil, errors.New(togglerError.InternalError)
	}
	return account, nil
}

func DeleteAccount(account *models.Account, hardDelete bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	if !hardDelete {
		updates := bson.M{
			"deletedAt": primitive.NewDateTimeFromTime(time.Now()),
		}
		_, err := accountCollection.UpdateByID(ctx, account.Id, bson.M{"$set": updates})
		return err

	}
	_, err := accountCollection.DeleteOne(ctx, bson.M{"_id": account.Id})
	if err != nil {
		log.Println("Couldn't delete account")
		log.Println(err.Error())
		return errors.New(togglerError.InternalError)
	}
	return nil
}
