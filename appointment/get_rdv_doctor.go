package appointment

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type GetRdvDoctorResponse struct {
	Rdv  []model.Rdv
	Code int
	Err  error
}

func GetRdvDoctor(doctorId string) GetRdvDoctorResponse {
	var res []model.Rdv

	rdv, err := graphql.GetDoctorRdv(doctorId, nil)
	if err != nil {
		return GetRdvDoctorResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a doctor")}
	}

	for _, appointment := range rdv {
		res = append(res, appointment)
	}

	return GetRdvDoctorResponse{res, 200, nil}
}
