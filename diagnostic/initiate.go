package diagnostic

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/diagnostic/utils"
	"github.com/edgar-care/edgarlib/graphql"
)

type InitiateResponse struct {
	Id   string
	Code int
	Err  error
}

func Initiate() InitiateResponse {
	gqlClient := graphql.CreateClient()

	utils.WakeNlpUp()

	session, err := graphql.CreateSession(context.Background(), gqlClient, []string{}, 0, 0, 0, "M", "", []graphql.LogsInput{}, []string{})

	if err != nil {
		return InitiateResponse{"", 500, errors.New("unable to create session")}
	}
	return InitiateResponse{session.CreateSession.Id, 200, nil}
}
