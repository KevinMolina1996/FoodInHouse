package Food_Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `json:"name"`
	Price      float32            `json:"price"`
	Type       string             `json:"_idType"`
	Restaurant string             `bson:"_idRestaurant" json:"_idRestaurant"`
	CreatedAt  time.Time          `bson:"created_At" json:"created_at"`
}

type Foods []*Food
