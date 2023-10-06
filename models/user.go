package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string   `json:"username" bson:"username"`
	Email            string   `json:"email" bson:"email"`
	Password         string   `json:"-" bson:"password"`
	FirstName        string   `json:"first_name" bson:"first_name"`
	LastName         string   `json:"last_name" bson:"last_name"`
	WrongPass        int      `json:"-" bson:"wrong_pass"`
	Permissions      []string `json:"permissions" bson:"permissions"`
	CreatedBy        *User    `json:"created_by" bson:"created_by"`
	UpdatedBy        *User    `json:"updated_by" bson:"updated_by"`
}

func (model *User) CollectionName() string {
	return "users"
}
