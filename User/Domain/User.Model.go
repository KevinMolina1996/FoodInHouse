package User_Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Address   string             `json:"address"`
	CreatedAt time.Time          `bson:"created_At" json:"created_at"`
}
type Users []*User
