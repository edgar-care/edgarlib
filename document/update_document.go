package document

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	updatedDocument, err := graphql.UpdateDocument(context.Background(), gqlClient, id, newDocumentInfo.Name, newDocumentInfo.IsFavorite)
	if err != nil {
		return UpdateDocumentResponse{Document: model.Document{}, Code: 500, Err: errors.New("unable to update document")}
	}

	return UpdateDocumentResponse{
		Document: model.Document{
			ID:           updatedDocument.UpdateDocument.Id,
			OwnerID:      updatedDocument.UpdateDocument.Owner_id,
			Name:         updatedDocument.UpdateDocument.Name,
			DocumentType: model.DocumentType(updatedDocument.UpdateDocument.Document_type),
			Category:     model.Category(updatedDocument.UpdateDocument.Category),
			IsFavorite:   updatedDocument.UpdateDocument.Is_favorite,
			DownloadURL:  updatedDocument.UpdateDocument.Download_url,
		},
		Code: 201,
		Err:  nil,
	}
}
