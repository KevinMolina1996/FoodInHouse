package Restaurant_Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `bson:"created_At" json:"created_at"`
	Address   string             `json:"address"`
}

type Restaurants []*Restaurant
