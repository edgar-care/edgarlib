package black_list

import (
	"errors"
	"net/http"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type UpdateBlackListResponse struct {
	BlackList model.BlackList
	Code      int
	Err       error
}

func UpdateBlackList(token string) UpdateBlackListResponse {
	list, err := graphql.GetBlackList(nil)
	if err != nil {
		return UpdateBlackListResponse{BlackList: model.BlackList{}, Code: http.StatusBadRequest, Err: errors.New("unable to fetch blacklist")}
	}

	var updatedBlackList model.BlackList
	if list == nil || len(list) == 0 {
		createdBlackList, err := graphql.CreateBlackList(model.CreateBlackListInput{Token: []string{token}})
		if err != nil {
			return UpdateBlackListResponse{BlackList: model.BlackList{}, Code: http.StatusInternalServerError, Err: errors.New("unable to create blacklist")}
		}
		updatedBlackList = createdBlackList
	} else {
		updatedTokens := append(list[0].Token, token)
		auth, err := graphql.UpdateBlackList(list[0].ID, model.UpdateBlackListInput{Token: updatedTokens})
		if err != nil {
			return UpdateBlackListResponse{BlackList: model.BlackList{}, Code: http.StatusInternalServerError, Err: errors.New("unable to update blacklist")}
		}
		updatedBlackList = auth
	}

	return UpdateBlackListResponse{
		BlackList: updatedBlackList,
		Code:      http.StatusOK,
		Err:       nil,
	}
}
