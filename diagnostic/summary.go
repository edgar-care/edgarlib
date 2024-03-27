package diagnostic

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetSummaryResponse struct {
	SessionId   string
	Diseases    []model.SessionDiseases
	Symptoms    []model.SessionSymptom
	Age         int
	Height      int
	Weight      int
	Sex         string
	AnteChirs   []model.AnteChir
	AnteDisease []model.AnteDisease
	Treatments  []model.Treatment
	Logs        []graphql.LogsInput
	Alerts      []model.Alert
	Code        int
	Err         error
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

	var sessionDiseases []model.SessionDiseases
	for _, sessionDisease := range session.GetSessionById.Diseases {
		var nSD model.SessionDiseases
		nSD.Name = sessionDisease.Name
		nSD.Presence = sessionDisease.Presence
		sessionDiseases = append(sessionDiseases, nSD)
	}

	var sessionSymptoms []model.SessionSymptom
	for _, sessionSymptom := range session.GetSessionById.Symptoms {
		var nSS model.SessionSymptom
		nSS.Name = sessionSymptom.Name
		nSS.Presence = &sessionSymptom.Presence
		nSS.Duration = &sessionSymptom.Duration
		sessionSymptoms = append(sessionSymptoms, nSS)
	}

	var anteChirs []model.AnteChir
	for _, anteChirId := range session.GetSessionById.Ante_chirs {
		anteChir, _ := graphql.GetAnteChirByID(context.Background(), gqlClient, anteChirId)
		var nAC model.AnteChir
		nAC.ID = anteChir.GetAnteChirByID.Id
		nAC.Name = anteChir.GetAnteChirByID.Name
		nAC.Localisation = anteChir.GetAnteChirByID.Localisation
		nAC.InducedSymptoms = anteChir.GetAnteChirByID.Induced_symptoms
		anteChirs = append(anteChirs, nAC)
	}

	var anteDiseases []model.AnteDisease
	for _, anteDiseaseId := range session.GetSessionById.Ante_diseases {
		anteDisease, _ := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, anteDiseaseId)
		var nAD model.AnteDisease
		nAD.ID = anteDisease.GetAnteDiseaseByID.Id
		nAD.Name = anteDisease.GetAnteDiseaseByID.Name
		nAD.Chronicity = anteDisease.GetAnteDiseaseByID.Chronicity
		nAD.SurgeryIds = anteDisease.GetAnteDiseaseByID.Surgery_ids
		nAD.Symptoms = anteDisease.GetAnteDiseaseByID.Symptoms
		nAD.TreatmentIds = anteDisease.GetAnteDiseaseByID.Treatment_ids
		nAD.StillRelevant = anteDisease.GetAnteDiseaseByID.Still_relevant
		anteDiseases = append(anteDiseases, nAD)
	}

	var treatments []model.Treatment
	for _, treatmentId := range session.GetSessionById.Treatments {
		treatment, _ := graphql.GetTreatmentByID(context.Background(), gqlClient, treatmentId)
		var nT model.Treatment
		nT.ID = treatment.GetTreatmentByID.Id
		nT.Name = treatment.GetTreatmentByID.Name
		nT.Disease = treatment.GetTreatmentByID.Disease
		nT.Symptoms = treatment.GetTreatmentByID.Symptoms
		nT.SideEffects = treatment.GetTreatmentByID.Side_effects
		treatments = append(treatments, nT)
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
		SessionId:   session.GetSessionById.Id,
		Diseases:    sessionDiseases,
		Symptoms:    sessionSymptoms,
		Age:         session.GetSessionById.Age,
		Height:      session.GetSessionById.Height,
		Weight:      session.GetSessionById.Weight,
		Sex:         session.GetSessionById.Sex,
		AnteChirs:   anteChirs,
		AnteDisease: anteDiseases,
		Treatments:  treatments,
		Logs:        logs,
		Alerts:      alerts,
		Code:        200,
		Err:         nil,
	}
}
