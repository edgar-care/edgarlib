package document

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type DeleteDocumentResponse struct {
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func DeleteDocument(docId string, patientId string) DeleteDocumentResponse {
	if docId == "" {
		return DeleteDocumentResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("document id is required")}
	}
	gqlClient := graphql.CreateClient()

	_, err := graphql.GetDocumentById(context.Background(), gqlClient, docId)
	if err != nil {
		return DeleteDocumentResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a document")}
	}

	_, err = graphql.DeleteDocument(context.Background(), gqlClient, docId)
	if err != nil {
		return DeleteDocumentResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a document, unable to delete")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return DeleteDocumentResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	updatedPatient, _ := graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, removeElement(patient.GetPatientById.Document_ids, docId), patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids)
	return DeleteDocumentResponse{
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
