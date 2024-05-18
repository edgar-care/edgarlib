package appointment

import (
	"context"
	"errors"
	diagnos "github.com/edgar-care/edgarlib/diagnostic"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type SessionSummary struct {
	Diseases  []model.SessionDiseases `json:"diseases"`
	Fiability float64                 `json:"fiability"`
	Symptoms  []model.SessionSymptom  `json:"symptoms"`
	Logs      []graphql.LogsInput     `json:"logs"`
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

func GetWaitingReview(doctorId string) GetWaitingReviewResponse {
	gqlClient := graphql.CreateClient()
	rdv, err := graphql.GetWaitingRdv(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetWaitingReviewResponse{RdvWithSession: nil, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	var rdvWithSession []RdvSessionPair

	for _, appointment := range rdv.GetWaitingRdv {
		sessionResponse := diagnos.GetSummary(appointment.Session_id)
		if sessionResponse.Err != nil {
			return GetWaitingReviewResponse{RdvWithSession: nil, Code: sessionResponse.Code, Err: sessionResponse.Err}
		}

		temp := appointment.Cancelation_reason

		rdvWithSession = append(rdvWithSession, RdvSessionPair{
			Rdv: model.Rdv{
				ID:                appointment.Id,
				DoctorID:          appointment.Doctor_id,
				IDPatient:         appointment.Id_patient,
				StartDate:         appointment.Start_date,
				EndDate:           appointment.End_date,
				CancelationReason: &temp,
				AppointmentStatus: model.AppointmentStatus(appointment.Appointment_status),
				SessionID:         sessionResponse.SessionId,
			},
			Session: SessionSummary{
				Diseases:  sessionResponse.Diseases,
				Fiability: sessionResponse.Fiability,
				Symptoms:  sessionResponse.Symptoms,
				Logs:      sessionResponse.Logs,
				Alerts:    sessionResponse.Alerts,
			},
		})
	}

	return GetWaitingReviewResponse{RdvWithSession: rdvWithSession, Code: 200, Err: nil}
}
