package server

import (
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOptions(option model.Options) *options.FindOptions {
	findOptions := options.Find()

	if option.Sort != nil {
		var sortOrder int
		if option.Sort.Order == "ASC" {
			sortOrder = 1
		} else {
			sortOrder = -1
		}
		findOptions.SetSort(bson.D{{Key: option.Sort.Key, Value: sortOrder}})
	}

	findOptions.SetLimit(int64(option.Limit))

	findOptions.SetSkip(int64(option.Offset))

	return findOptions
}
