package document

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

func Updatefavorite(id string, favorite bool) UpdateDocumentResponse {
	gqlClient := graphql.CreateClient()

	document, err := graphql.GetDocumentById(context.Background(), gqlClient, id)
	if err != nil {
		return UpdateDocumentResponse{model.Document{}, 400, errors.New("id does not correspond to a document")}
	}

	updatedDocument, err := graphql.UpdateDocument(context.Background(), gqlClient, id, document.GetDocumentById.Name, favorite)
	if err != nil {
		return UpdateDocumentResponse{Document: model.Document{}, Code: 500, Err: errors.New("unable to update document")}
	}

	return UpdateDocumentResponse{
		Document: model.Document{
			ID:         updatedDocument.UpdateDocument.Id,
			OwnerID:    updatedDocument.UpdateDocument.Owner_id,
			Name:       updatedDocument.UpdateDocument.Name,
			IsFavorite: updatedDocument.UpdateDocument.Is_favorite,
		},
		Code: 201,
		Err:  nil,
	}
}
