package appointment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	rdv, err := graphql.GetRdvById(context.Background(), gqlClient, rdvId)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	var appointment_status = graphql.AppointmentStatusCanceled
	_, err = graphql.UpdateRdv(context.Background(), gqlClient, rdv.GetRdvById.Id, rdv.GetRdvById.Id_patient, rdv.GetRdvById.Doctor_id, rdv.GetRdvById.Start_date, rdv.GetRdvById.End_date, rdv.GetRdvById.Cancelation_reason, appointment_status, rdv.GetRdvById.Session_id, rdv.GetRdvById.Health_method)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("could not update appointment")}
	}

	var new_appointment_status graphql.AppointmentStatus = "OPENED"
	new_slot, err := graphql.CreateRdv(context.Background(), gqlClient, "", rdv.GetRdvById.Doctor_id, rdv.GetRdvById.Start_date, rdv.GetRdvById.End_date, new_appointment_status, "")
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, rdv.GetRdvById.Doctor_id)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id)
	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, rdv.GetRdvById.Doctor_id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, append(doctor.GetDoctorById.Rendez_vous_ids, new_slot.CreateRdv.Id), doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}
	return DeleteRdvResponse{
		UpdatedPatient: model.Patient{
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

func removeElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
