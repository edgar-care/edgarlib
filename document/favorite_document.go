package document

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

func Updatefavorite(id string, favorite bool, ownerID string) UpdateDocumentResponse {

	document, err := graphql.GetDocumentById(id)
	if err != nil {
		return UpdateDocumentResponse{model.Document{}, 400, errors.New("id does not correspond to a document")}
	}

	if document.OwnerID != ownerID {
		return UpdateDocumentResponse{model.Document{}, 403, errors.New("you do not have permission to update this document")}
	}

	updatedDocument, err := graphql.UpdateDocument(id, model.UpdateDocumentInput{
		IsFavorite: &favorite,
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
