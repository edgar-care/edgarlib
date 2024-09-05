package ordonnance

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/document"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type GetOrdonnanceByIdResponse struct {
	Ordonnance model.Ordonnance
	Url        string
	Code       int
	Err        error
}

type OrdonnanceWithURL struct {
	model.Ordonnance
	URL string
}

type GetOrdonnancesResponse struct {
	Ordonnance []OrdonnanceWithURL
	Code       int
	Err        error
}

func GetOrdonnancebyID(id string) GetOrdonnanceByIdResponse {
	ordonnance, err := graphql.GetOrdonnanceById(id)
	if err != nil {
		return GetOrdonnanceByIdResponse{model.Ordonnance{}, "", 400, errors.New("id does not correspond to a medicament")}
	}
	filename := ordonnance.ID + ".pdf"

	url, err := document.GenerateURL("doctor-ordonnance", filename, "Ordonnance.pdf")
	if err != nil {
		return GetOrdonnanceByIdResponse{model.Ordonnance{}, "", 500, errors.New("error generating signed URL")}
	}

	return GetOrdonnanceByIdResponse{ordonnance, url, 200, nil}
}

func GetOrdonnancesDoctor(ownerID string) GetOrdonnancesResponse {
	ordonnances, err := graphql.GetOrdonnanceByDoctorId(ownerID, nil)
	if err != nil {
		return GetOrdonnancesResponse{[]OrdonnanceWithURL{}, 400, errors.New("invalid input: " + err.Error())}
	}

	ordonnancesWithURL := make([]OrdonnanceWithURL, len(ordonnances))
	for i, ordonnance := range ordonnances {
		filename := ordonnance.ID + ".pdf"
		url, err := document.GenerateURL("doctor-ordonnance", filename, "Ordonnance.pdf")
		if err != nil {
			return GetOrdonnancesResponse{[]OrdonnanceWithURL{}, 500, errors.New("error generating signed URL for ordonnance ID: " + ordonnance.ID)}
		}
		ordonnancesWithURL[i] = OrdonnanceWithURL{
			Ordonnance: ordonnance,
			URL:        url,
		}
	}

	return GetOrdonnancesResponse{ordonnancesWithURL, 200, nil}
}
