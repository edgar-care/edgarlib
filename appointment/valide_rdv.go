package appointment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type ValidateRdvResponse struct {
	Rdv  model.Rdv
	Code int
	Err  error
}

type ReviewInput struct {
	Reason        string `json:"reason,omitempty"`
	Validation    bool   `json:"validation"`
	HealthMethode string `json:"health_methode"`
}

func ValidateRdv(appointmentId string, input ReviewInput) EditRdvResponse {
	gqlClient := graphql.CreateClient()
	if appointmentId == "" {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	var appointment_status graphql.AppointmentStatus

	if input.Validation == true {
		appointment_status = graphql.AppointmentStatusAcceptedDueToReview
	} else {
		appointment_status = graphql.AppointmentStatusCanceledDueToReview
	}
	updatedRdv, err := graphql.UpdateRdv(context.Background(), gqlClient, appointmentId, appointment.GetRdvById.Id_patient, appointment.GetRdvById.Doctor_id, appointment.GetRdvById.Start_date, appointment.GetRdvById.End_date, input.Reason, appointment_status, appointment.GetRdvById.Session_id, input.HealthMethode)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	if input.Validation == false {
		var new_appointment_status = graphql.AppointmentStatusOpened
		_, err = graphql.CreateRdv(context.Background(), gqlClient, "", appointment.GetRdvById.Doctor_id, appointment.GetRdvById.Start_date, appointment.GetRdvById.End_date, new_appointment_status, "")
		if err != nil {
			return EditRdvResponse{Rdv: model.Rdv{}, Code: 500, Err: errors.New("unable to create appointment")}
		}
	}

	return EditRdvResponse{
		Rdv: model.Rdv{
			ID:                updatedRdv.UpdateRdv.Id,
			DoctorID:          updatedRdv.UpdateRdv.Doctor_id,
			IDPatient:         updatedRdv.UpdateRdv.Id_patient,
			StartDate:         updatedRdv.UpdateRdv.Start_date,
			EndDate:           updatedRdv.UpdateRdv.End_date,
			CancelationReason: &updatedRdv.UpdateRdv.Cancelation_reason,
			AppointmentStatus: model.AppointmentStatus(updatedRdv.UpdateRdv.Appointment_status),
			SessionID:         updatedRdv.UpdateRdv.Session_id,
			HealthMethod:      &updatedRdv.UpdateRdv.Health_method,
		},
		Code: 200,
		Err:  nil,
	}
}
