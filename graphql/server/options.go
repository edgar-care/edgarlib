package server

import (
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOptions(option model.Options) *options.FindOptions {
	findOptions := options.Find()

	// Apply sorting options
	if option.Sort != nil {
		var sortOrder int
		if option.Sort.Order == "ASC" {
			sortOrder = 1
		} else {
			sortOrder = -1
		}
		findOptions.SetSort(bson.D{{Key: option.Sort.Key, Value: sortOrder}})
	}

	// Apply limit
	if option.Limit != nil {
		findOptions.SetLimit(int64(*option.Limit))
	}

	// Apply offset
	if option.Offset != nil {
		findOptions.SetSkip(int64(*option.Offset))
	}

	return findOptions
}
