package appointment

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type GetRdvResponse struct {
	Rdv  []model.Rdv
	Code int
	Err  error
}

func GetRdv(patientId string) GetRdvResponse {
	var res []model.Rdv
	rdv, err := graphql.GetPatientRdv(patientId, nil)

	if err != nil {
		return GetRdvResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a patient")}
	}
	for _, appointment := range rdv {
		res = append(res, appointment)
	}
	return GetRdvResponse{res, 200, nil}
}
