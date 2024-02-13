package diagnostic

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/diagnostic/utils"
	"github.com/edgar-care/edgarlib/graphql"
)

type GetSummaryResponse struct {
	SessionId string
	Symptoms  []utils.Symptom
	Age       int
	Height    int
	Weight    int
	Sex       string
	Logs      []graphql.LogsInput
	Code      int
	Err       error
}

func GetSummary(id string) GetSummaryResponse {
	if id == "" {
		return GetSummaryResponse{Code: 400, Err: errors.New("id is required")}
	}
	gqlClient := graphql.CreateClient()
	session, err := graphql.GetSessionById(context.Background(), gqlClient, id)
	if err != nil {
		return GetSummaryResponse{Code: 400, Err: errors.New("id does not correspond to a session")}
	}
	var logs []graphql.LogsInput
	for _, log := range session.GetSessionById.Logs {
		logs = append(logs, graphql.LogsInput{
			Question: log.Question,
			Answer:   log.Answer,
		})
	}

	return GetSummaryResponse{
		SessionId: session.GetSessionById.Id,
		Symptoms:  utils.StringToSymptoms(session.GetSessionById.Symptoms),
		Age:       session.GetSessionById.Age,
		Height:    session.GetSessionById.Height,
		Weight:    session.GetSessionById.Weight,
		Sex:       session.GetSessionById.Sex,
		Logs:      logs,
		Code:      200,
		Err:       nil,
	}
}
