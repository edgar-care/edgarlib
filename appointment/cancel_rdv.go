package appointment

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type CancelRdvResponse struct {
	Reason string
	Code   int
	Err    error
}

func CancelRdv(id string, reason string) CancelRdvResponse {
	if id == "" {
		return CancelRdvResponse{Reason: "", Code: 400, Err: errors.New("id is required")}
	}

	rdv, err := graphql.GetRdvById(id)
	if err != nil {
		return CancelRdvResponse{Reason: "", Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	var appointment_status = model.AppointmentStatusCanceled
	_, err = graphql.UpdateRdv(id, model.UpdateRdvInput{
		CancelationReason: &reason,
		AppointmentStatus: &appointment_status,
	})
	if err != nil {
		return CancelRdvResponse{Reason: "", Code: 500, Err: errors.New("unable to update appointment")}
	}

	var new_appointment_status = model.AppointmentStatusOpened
	_, err = graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          rdv.DoctorID,
		StartDate:         rdv.StartDate,
		EndDate:           rdv.EndDate,
		AppointmentStatus: new_appointment_status,
		SessionID:         "",
	})
	if err != nil {
		return CancelRdvResponse{Reason: "", Code: 500, Err: errors.New("unable to update appointment")}
	}
	return CancelRdvResponse{reason, 200, nil}
}
