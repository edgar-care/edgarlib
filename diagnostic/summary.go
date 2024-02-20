package diagnostic

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetSummaryResponse struct {
	SessionId string
	Symptoms  []model.SessionSymptom
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
	var sessionSymptoms []model.SessionSymptom
	for _, sessionSymptom := range session.GetSessionById.Symptoms {
		var nSS model.SessionSymptom
		nSS.Name = sessionSymptom.Name
		nSS.Presence = &sessionSymptom.Presence
		nSS.Duration = &sessionSymptom.Duration
		sessionSymptoms = append(sessionSymptoms, nSS)
	}

	return GetSummaryResponse{
		SessionId: session.GetSessionById.Id,
		Symptoms:  sessionSymptoms,
		Age:       session.GetSessionById.Age,
		Height:    session.GetSessionById.Height,
		Weight:    session.GetSessionById.Weight,
		Sex:       session.GetSessionById.Sex,
		Logs:      logs,
		Code:      200,
		Err:       nil,
	}
}
