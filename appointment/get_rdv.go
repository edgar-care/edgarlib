package appointment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetRdvResponse struct {
	Rdv  []model.Rdv
	Code int
	Err  error
}

func GetRdv(patientId string) GetRdvResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Rdv
	rdv, err := graphql.GetPatientRdv(context.Background(), gqlClient, patientId)

	if err != nil {
		return GetRdvResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a patient")}
	}
	for _, appointment := range rdv.GetPatientRdv {
		temp := appointment.Cancelation_reason
		res = append(res, model.Rdv{
			ID:                appointment.Id,
			DoctorID:          appointment.Doctor_id,
			IDPatient:         appointment.Id_patient,
			StartDate:         appointment.Start_date,
			EndDate:           appointment.End_date,
			CancelationReason: &temp,
		})
	}
	return GetRdvResponse{res, 200, nil}
}
