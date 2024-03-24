package appointment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type EditRdvResponse struct {
	Rdv     model.Rdv
	Patient model.Patient
	Code    int
	Err     error
}

func EditRdv(newAppointmentId string, appointmentId string, patientId string) EditRdvResponse {
	gqlClient := graphql.CreateClient()
	if appointmentId == "" {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetSlotById(context.Background(), gqlClient, newAppointmentId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	if appointment.GetSlotById.Id_patient != "" {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment is booked, cannot be edited")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
	}

	updatedRdv, err := graphql.UpdateRdv(context.Background(), gqlClient, newAppointmentId, patientId, appointment.GetSlotById.Doctor_id, appointment.GetSlotById.Start_date, appointment.GetSlotById.End_date, appointment.GetSlotById.Cancelation_reason, appointment.GetSlotById.Appointment_status, appointment.GetSlotById.Session_id)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	oldAppointment, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentId)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	_, err = graphql.UpdateRdv(context.Background(), gqlClient, appointmentId, "", oldAppointment.GetRdvById.Doctor_id, oldAppointment.GetRdvById.Start_date, oldAppointment.GetRdvById.End_date, oldAppointment.GetRdvById.Cancelation_reason, oldAppointment.GetRdvById.Appointment_status, oldAppointment.GetRdvById.Session_id)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, append(removeElement(patient.GetPatientById.Rendez_vous_ids, appointmentId), newAppointmentId), patient.GetPatientById.Document_ids)
	if err != nil {
		return EditRdvResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
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
