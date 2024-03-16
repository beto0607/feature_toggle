package data

import (
	"time"
	"toggler/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var featureCollection *mongo.Collection = configs.GetCollection(configs.DB, "features")

const listTimeout = 20 * time.Second

// func GetFeatures() ([]models.Feature, int) {
// 	ctx, cancel := context.WithTimeout(context.Background(), listTimeout)
// 	defer cancel()
// 	cur, err := featureCollection.Find(ctx, bson.M{"accountId": accountId})
// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, http.StatusInternalServerError
// 	}
// 	//reading from the db in an optimal way
// 	defer cur.Close(ctx)
// 	var features []models.Feature
// 	features = []models.Feature{}
//
// 	for cur.Next(ctx) {
// 		var singleFeature models.Feature
// 		if err = cur.Decode(&singleFeature); err != nil {
// 			log.Println("Couldn't parse feature")
// 			log.Println(err.Error())
// 			return nil, errors.New(togglerError.InternalError)
// 		}
//
// 		features = append(features, singleFeature)
// 	}
//
// 	return features, nil
// }
//
// func AddFeature(feature *models.Feature) (*models.Feature, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
// 	defer cancel()
//
// 	newFeature := models.Feature{
// 		Id:        primitive.NewObjectID(),
// 		AccountId: feature.AccountId,
// 		Name:      feature.Name,
// 		Enabled:   feature.Enabled,
// 		Flags:     feature.Flags,
// 		CreatedBy: feature.CreatedBy,
// 		CreatedAt: time.Now().UTC().String(),
// 	}
// 	_, err := featureCollection.InsertOne(ctx, newFeature)
// 	if err != nil {
// 		log.Println("Couldn't insert feature")
// 		log.Println(err.Error())
// 		return nil, errors.New(togglerError.InternalError)
// 	}
//
// 	return &newFeature, nil
//
// }
