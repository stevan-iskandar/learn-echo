package helpers

import (
	"context"
	"fmt"
	"learn-echo/structs"
	"strconv"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagination(c echo.Context, model mgm.Model, data interface{}) (*structs.Pagination, error) {
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))          // Page number
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size")) // Items per page

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10 // Default page size
	}

	// Calculate skip value to skip records on previous pages
	skip := (page - 1) * pageSize

	// Define options for MongoDB query
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64(skip))

	// Create a filter to find all users (you can customize this filter)
	filter := bson.M{}

	// Count the total number of records
	total, err := mgm.Coll(model).CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	sortFields := bson.D{}
	for i := 0; ; i++ {
		field := c.QueryParam(fmt.Sprintf("sort[%d]field", i))
		// println(fmt.Sprintf("field_%d: %s", i, field))
		if field == "" {
			break
		}

		order := c.QueryParam(fmt.Sprintf("sort[%d]order", i))
		// println(fmt.Sprintf("order_%d: %s", i, order))
		if order == "" {
			return nil, CustomError("Invalid sort parameter")
		}

		// Check if the order is valid
		if order != "asc" && order != "desc" {
			return nil, CustomError("Invalid sort order")
		}

		// Append the field and order to the sortFields
		sortFields = append(sortFields, bson.E{
			Key:   field,
			Value: order,
		})
	}
	println(sortFields)

	// Set the sorting options
	if len(sortFields) > 0 {
		findOptions.SetSort(sortFields)
	}

	// Query the database
	if err := mgm.Coll(model).SimpleFind(&data, filter, findOptions); err != nil {
		return nil, err
	}

	return &structs.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     data,
	}, nil
}
