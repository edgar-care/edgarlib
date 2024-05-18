package document

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

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

	ext := filepath.Ext(document.GetDocumentById.Name)
	if ext == "" {
		return GetDocumentByIdResponse{model.Document{}, 500, errors.New("invalid file extension")}
	}

	filename := document.GetDocumentById.Id + ext

	signedURL, err := generateURL("document-patient", filename, document.GetDocumentById.Name)
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
		ext := filepath.Ext(document.Name)
		if ext == "" {
			return GetDocumentsResponse{[]model.Document{}, 500, errors.New("invalid file extension")}
		}

		filename := document.Id + ext
		signedURL, err := generateURL("document-patient", filename, document.Name)
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
