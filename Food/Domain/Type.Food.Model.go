package Food_Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Type struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `json:"name"`
}

type Types []*Type
