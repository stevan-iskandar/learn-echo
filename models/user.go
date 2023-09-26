package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USER string = "users"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Age       int                `bson:"age"`
	Safe      int                `bson:"safe"`
	Code      string             `bson:"code"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
