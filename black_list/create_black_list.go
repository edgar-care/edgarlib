package black_list

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateBlackListResponse struct {
	BlackList model.BlackList
	Code      int
	Err       error
}

func CreateBlackList(token string) CreateBlackListResponse {
	gqlClient := graphql.CreateClient()

	device, err := graphql.CreateBlackList(context.Background(), gqlClient, []string{token})
	if err != nil {
		return CreateBlackListResponse{BlackList: model.BlackList{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return CreateBlackListResponse{
		BlackList: model.BlackList{
			ID:    device.CreateBlackList.Id,
			Token: device.CreateBlackList.Token,
		},
		Code: 201,
		Err:  nil,
	}
}
