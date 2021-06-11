package User_Domain

import (
	database "FoodInHouse/Database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user User) error {

	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (Users, error) {
	var users Users

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func Update(user User, UserId string) error {
	var err error
	oid, _ := primitive.ObjectIDFromHex(UserId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":    user.Name,
			"email":   user.Email,
			"address": user.Address,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(UserId string) error {
	return nil
}
