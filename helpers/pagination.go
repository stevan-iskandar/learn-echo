package helpers

import (
	"context"
	"fmt"
	"learn-echo/structs"
	"strconv"
	"strings"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagination(c echo.Context, model mgm.Model, data interface{}, filterMap map[string]interface{}) (*structs.Pagination, error) {
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))          // Page number
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size")) // Items per page
	sorts := c.QueryParams()["sort"]                       // Get an array of sorting parameters

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

	// Create a filter to find all data (you can customize this filter)
	filter := bson.M{}
	if filterMap != nil {
		filter = FilterFromMap(filterMap)
	}

	// Count the total number of records
	total, err := mgm.Coll(model).CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	// Apply sorting based on the request parameters
	sortFields := bson.D{}
	for _, sortParam := range sorts {
		// Split the sort parameter into field and order
		parts := strings.Split(sortParam, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid sort parameter")
		}

		// Check if the order is valid
		if parts[1] != "asc" && parts[1] != "desc" {
			return nil, fmt.Errorf("invalid sort order")
		}

		field := parts[0]
		order := 1 // 1 = asc, -1 = desc

		if parts[1] == "desc" {
			order = -1
		}

		// Append the field and order to the sortFields
		sortFields = append(sortFields, bson.E{
			Key:   field,
			Value: order,
		})
	}

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
