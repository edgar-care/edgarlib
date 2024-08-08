package appointment

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type UpdateDoctorAppointmentStruct struct {
	Rdv     model.Rdv
	Patient model.Patient
	Code    int
	Err     error
}

func UpdateDoctorAppointment(newAppointmentId string, appointmentId string) UpdateDoctorAppointmentStruct {
	if appointmentId == "" {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetSlotById(newAppointmentId)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	if appointment.IDPatient != "" {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment is booked, cannot be edited")}
	}

	oldAppointment, err := graphql.GetRdvById(appointmentId)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	updatedRdv, err := graphql.UpdateRdv(newAppointmentId, model.UpdateRdvInput{
		IDPatient:         &oldAppointment.IDPatient,
		AppointmentStatus: &oldAppointment.AppointmentStatus,
		SessionID:         &oldAppointment.SessionID,
		HealthMethod:      oldAppointment.HealthMethod,
	})

	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	patient, err := graphql.GetPatientById(oldAppointment.IDPatient)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
	}

	var appointment_status = model.AppointmentStatusOpened
	empty := ""
	_, err = graphql.UpdateRdv(appointmentId, model.UpdateRdvInput{
		IDPatient:         &empty,
		CancelationReason: &empty,
		AppointmentStatus: &appointment_status,
		SessionID:         &empty,
		HealthMethod:      &empty,
	})

	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	updatedPatient, err := graphql.UpdatePatient(oldAppointment.IDPatient, model.UpdatePatientInput{
		RendezVousIds: append(removeElement(patient.RendezVousIds, &appointmentId), &newAppointmentId),
	})
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	return UpdateDoctorAppointmentStruct{
		Rdv:     updatedRdv,
		Patient: updatedPatient,
		Code:    200,
		Err:     nil,
	}
}
