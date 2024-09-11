package appointment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/paging"
)

type SessionSummary struct {
	Diseases  []model.SessionDiseases `json:"diseases"`
	Fiability float64                 `json:"fiability"`
	Symptoms  []model.SessionSymptom  `json:"symptoms"`
	Logs      []model.LogsInput       `json:"logs"`
	Alerts    []model.Alert           `json:"alerts"`
}

type RdvSessionPair struct {
	Rdv     model.Rdv      `json:"rdv"`
	Session SessionSummary `json:"session"`
}

type GetWaitingReviewResponse struct {
	RdvWithSession []RdvSessionPair
	Code           int
	Err            error
}

func GetWaitingReview(doctorId string, page int, size int) GetWaitingReviewResponse {
	rdv, err := graphql.GetWaitingRdv(doctorId, paging.CreatePagingOption(page, size))
	if err != nil {
		return GetWaitingReviewResponse{RdvWithSession: nil, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	var rdvWithSession []RdvSessionPair

	for _, appointment := range rdv {
		//sessionResponse := diagnos.GetSummary(appointment.SessionID)
		//if sessionResponse.Err != nil {
		//	return GetWaitingReviewResponse{RdvWithSession: nil, Code: sessionResponse.Code, Err: sessionResponse.Err}
		//}

		rdvWithSession = append(rdvWithSession, RdvSessionPair{
			Rdv: appointment,
			//Session: SessionSummary{
			//	Diseases:  sessionResponse.Diseases,
			//	Fiability: sessionResponse.Fiability,
			//	Symptoms:  sessionResponse.Symptoms,
			//	Logs:      sessionResponse.Logs,
			//	Alerts:    sessionResponse.Alerts,
			//},
		})
	}

	return GetWaitingReviewResponse{RdvWithSession: rdvWithSession, Code: 200, Err: nil}
}
