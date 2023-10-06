package helpers

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func FilterFromMap(input map[string]interface{}) bson.M {
	bsonData := bson.M{}

	for key, value := range input {
		if value == nil || value == "" {
			continue
		}

		valueType := reflect.TypeOf(value)
		switch valueType.Kind() {
		case reflect.Slice:
			bsonData[key] = bson.M{"$in": value}
		default:
			bsonData[key] = value
		}
	}

	return bsonData
}
