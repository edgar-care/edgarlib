package black_list

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
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
	device, err := graphql.GetBlackListById(id)
	if err != nil {
		return GetBlackListByIdResponse{model.BlackList{}, 400, errors.New("id does not correspond to a slot")}
	}
	return GetBlackListByIdResponse{device, 200, nil}
}

func GetBlackList() GetBlackListResponse {
	devices, err := graphql.GetBlackList(nil)
	if err != nil {
		return GetBlackListResponse{[]model.BlackList{}, 400, errors.New("invalid input: " + err.Error())}
	}

	return GetBlackListResponse{devices, 200, nil}
}
