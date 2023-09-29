package autoload

import (
	"os"

	"learn-echo/constants"

	"github.com/gookit/validate"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := mgm.SetDefaultConfig(nil, os.Getenv(constants.DB_NAME), options.Client().ApplyURI(os.Getenv(constants.MONGO_URI))); err != nil {
		panic(err)
	}

	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})
}
