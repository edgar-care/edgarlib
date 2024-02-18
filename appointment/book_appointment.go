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

func BookAppointment(appointmentId string, patientId string) BookAppointmentResponse {
	gqlClient := graphql.CreateClient()
	if appointmentId == "" {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment id is required")}
	}

	appointment, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentId)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	if appointment.GetRdvById.Id_patient != "" {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("appointment is already booked")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
	}

	updatedRdv, err := graphql.UpdateRdv(context.Background(), gqlClient, appointmentId, patientId, appointment.GetRdvById.Doctor_id, appointment.GetRdvById.Start_date, appointment.GetRdvById.End_date, appointment.GetRdvById.Cancelation_reason)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update appointment")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, append(patient.GetPatientById.Rendez_vous_ids, appointmentId), patient.GetPatientById.Document_ids)
	if err != nil {
		return BookAppointmentResponse{Rdv: model.Rdv{}, Patient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	return BookAppointmentResponse{
		Rdv: model.Rdv{
			ID:                updatedRdv.UpdateRdv.Id,
			DoctorID:          updatedRdv.UpdateRdv.Doctor_id,
			IDPatient:         updatedRdv.UpdateRdv.Id_patient,
			StartDate:         updatedRdv.UpdateRdv.Start_date,
			EndDate:           updatedRdv.UpdateRdv.End_date,
			CancelationReason: &updatedRdv.UpdateRdv.Cancelation_reason,
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
