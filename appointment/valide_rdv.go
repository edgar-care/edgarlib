package appointment

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type ValidateRdvResponse struct {
	Rdv  model.Rdv
	Code int
	Err  error
}

type ReviewInput struct {
	Reason        string `json:"reason,omitempty"`
	Validation    bool   `json:"validation"`
	HealthMethode string `json:"health_method"`
}

func ValidateRdv(appointmentId string, input ReviewInput) EditRdvResponse {
	if appointmentId == "" {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetRdvById(appointmentId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	var appointment_status model.AppointmentStatus

	if input.Validation == true {
		appointment_status = model.AppointmentStatusAcceptedDueToReview
	} else {
		appointment_status = model.AppointmentStatusCanceledDueToReview
	}
	updatedRdv, err := graphql.UpdateRdv(appointmentId, model.UpdateRdvInput{
		CancelationReason: &input.Reason,
		AppointmentStatus: &appointment_status,
		HealthMethod:      &input.HealthMethode,
	})
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	if input.Validation == false {
		var new_appointment_status = model.AppointmentStatusOpened
		_, err = graphql.CreateRdv(model.CreateRdvInput{
			IDPatient:         "",
			DoctorID:          appointment.DoctorID,
			StartDate:         appointment.StartDate,
			EndDate:           appointment.EndDate,
			AppointmentStatus: new_appointment_status,
			SessionID:         "",
		})
		if err != nil {
			return EditRdvResponse{Rdv: model.Rdv{}, Code: 500, Err: errors.New("unable to create appointment")}
		}
	}

	return EditRdvResponse{
		Rdv:  updatedRdv,
		Code: 200,
		Err:  nil,
	}
}
