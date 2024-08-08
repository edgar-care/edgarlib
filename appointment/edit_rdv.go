package appointment

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type EditRdvResponse struct {
	Rdv     model.Rdv
	Patient model.Patient
	Code    int
	Err     error
}

func EditRdv(newAppointmentId string, appointmentId string, patientId string) EditRdvResponse {
	if appointmentId == "" {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetSlotById(newAppointmentId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	if appointment.IDPatient != "" {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment is booked, cannot be edited")}
	}

	oldAppointment, err := graphql.GetRdvById(appointmentId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
	}

	updatedRdv, err := graphql.UpdateRdv(newAppointmentId, model.UpdateRdvInput{
		IDPatient:         &patientId,
		DoctorID:          nil,
		StartDate:         nil,
		EndDate:           nil,
		CancelationReason: nil,
		AppointmentStatus: &oldAppointment.AppointmentStatus,
		SessionID:         &oldAppointment.SessionID,
		HealthMethod:      oldAppointment.HealthMethod,
	})
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	var appointment_status = model.AppointmentStatusOpened
	empty := ""
	_, err = graphql.UpdateRdv(appointmentId, model.UpdateRdvInput{
		IDPatient:         &empty,
		DoctorID:          nil,
		StartDate:         nil,
		EndDate:           nil,
		CancelationReason: nil,
		AppointmentStatus: &appointment_status,
		SessionID:         &empty,
		HealthMethod:      &empty,
	})
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	updatedPatient, err := graphql.UpdatePatient(patientId, model.UpdatePatientInput{
		RendezVousIds: append(removeElement(patient.RendezVousIds, &appointmentId), &newAppointmentId),
	})
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	return EditRdvResponse{
		Rdv:     updatedRdv,
		Patient: updatedPatient,
		Code:    200,
		Err:     nil,
	}
}
