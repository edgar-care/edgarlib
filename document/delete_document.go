package document

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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

	_, err := graphql.GetDocumentById(docId)
	if err != nil {
		return DeleteDocumentResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a document")}
	}

	_, err = graphql.DeleteDocument(docId)
	if err != nil {
		return DeleteDocumentResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a document, unable to delete")}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return DeleteDocumentResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	updatedPatient, _ := graphql.UpdatePatient(patientId, model.UpdatePatientInput{
		DocumentIds: removeElement(patient.DocumentIds, &docId),
	})
	return DeleteDocumentResponse{
		UpdatedPatient: updatedPatient,
		Code:           200,
		Err:            nil,
	}
}

func removeElement(slice []*string, element *string) []*string {
	var result []*string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
