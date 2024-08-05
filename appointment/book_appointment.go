package appointment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
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

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, append(patient.GetPatientById.Rendez_vous_ids, appointmentId), patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id, patient.GetPatientById.Trust_devices)
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
		_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.GetDoctorById.Id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, append(doctor.GetDoctorById.Patient_ids, patientId), graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, doctor.GetDoctorById.Device_connect, doctor.GetDoctorById.Double_auth_methods_id, doctor.GetDoctorById.Trust_devices)
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
