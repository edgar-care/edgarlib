package medical_folder

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type DeleteMedicalInfoResponse struct {
	Deleted        bool
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func remElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}

func DeleteMedicalInfo(medicalId string, patientId string) DeleteMedicalInfoResponse {
	if medicalId == "" {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("medical id is required")}
	}

	gqlClient := graphql.CreateClient()
	if gqlClient == nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("failed to create GraphQL client")}
	}

	_, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, medicalId)
	if err != nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a medical folder")}
	}

	deleted, err := graphql.DeleteMedicalFolder(context.Background(), gqlClient, medicalId)
	if err != nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	updatedPatient, err := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, "", patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids)

	if err != nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	return DeleteMedicalInfoResponse{
		Deleted: deleted.DeleteMedicalFolder,
		UpdatedPatient: model.Patient{
			ID:                   updatedPatient.UpdatePatient.Id,
			Email:                updatedPatient.UpdatePatient.Email,
			Password:             updatedPatient.UpdatePatient.Password,
			MedicalInfoID:        &updatedPatient.UpdatePatient.Medical_info_id,
			DocumentIds:          graphql.ConvertStringSliceToPointerSlice(updatedPatient.UpdatePatient.Document_ids),
			RendezVousIds:        graphql.ConvertStringSliceToPointerSlice(updatedPatient.UpdatePatient.Rendez_vous_ids),
			TreatmentFollowUpIds: graphql.ConvertStringSliceToPointerSlice(updatedPatient.UpdatePatient.Treatment_follow_up_ids),
		},
		Code: 200,
		Err:  errors.New(""),
	}
}
