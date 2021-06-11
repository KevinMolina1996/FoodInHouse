package Order_Domain

import (
	database "FoodInHouse/Database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("orders")
var ctx = context.Background()

func Create(order Order) error {

	var err error

	_, err = collection.InsertOne(ctx, order)

	if err != nil {
		return err
	}

	return nil
}

func Read() (Orders, error) {
	var orders Orders

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var order Order
		err = cur.Decode(&order)

		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

func Update(order Order, OrderId string) error {
	return nil
}

func Delete(OrderId string) error {
	return nil
}
