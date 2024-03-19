package appointment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type UpdateDoctorAppointmentStruct struct {
	Rdv     model.Rdv
	Patient model.Patient
	Code    int
	Err     error
}

func UpdateDoctorAppointment(newAppointmentId string, appointmentId string) UpdateDoctorAppointmentStruct {
	gqlClient := graphql.CreateClient()
	if appointmentId == "" {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetRdvById(context.Background(), gqlClient, newAppointmentId)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	if appointment.GetRdvById.Id_patient != "" {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment is booked, cannot be edited")}
	}

	oldAppointment, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentId)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	updatedRdv, err := graphql.UpdateRdv(context.Background(), gqlClient, newAppointmentId, oldAppointment.GetRdvById.Id_patient, appointment.GetRdvById.Doctor_id, appointment.GetRdvById.Start_date, appointment.GetRdvById.End_date, appointment.GetRdvById.Cancelation_reason, appointment.GetRdvById.Appointment_status, appointment.GetRdvById.Session_id)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, oldAppointment.GetRdvById.Id_patient)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
	}
	_, err = graphql.UpdateRdv(context.Background(), gqlClient, appointmentId, "", oldAppointment.GetRdvById.Doctor_id, oldAppointment.GetRdvById.Start_date, oldAppointment.GetRdvById.End_date, oldAppointment.GetRdvById.Cancelation_reason, oldAppointment.GetRdvById.Appointment_status, oldAppointment.GetRdvById.Session_id)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, oldAppointment.GetRdvById.Id_patient, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, append(removeElement(patient.GetPatientById.Rendez_vous_ids, appointmentId), newAppointmentId), patient.GetPatientById.Document_ids)
	if err != nil {
		return UpdateDoctorAppointmentStruct{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	return UpdateDoctorAppointmentStruct{
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
