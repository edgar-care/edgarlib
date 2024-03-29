package document

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetDocumentByIdResponse struct {
	Document model.Document
	Code     int
	Err      error
}

type GetDocumentsResponse struct {
	Documents []model.Document
	Code      int
	Err       error
}

func GetDocument(id string) GetDocumentByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.Document

	document, err := graphql.GetDocumentById(context.Background(), gqlClient, id)
	if err != nil {
		return GetDocumentByIdResponse{model.Document{}, 400, errors.New("id does not correspond to a document")}
	}
	res = model.Document{
		ID:           document.GetDocumentById.Id,
		OwnerID:      document.GetDocumentById.Owner_id,
		Name:         document.GetDocumentById.Name,
		DocumentType: model.DocumentType(document.GetDocumentById.Document_type),
		Category:     model.Category(document.GetDocumentById.Category),
		IsFavorite:   document.GetDocumentById.Is_favorite,
		DownloadURL:  document.GetDocumentById.Download_url,
	}
	return GetDocumentByIdResponse{res, 200, nil}
}

func GetDocuments(id string) GetDocumentsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Document

	documents, err := graphql.GetPatientDocument(context.Background(), gqlClient, id)
	if err != nil {
		return GetDocumentsResponse{[]model.Document{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, document := range documents.GetPatientDocument {
		res = append(res, model.Document{
			ID:           document.Id,
			OwnerID:      document.Owner_id,
			Name:         document.Name,
			DocumentType: model.DocumentType(document.Document_type),
			Category:     model.Category(document.Category),
			IsFavorite:   document.Is_favorite,
			DownloadURL:  document.Download_url,
		})
	}
	return GetDocumentsResponse{res, 200, nil}
}
