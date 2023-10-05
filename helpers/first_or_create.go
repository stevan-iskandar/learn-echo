package helpers

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func FirstOrCreate(model mgm.Model, filter bson.M, data mgm.Model) error {
	if err := mgm.Coll(model).First(filter, model); err == nil {
		return nil
	}
	if err := mgm.Coll(data).Create(data); err != nil {
		return err
	}
	return nil
}
