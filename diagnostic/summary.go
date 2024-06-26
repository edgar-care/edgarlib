package diagnostic

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetSummaryResponse struct {
	SessionId string
	Diseases  []model.SessionDiseases
	Fiability float64
	Symptoms  []model.SessionSymptom
	Logs      []graphql.LogsInput
	Alerts    []model.Alert
	Code      int
	Err       error
}

func GetSummary(id string) GetSummaryResponse {
	if id == "" {
		return GetSummaryResponse{Code: 400, Err: errors.New("id is required")}
	}
	gqlClient := graphql.CreateClient()
	session, err := graphql.GetSessionById(context.Background(), gqlClient, id)
	symptoms, err := graphql.GetSymptoms(context.Background(), gqlClient)
	diseases, err := graphql.GetDiseases(context.Background(), gqlClient)
	if err != nil {
		return GetSummaryResponse{Code: 400, Err: errors.New("id does not correspond to a session")}
	}

	var sessionDiseases []model.SessionDiseases
	for _, sessionDisease := range session.GetSessionById.Diseases {
		var nSD model.SessionDiseases
		for _, d := range diseases.GetDiseases {
			if sessionDisease.Name == d.Code && d.Name != "" {
				nSD.Name = d.Name
				break
			} else {
				nSD.Name = sessionDisease.Name
			}
		}
		nSD.Presence = sessionDisease.Presence
		nSD.UnknownPresence = sessionDisease.Unknown_presence
		sessionDiseases = append(sessionDiseases, nSD)
	}

	var fiability float64
	fiability = 0.42 //todo: Add a fiability system

	var sessionSymptoms []model.SessionSymptom
	for _, sessionSymptom := range session.GetSessionById.Symptoms {
		var nSS model.SessionSymptom
		for _, s := range symptoms.GetSymptoms {
			if sessionSymptom.Name == s.Code && s.Name != "" {
				nSS.Name = s.Name
				break
			} else {
				nSS.Name = sessionSymptom.Name
			}
		}
		nSS.Presence = sessionSymptom.Presence
		dura := sessionSymptom.Duration
		nSS.Duration = &dura
		sessionSymptoms = append(sessionSymptoms, nSS)
	}

	var logs []graphql.LogsInput
	for _, log := range session.GetSessionById.Logs {
		logs = append(logs, graphql.LogsInput{
			Question: log.Question,
			Answer:   log.Answer,
		})
	}

	var alerts []model.Alert
	for _, alertId := range session.GetSessionById.Alerts {
		alert, _ := graphql.GetAlertById(context.Background(), gqlClient, alertId)
		var nA model.Alert
		nA.ID = alert.GetAlertById.Id
		nA.Name = alert.GetAlertById.Name
		nA.Sex = &alert.GetAlertById.Sex
		nA.Height = &alert.GetAlertById.Height
		nA.Weight = &alert.GetAlertById.Weight
		nA.Symptoms = alert.GetAlertById.Symptoms
		nA.Comment = alert.GetAlertById.Comment
		alerts = append(alerts, nA)
	}

	return GetSummaryResponse{
		SessionId: session.GetSessionById.Id,
		Diseases:  sessionDiseases,
		Fiability: fiability,
		Symptoms:  sessionSymptoms,
		Logs:      logs,
		Alerts:    alerts,
		Code:      200,
		Err:       nil,
	}
}
