package document

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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
	document, err := graphql.GetDocumentById(id)
	if err != nil {
		return GetDocumentByIdResponse{model.Document{}, 400, errors.New("id does not correspond to a document")}
	}

	ext := filepath.Ext(document.Name)
	if ext == "" {
		return GetDocumentByIdResponse{model.Document{}, 500, errors.New("invalid file extension")}
	}

	filename := document.ID + ext

	signedURL, err := GenerateURL("document-patient", filename, document.Name)
	if err != nil {
		return GetDocumentByIdResponse{model.Document{}, 500, fmt.Errorf("error generating signed URL: %v", err)}
	}

	res := model.Document{
		ID:           document.ID,
		OwnerID:      document.OwnerID,
		Name:         document.Name,
		DocumentType: document.DocumentType,
		Category:     document.Category,
		IsFavorite:   document.IsFavorite,
		DownloadURL:  signedURL,
		UploaderID:   document.UploaderID,
		CreatedAt:    document.CreatedAt,
	}

	return GetDocumentByIdResponse{res, 200, nil}
}

func GetDocuments(id string) GetDocumentsResponse {
	var res []model.Document

	documents, err := graphql.GetPatientDocument(id, nil)
	if err != nil {
		return GetDocumentsResponse{[]model.Document{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, document := range documents {
		ext := filepath.Ext(document.Name)
		if ext == "" {
			return GetDocumentsResponse{[]model.Document{}, 500, errors.New("invalid file extension")}
		}

		filename := document.ID + ext
		signedURL, err := GenerateURL("document-patient", filename, document.Name)
		if err != nil {
			return GetDocumentsResponse{[]model.Document{}, 500, fmt.Errorf("error generating signed URL: %v", err)}
		}

		res = append(res, model.Document{
			ID:           document.ID,
			OwnerID:      document.OwnerID,
			Name:         document.Name,
			DocumentType: document.DocumentType,
			Category:     document.Category,
			IsFavorite:   document.IsFavorite,
			DownloadURL:  signedURL,
			UploaderID:   document.UploaderID,
			CreatedAt:    document.CreatedAt,
		})
	}
	return GetDocumentsResponse{res, 200, nil}
}
