package appointment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetWaitingReviewResponse struct {
	Rdv  []model.Rdv
	Code int
	Err  error
}

func GetWaitingReview(doctorId string) GetWaitingReviewResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Rdv
	rdv, err := graphql.GetWaitingRdv(context.Background(), gqlClient, doctorId)

	if err != nil {
		return GetWaitingReviewResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a doctor")}
	}
	for _, appointment := range rdv.GetWaitingRdv {
		temp := appointment.Cancelation_reason
		res = append(res, model.Rdv{
			ID:                appointment.Id,
			DoctorID:          appointment.Doctor_id,
			IDPatient:         appointment.Id_patient,
			StartDate:         appointment.Start_date,
			EndDate:           appointment.End_date,
			CancelationReason: &temp,
			AppointmentStatus: model.AppointmentStatus(appointment.Appointment_status),
			SessionID:         appointment.Session_id,
		})
	}
	return GetWaitingReviewResponse{res, 200, nil}
}
