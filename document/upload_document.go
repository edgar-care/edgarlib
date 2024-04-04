package document

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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

func CreateDocument(newdoc UploadDocumentInput, id string) CreateDocumentResponse {
	gqlClient := graphql.CreateClient()

	document, err := graphql.CreateDocument(context.Background(), gqlClient, id, newdoc.Name, newdoc.DocumentType, newdoc.Category, newdoc.IsFavorite, newdoc.DownloadURL)
	if err != nil {
		return CreateDocumentResponse{Document: model.Document{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err != nil {
		return CreateDocumentResponse{Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	_, _ = graphql.UpdatePatient(context.Background(), gqlClient, id, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, append(patient.GetPatientById.Document_ids, document.CreateDocument.Id), patient.GetPatientById.Treatment_follow_up_ids)

	return CreateDocumentResponse{
		Document: model.Document{
			ID:           document.CreateDocument.Id,
			OwnerID:      document.CreateDocument.Owner_id,
			Name:         document.CreateDocument.Name,
			DocumentType: model.DocumentType(document.CreateDocument.Category),
			Category:     model.Category(document.CreateDocument.Category),
			IsFavorite:   document.CreateDocument.Is_favorite,
			DownloadURL:  document.CreateDocument.Download_url,
		},
		Code: 201,
		Err:  nil,
	}
}
