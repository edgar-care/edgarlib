package document

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type UploadDocumentInput struct {
	ID           string `json:"id"`
	OwnerID      string `json:"owner_id"`
	Name         string `json:"name"`
	DocumentType string `json:"document_type"`
	Category     string `json:"category"`
	IsFavorite   bool   `json:"is_favorite"`
	DownloadURL  string `json:"download_url"`
}

type CreateDocumentResponse struct {
	Document       model.Document
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func CreateDocument(newdoc UploadDocumentInput, ownerId string, uploaderID string) CreateDocumentResponse {
	document, err := graphql.CreateDocument(model.CreateDocumentInput{
		OwnerID:      ownerId,
		Name:         newdoc.Name,
		DocumentType: newdoc.DocumentType,
		Category:     newdoc.Category,
		IsFavorite:   newdoc.IsFavorite,
		DownloadURL:  newdoc.DownloadURL,
		UploaderID:   uploaderID,
	})
	if err != nil {
		return CreateDocumentResponse{Document: model.Document{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	patient, err := graphql.GetPatientById(ownerId)
	if err != nil {
		return CreateDocumentResponse{Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	_, _ = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{DocumentIds: append(patient.DocumentIds, &document.ID)})

	return CreateDocumentResponse{
		Document: document,
		Code:     201,
		Err:      nil,
	}
}
