package black_list

import (
	"context"
	"errors"
	"net/http"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type UpdateBlackListResponse struct {
	BlackList model.BlackList
	Code      int
	Err       error
}

func UpdateBlackList(token string) UpdateBlackListResponse {
	gqlClient := graphql.CreateClient()

	list, err := graphql.GetBlackList(context.Background(), gqlClient)
	if err != nil {
		return UpdateBlackListResponse{BlackList: model.BlackList{}, Code: http.StatusBadRequest, Err: errors.New("unable to fetch blacklist")}
	}

	var updatedBlackList model.BlackList
	if list == nil || len(list.GetBlackList) == 0 {
		createdBlackList, err := graphql.CreateBlackList(context.Background(), gqlClient, []string{token})
		if err != nil {
			return UpdateBlackListResponse{BlackList: model.BlackList{}, Code: http.StatusInternalServerError, Err: errors.New("unable to create blacklist")}
		}
		updatedBlackList = model.BlackList{
			ID:    createdBlackList.CreateBlackList.Id,
			Token: createdBlackList.CreateBlackList.Token,
		}
	} else {
		updatedTokens := append(list.GetBlackList[0].Token, token)
		auth, err := graphql.UpdateBlackList(context.Background(), gqlClient, list.GetBlackList[0].Id, updatedTokens)
		if err != nil {
			return UpdateBlackListResponse{BlackList: model.BlackList{}, Code: http.StatusInternalServerError, Err: errors.New("unable to update blacklist")}
		}
		updatedBlackList = model.BlackList{
			ID:    auth.UpdateBlackList.Id,
			Token: auth.UpdateBlackList.Token,
		}
	}

	return UpdateBlackListResponse{
		BlackList: updatedBlackList,
		Code:      http.StatusOK,
		Err:       nil,
	}
}
