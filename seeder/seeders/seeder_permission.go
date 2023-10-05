package seeders

import (
	"learn-echo/constants"
	"learn-echo/helpers"
	"learn-echo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func SeedPermission() error {
	permissions := []models.Permission{
		{Code: constants.PER_PERMISSION_CREATE},
		{Code: constants.PER_PERMISSION_DELETE},
		{Code: constants.PER_PERMISSION_UPDATE},
		{Code: constants.PER_PERMISSION_VIEW},

		{Code: constants.PER_USER_CREATE},
		{Code: constants.PER_USER_DELETE},
		{Code: constants.PER_USER_UPDATE},
		{Code: constants.PER_USER_VIEW},
	}

	for _, permission := range permissions {
		if err := helpers.FirstOrCreate(&models.Permission{}, bson.M{"code": permission.Code}, &permission); err != nil {
			return err
		}
	}

	return nil
}
