package Food_Domain

import (
	database "FoodInHouse/Database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("foods")
var ctx = context.Background()

func Create(food Food) error {

	var err error

	_, err = collection.InsertOne(ctx, food)

	if err != nil {
		return err
	}

	return nil
}

func Read() (Foods, error) {
	var foods Foods

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var food Food
		err = cur.Decode(&food)

		if err != nil {
			return nil, err
		}

		foods = append(foods, &food)
	}

	return foods, nil
}

func ReadById(id string) (Foods, error) {
	var foods Foods

	filter := bson.M{"_idRestaurant": id}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var food Food
		err = cur.Decode(&food)

		if err != nil {
			return nil, err
		}

		foods = append(foods, &food)
	}

	return foods, nil
}

func ReadTypes() (Types, error) {
	var types Types

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var ty Type
		err = cur.Decode(&ty)

		if err != nil {
			return nil, err
		}

		types = append(types, &ty)
	}

	return types, nil
}

func Update(food Food, FoodId string) error {

	var err error
	oid, _ := primitive.ObjectIDFromHex(FoodId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"Name":  food.Name,
			"Price": food.Price,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(FoodId string) error {
	return nil
}
