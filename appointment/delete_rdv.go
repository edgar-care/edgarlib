package appointment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type DeleteRdvResponse struct {
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func DeleteRdv(rdvId string, patientId string) DeleteRdvResponse {
	if rdvId == "" {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("rdv id is required")}
	}

	rdv, err := graphql.GetRdvById(rdvId)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	var appointment_status = model.AppointmentStatusCanceled
	_, err = graphql.UpdateRdv(rdv.ID, model.UpdateRdvInput{
		AppointmentStatus: &appointment_status,
	})
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("could not update appointment")}
	}

	var new_appointment_status model.AppointmentStatus = "OPENED"
	new_slot, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          rdv.DoctorID,
		StartDate:         rdv.StartDate,
		EndDate:           rdv.EndDate,
		AppointmentStatus: new_appointment_status,
		SessionID:         "",
	})
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	doctor, err := graphql.GetDoctorById(rdv.DoctorID)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id, patient.GetPatientById.Trust_devices)
	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, rdv.GetRdvById.Doctor_id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, append(doctor.GetDoctorById.Rendez_vous_ids, new_slot.CreateRdv.Id), doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, doctor.GetDoctorById.Device_connect, doctor.GetDoctorById.Double_auth_methods_id, doctor.GetDoctorById.Trust_devices)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}
	return DeleteRdvResponse{
		UpdatedPatient: patient,
		Code:           200,
		Err:            nil,
	}
}

func removeElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
