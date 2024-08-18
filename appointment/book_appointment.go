package appointment

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type BookAppointmentResponse struct {
	Rdv     model.Rdv
	Patient model.Patient
	Code    int
	Err     error
}

func BookAppointment(appointmentId string, patientId string, session_id string) BookAppointmentResponse {
	if appointmentId == "" {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	if session_id == "" {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("session id is required")}
	}

	appointment, err := graphql.GetSlotById(appointmentId)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	if appointment.IDPatient != "" {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment is already booked")}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
	}
	var appointment_status = model.AppointmentStatusWaitingForReview

	updatedRdv, err := graphql.UpdateRdv(appointmentId, model.UpdateRdvInput{
		IDPatient:         &patientId,
		AppointmentStatus: &appointment_status,
		SessionID:         &session_id,
	})
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	updatedPatient, err := graphql.UpdatePatient(patientId, model.UpdatePatientInput{
		RendezVousIds: append(patient.RendezVousIds, &appointmentId),
	})
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	doctor, err := graphql.GetDoctorById(appointment.DoctorID)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	patientFound := false
	for _, patientID := range doctor.PatientIds {
		if patientID == &patientId {
			patientFound = true
			break
		}
	}

	if !patientFound {
		_, err = graphql.UpdateDoctor(doctor.ID, model.UpdateDoctorInput{
			PatientIds: append(doctor.PatientIds, &patientId),
		})
		if err != nil {
			return BookAppointmentResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("update failed" + err.Error())}
		}
	}

	return BookAppointmentResponse{
		Rdv:     updatedRdv,
		Patient: updatedPatient,
		Code:    200,
		Err:     nil,
	}
}
