package diagnostic

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"math"
	"sort"
)

type GetSummaryResponse struct {
	SessionId string
	Diseases  []model.SessionDiseases
	Fiability float64
	Symptoms  []model.SessionSymptom
	Logs      []model.LogsInput
	Alerts    []model.Alert
	Code      int
	Err       error
}

func calcFiability(sessionDiseases []model.SessionDiseases) float64 {
	if len(sessionDiseases) == 0 {
		return 0
	}

	sort.Slice(sessionDiseases, func(i, j int) bool {
		return sessionDiseases[i].Presence > sessionDiseases[j].Presence
	})

	firstElementPresence := sessionDiseases[0].Presence

	factorGap := 1.0
	if len(sessionDiseases) >= 2 {
		gap := sessionDiseases[0].Presence - sessionDiseases[1].Presence/2
		factorGap = math.Min(gap, 1.0)
	}

	factorPresence := firstElementPresence * 1.2

	reliability := factorPresence * factorGap

	return math.Max(0, math.Min(reliability, 0.946))
}

func GetSummary(id string) GetSummaryResponse {
	if id == "" {
		return GetSummaryResponse{Code: 400, Err: errors.New("id is required")}
	}
	session, err := graphql.GetSessionById(id)
	symptoms, err := graphql.GetSymptoms(nil)
	diseases, err := graphql.GetDiseases(nil)
	if err != nil {
		return GetSummaryResponse{Code: 400, Err: errors.New("id does not correspond to a session")}
	}

	var sessionDiseases []model.SessionDiseases
	for _, sessionDisease := range session.Diseases {
		var nSD model.SessionDiseases
		for _, d := range diseases {
			if sessionDisease.Name == d.Code && d.Name != "" {
				nSD.Name = d.Name
				break
			} else {
				nSD.Name = sessionDisease.Name
			}
		}
		nSD.Presence = sessionDisease.Presence
		nSD.UnknownPresence = sessionDisease.UnknownPresence
		sessionDiseases = append(sessionDiseases, nSD)
	}

	var fiability float64
	fiability = calcFiability(sessionDiseases)

	var sessionSymptoms []model.SessionSymptom
	for _, sessionSymptom := range session.Symptoms {
		var nSS model.SessionSymptom
		for _, s := range symptoms {
			if sessionSymptom.Name == s.Code && s.Name != "" {
				nSS.Name = s.Name
				break
			} else {
				nSS.Name = sessionSymptom.Name
			}
		}
		nSS.Presence = sessionSymptom.Presence
		nSS.Duration = sessionSymptom.Duration
		sessionSymptoms = append(sessionSymptoms, nSS)
	}

	var logs []model.LogsInput
	for _, log := range session.Logs {
		logs = append(logs, model.LogsInput{
			Question: log.Question,
			Answer:   log.Answer,
		})
	}

	var alerts []model.Alert
	for _, alertId := range session.Alerts {
		alert, _ := graphql.GetAlertById(alertId)
		var nA model.Alert
		nA.ID = alert.ID
		nA.Name = alert.Name
		nA.Sex = alert.Sex
		nA.Height = alert.Height
		nA.Weight = alert.Weight
		nA.Symptoms = alert.Symptoms
		nA.Comment = alert.Comment
		alerts = append(alerts, nA)
	}

	return GetSummaryResponse{
		SessionId: session.ID,
		Diseases:  sessionDiseases,
		Fiability: fiability,
		Symptoms:  sessionSymptoms,
		Logs:      logs,
		Alerts:    alerts,
		Code:      200,
		Err:       nil,
	}
}
