package document

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

func Updatefavorite(newDocumentInfo CreateDocumentInput, id string, favorite bool) UpdateDocumentResponse {
	gqlClient := graphql.CreateClient()

	updatedDocument, err := graphql.UpdateDocument(context.Background(), gqlClient, id, newDocumentInfo.Name, favorite)
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
