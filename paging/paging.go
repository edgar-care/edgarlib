package paging

import "github.com/edgar-care/edgarlib/v2/graphql/model"

func CreatePagingOption(page int, size int) *model.Options {
	offset := page * size
	return &model.Options{
		Sort: &model.SortingOptions{
			Order: "ASC",
			Key:   "createdAt",
		},
		Limit:  size,
		Offset: offset,
	}
}
