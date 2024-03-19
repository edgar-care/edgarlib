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
	Reason     string `json:"reason"`
	Validation bool   `json:"validation"`
}

func ValidateRdv(appointmentId string, reason string, validation bool) EditRdvResponse {
	gqlClient := graphql.CreateClient()
	if appointmentId == "" {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	var appointment_status graphql.AppointmentStatus

	if validation == true {
		appointment_status = graphql.AppointmentStatusAcceptedduetoreview
	} else {
		appointment_status = graphql.AppointmentStatusCanceledduetoreview
	}
	updatedRdv, err := graphql.UpdateRdv(context.Background(), gqlClient, appointmentId, appointment.GetRdvById.Id_patient, appointment.GetRdvById.Doctor_id, appointment.GetRdvById.Start_date, appointment.GetRdvById.End_date, reason, appointment_status, appointment.GetRdvById.Sessions_ids)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Code: 500, Err: errors.New("unable to update appointment")}
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
			SessionsIds:       updatedRdv.UpdateRdv.Sessions_ids,
		},
		Code: 200,
		Err:  nil,
	}
}
