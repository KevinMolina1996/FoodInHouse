package Order_Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Food      string             `json:"food"`
	Price     string             `json:"price"`
	User      string             `json:"_userId"`
	CreatedAt time.Time          `bson:"created_At" json:"created_at"`
}
type Orders []*Order
