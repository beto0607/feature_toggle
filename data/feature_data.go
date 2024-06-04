package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"toggler/configs"
	"toggler/models"
	"toggler/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func featureCollection() *mongo.Collection {
	return configs.GetCollection("features")
}

const listTimeout = 20 * time.Second

func GetFeatures() ([]models.Feature, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), listTimeout)
	defer cancel()
	cur, err := featureCollection().Find(ctx, bson.M{"deleted_at": ""})
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}
	//reading from the db in an optimal way
	defer cur.Close(ctx)
	var features []models.Feature
	features = []models.Feature{}

	for cur.Next(ctx) {
		var singleFeature models.Feature
		if err = cur.Decode(&singleFeature); err != nil {
			log.Println("Couldn't parse feature")
			log.Println(err.Error())
			return nil, false
		}

		features = append(features, singleFeature)
	}

	return features, true
}

func GetFeature(featureId string) (*models.Feature, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), listTimeout)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(featureId)
	if err != nil {
		log.Println("Coudln't convert to ObjectID from: " + featureId)
		return nil, false
	}

    feature := models.Feature{}
	err = featureCollection().FindOne(ctx, bson.M{"_id": objectId, "deleted_at": ""}).Decode(&feature)
	if err != nil {
        if err == mongo.ErrNoDocuments{
            return nil, false
        }
		return nil, false
	}
	return &feature, true
}

func AddFeature(feature *models.Feature) (*models.Feature, bool) {
	insertTimeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), insertTimeout)
	defer cancel()

	newFeature := models.Feature{
		Id:        primitive.NewObjectID(),
		Name:      feature.Name,
		Enabled:   feature.Enabled,
		Flags:     feature.Flags,
		CreatedAt: utils.NowTimestamp(),
	}
	log.Println(newFeature.Name)
	_, err := featureCollection().InsertOne(ctx, newFeature)
	if err != nil {
		log.Println("Couldn't insert feature")
		log.Println(err.Error())
		return nil, false
	}

	return &newFeature, true

}

func EditFeature(featureId string, featureDto models.FeatureDto) (*models.Feature, bool) {
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(featureId)
	if err != nil {
		log.Println("Coudln't convert to ObjectID from: " + featureId)
		return nil, false
	}
	updates := bson.M{
		"updated_at": utils.NowTimestamp(),
	}

	if featureDto.Enabled != nil {
		updates["enabled"] = *featureDto.Enabled
	}
	if featureDto.Flags != nil {
		updates["flags"] = *featureDto.Flags
	}
	if featureDto.Name != nil {
		updates["name"] = *featureDto.Name
	}

	out, err := json.Marshal(updates)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	result := featureCollection().FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectId, "deleted_at": ""},
		bson.M{"$set": updates},
		options.FindOneAndUpdate().SetReturnDocument(options.After), // <- Set option to return document after update (important)
	)
	if result.Err() != nil {
		log.Println("Couldn't update feature")
		return nil, false
	}

	feature := models.Feature{}
	result.Decode(&feature)
	return &feature, true
}

func DeleteFeature(featureId string) bool {
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(featureId)
	if err != nil {
		log.Println("Coudln't convert to ObjectID from: " + featureId)
		return false
	}

	r, _ := featureCollection().DeleteOne(ctx, bson.M{"_id": objectId})
	return r.DeletedCount == 1

}

func SoftDeleteFeature(featureId string) bool {
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(featureId)
	if err != nil {
		return false
	}
	timestamp := utils.NowTimestamp()
	updates := bson.M{
		"updated_at": timestamp,
		"deleted_at": timestamp,
	}

	result := featureCollection().FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectId},
		bson.M{"$set": updates},
		options.FindOneAndUpdate().SetReturnDocument(options.After), // <- Set option to return document after update (important)
	)
	if result.Err() != nil {
		log.Println("Couldn't update feature")
		return false
	}

	return true
}
