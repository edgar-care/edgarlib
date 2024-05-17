package document

import (
	"context"
	"errors"
	"fmt"

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

	signedURL, err := generateSignedURL("document-patient", document.GetDocumentById.Name)
	if err != nil {
		return GetDocumentByIdResponse{model.Document{}, 500, fmt.Errorf("error generating signed URL: %v", err)}
	}

	res = model.Document{
		ID:           document.GetDocumentById.Id,
		OwnerID:      document.GetDocumentById.Owner_id,
		Name:         document.GetDocumentById.Name,
		DocumentType: model.DocumentType(document.GetDocumentById.Document_type),
		Category:     model.Category(document.GetDocumentById.Category),
		IsFavorite:   document.GetDocumentById.Is_favorite,
		DownloadURL:  signedURL,
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
		signedURL, err := generateSignedURL("document-patient", document.Name)
		if err != nil {
			return GetDocumentsResponse{[]model.Document{}, 500, fmt.Errorf("error generating signed URL: %v", err)}
		}

		res = append(res, model.Document{
			ID:           document.Id,
			OwnerID:      document.Owner_id,
			Name:         document.Name,
			DocumentType: model.DocumentType(document.Document_type),
			Category:     model.Category(document.Category),
			IsFavorite:   document.Is_favorite,
			DownloadURL:  signedURL,
		})
	}
	return GetDocumentsResponse{res, 200, nil}
}
