package black_list

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetBlackListByIdResponse struct {
	BlackList model.BlackList
	Code      int
	Err       error
}

type GetBlackListResponse struct {
	BlackList []model.BlackList
	Code      int
	Err       error
}

func GetBlackListById(id string) GetBlackListByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.BlackList

	device, err := graphql.GetBlackListById(context.Background(), gqlClient, id)
	if err != nil {
		return GetBlackListByIdResponse{model.BlackList{}, 400, errors.New("id does not correspond to a slot")}
	}
	res = model.BlackList{
		ID:    device.GetBlackListById.Id,
		Token: device.GetBlackListById.Token,
	}
	return GetBlackListByIdResponse{res, 200, nil}
}

func GetBlackList() GetBlackListResponse {
	gqlClient := graphql.CreateClient()
	var res []model.BlackList

	devices, err := graphql.GetBlackList(context.Background(), gqlClient)
	if err != nil {
		return GetBlackListResponse{[]model.BlackList{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, device := range devices.GetBlackList {
		res = append(res, model.BlackList{
			ID:    device.Id,
			Token: device.Token,
		})
	}
	return GetBlackListResponse{res, 200, nil}
}
