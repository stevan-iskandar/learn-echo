package models

import "github.com/kamva/mgm/v3"

type Permission struct {
	mgm.DefaultModel `bson:",inline"`
	Code             string `json:"code" bson:"code"`
	CreatedBy        *User  `json:"created_by" bson:"created_by"`
	UpdatedBy        *User  `json:"updated_by" bson:"updated_by"`
}

func (model *Permission) CollectionName() string {
	return "permissions"
}
