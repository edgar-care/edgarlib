package document

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type CreateDocumentInput struct {
	ID           string `json:"id"`
	OwnerID      string `json:"owner_id"`
	Name         string `json:"name"`
	DocumentType string `json:"document_type"`
	Category     string `json:"category"`
	IsFavorite   bool   `json:"is_favorite"`
	DownloadURL  string `json:"download_url"`
}

type UpdateDocumentResponse struct {
	Document model.Document
	Code     int
	Err      error
}

func UpdateDocument(newDocumentInfo CreateDocumentInput, id string) UpdateDocumentResponse {
	updatedDocument, err := graphql.UpdateDocument(id, model.UpdateDocumentInput{
		Name:       &newDocumentInfo.Name,
		IsFavorite: &newDocumentInfo.IsFavorite,
	})
	if err != nil {
		return UpdateDocumentResponse{Document: model.Document{}, Code: 500, Err: errors.New("unable to update document")}
	}

	return UpdateDocumentResponse{
		Document: updatedDocument,
		Code:     201,
		Err:      nil,
	}
}
