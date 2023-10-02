package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"learn-echo/helpers"
	"learn-echo/models"
	"learn-echo/validations"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type UserMGM struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string   `bson:"username" mgm:"unique"`
	Email            string   `json:"email" bson:"email"`
	Password         string   `json:"password" bson:"password"`
	FirstName        string   `json:"first_name" bson:"first_name"`
	LastName         string   `json:"last_name" bson:"last_name"`
	WrongPass        int      `json:"wrong_pass" bson:"wrong_pass"`
	CreatedBy        *UserMGM `json:"created_by" bson:"created_by"`
	UpdatedBy        *UserMGM `json:"updated_by" bson:"updated_by"`
}

func (model *UserMGM) CollectionName() string {
	return "users_mgm"
}

func Root(c echo.Context) error {
	// Test write speed.
	startTime := time.Now()

	for i := 1; i <= 1000; i++ {
		password, _ := helpers.HashPassword(fmt.Sprintf("password*%d*", i))
		user := &UserMGM{
			Username:  fmt.Sprintf("user%d", i),
			Email:     fmt.Sprintf("user%d@email.com", i),
			Password:  password,
			FirstName: fmt.Sprintf("first%d", i),
			LastName:  fmt.Sprintf("last%d", i),
			WrongPass: i % 2,
		}
		if err := mgm.Coll(user).Create(user); err != nil {
			log.Fatal(err)
		}
	}

	writeDuration := fmt.Sprintf("Write Time: %v\n", time.Since(startTime))

	// Test read speed.
	startTime = time.Now()
	var users []UserMGM
	if err := mgm.Coll(&UserMGM{}).SimpleFind(&users, bson.M{}); err != nil {
		log.Fatal(err)
	}

	readDuration := fmt.Sprintf("Read Time: %v\n", time.Since(startTime))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"write_duration": writeDuration,
		"read_duration":  readDuration,
		"users":          users,
	})
}

func Store(c echo.Context) error {
	userForm := c.Get(validations.STORE_VALIDATION).(*validations.UserStoreForm)

	user := &models.User{
		Name:  userForm.Name,
		Email: userForm.Email,
		Age:   userForm.Age,
		Safe:  userForm.Safe,
		Code:  userForm.Code,
	}

	if err := mgm.Coll(user).Create(user); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, user)
}
