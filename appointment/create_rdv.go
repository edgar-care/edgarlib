package appointment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type CreateRdvResponse struct {
	Rdv    model.Rdv
	Doctor model.Doctor
	Code   int
	Err    error
}

func CreateRdv(patientId string, doctorId string, startDate int, endDate int, session_id string) CreateRdvResponse {

	var appointment_status model.AppointmentStatus = "WAITING_FOR_REVIEW"
	rdv, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         patientId,
		DoctorID:          doctorId,
		StartDate:         startDate,
		EndDate:           endDate,
		AppointmentStatus: appointment_status,
		SessionID:         session_id,
	})
	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	doctor, err := graphql.GetDoctorById(doctorId)
	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedDoctor, err := graphql.UpdateDoctor(context.Background(), gqlClient, doctorId, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, append(doctor.GetDoctorById.Rendez_vous_ids, rdv.CreateRdv.Id), doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, doctor.GetDoctorById.Device_connect, doctor.GetDoctorById.Double_auth_methods_id, doctor.GetDoctorById.Trust_devices)

	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 500, Err: errors.New("update failed" + err.Error())}
	}

	if patientId != "" {
		patient, err := graphql.GetPatientById(patientId)
		if err != nil {
			return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
		}
		_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, append(patient.GetPatientById.Rendez_vous_ids, rdv.CreateRdv.Id), patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id, patient.GetPatientById.Trust_devices)
		if err != nil {
			return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 500, Err: errors.New("unable to update patient")}
		}
	}

	return CreateRdvResponse{
		Rdv:    rdv,
		Doctor: updatedDoctor,
		Code:   200,
		Err:    nil,
	}
}
