package appointment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	if session_id == "" {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("session id is required")}
	}

	appointment, err := graphql.GetSlotById(context.Background(), gqlClient, appointmentId)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	if appointment.GetSlotById.Id_patient != "" {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment is already booked")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
	}
	var appointment_status = graphql.AppointmentStatusWaitingForReview

	updatedRdv, err := graphql.UpdateRdv(context.Background(), gqlClient, appointmentId, patientId, appointment.GetSlotById.Doctor_id, appointment.GetSlotById.Start_date, appointment.GetSlotById.End_date, appointment.GetSlotById.Cancelation_reason, appointment_status, session_id)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, append(patient.GetPatientById.Rendez_vous_ids, appointmentId), patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, appointment.GetSlotById.Doctor_id)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	patientFound := false
	for _, patientID := range doctor.GetDoctorById.Patient_ids {
		if patientID == patientId {
			patientFound = true
			break
		}
	}

	if !patientFound {
		_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.GetDoctorById.Id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, append(doctor.GetDoctorById.Patient_ids, patientId), graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids)
		if err != nil {
			return BookAppointmentResponse{Rdv: model.Rdv{}, Code: 400, Err: errors.New("update failed" + err.Error())}
		}
	}

	return BookAppointmentResponse{
		Rdv: model.Rdv{
			ID:                updatedRdv.UpdateRdv.Id,
			DoctorID:          updatedRdv.UpdateRdv.Doctor_id,
			IDPatient:         updatedRdv.UpdateRdv.Id_patient,
			StartDate:         updatedRdv.UpdateRdv.Start_date,
			EndDate:           updatedRdv.UpdateRdv.End_date,
			CancelationReason: &updatedRdv.UpdateRdv.Cancelation_reason,
			AppointmentStatus: model.AppointmentStatus(updatedRdv.UpdateRdv.Appointment_status),
			SessionID:         updatedRdv.UpdateRdv.Session_id,
		},
		Patient: model.Patient{
			ID:            updatedPatient.UpdatePatient.Id,
			Email:         updatedPatient.UpdatePatient.Email,
			Password:      updatedPatient.UpdatePatient.Password,
			RendezVousIds: graphql.ConvertStringSliceToPointerSlice(updatedPatient.UpdatePatient.Rendez_vous_ids),
			MedicalInfoID: &updatedPatient.UpdatePatient.Medical_info_id,
			DocumentIds:   graphql.ConvertStringSliceToPointerSlice(updatedPatient.UpdatePatient.Document_ids),
		},
		Code: 200,
		Err:  nil,
	}
}
