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
	_, err = graphql.UpdateRdv(context.Background(), gqlClient, rdv.GetRdvById.Id, "", rdv.GetRdvById.Doctor_id, rdv.GetRdvById.Start_date, rdv.GetRdvById.End_date, rdv.GetRdvById.Cancelation_reason, appointment_status, rdv.GetRdvById.Session_id)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("could not update appointment")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, removeElement(patient.GetPatientById.Rendez_vous_ids, rdvId), patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids)
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
