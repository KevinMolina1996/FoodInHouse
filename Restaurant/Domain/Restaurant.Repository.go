package Restaurant_Domain

import (
	database "FoodInHouse/Database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("restaurants")
var ctx = context.Background()

func Create(restaurant Restaurant) error {

	var err error

	_, err = collection.InsertOne(ctx, restaurant)

	if err != nil {
		return err
	}

	return nil
}

func Read() (Restaurants, error) {
	var restaurants Restaurants

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var restaurant Restaurant
		err = cur.Decode(&restaurant)

		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, &restaurant)
	}

	return restaurants, nil
}

func Update(restaurant Restaurant, RestaurantId string) error {
	var err error
	oid, _ := primitive.ObjectIDFromHex(RestaurantId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"Name":    restaurant.Name,
			"Address": restaurant.Address,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(RestaurantId string) error {
	return nil
}
