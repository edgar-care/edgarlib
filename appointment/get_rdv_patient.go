package appointment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type GetRdvPatientResponse struct {
	Rdv  model.Rdv
	Code int
	Err  error
}

func GetRdvPatient(appointmentId string, patientId string) GetRdvPatientResponse {
	rdv, err := graphql.GetRdvById(appointmentId)

	if err != nil {
		return GetRdvPatientResponse{model.Rdv{}, 400, errors.New("id does not correspond to an appointment")}
	}

	if rdv.IDPatient != patientId {
		return GetRdvPatientResponse{model.Rdv{}, 403, errors.New("unauthorized to access to this appointment")}
	}
	return GetRdvPatientResponse{rdv, 200, nil}
}
