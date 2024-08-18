package black_list

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateBlackListResponse struct {
	BlackList model.BlackList
	Code      int
	Err       error
}

func CreateBlackList(token string) CreateBlackListResponse {
	device, err := graphql.CreateBlackList(model.CreateBlackListInput{Token: []string{token}})
	if err != nil {
		return CreateBlackListResponse{BlackList: model.BlackList{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return CreateBlackListResponse{
		BlackList: device,
		Code:      201,
		Err:       nil,
	}
}
