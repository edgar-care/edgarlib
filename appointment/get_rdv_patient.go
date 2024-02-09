package appointment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetRdvPatientResponse struct {
	Rdv  model.Rdv
	Code int
	Err  error
}

func GetRdvPatient(appointmentId string, patientId string) GetRdvPatientResponse {
	gqlClient := graphql.CreateClient()
	rdv, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentId)

	if err != nil {
		return GetRdvPatientResponse{model.Rdv{}, 400, errors.New("id does not correspond to an appointment")}
	}

	if rdv.GetRdvById.Id_patient != patientId {
		return GetRdvPatientResponse{model.Rdv{}, 403, errors.New("unauthorized to access to this appointment")}
	}
	return GetRdvPatientResponse{model.Rdv{
		ID:                rdv.GetRdvById.Id,
		DoctorID:          rdv.GetRdvById.Doctor_id,
		IDPatient:         rdv.GetRdvById.Id_patient,
		StartDate:         rdv.GetRdvById.Start_date,
		EndDate:           rdv.GetRdvById.End_date,
		CancelationReason: &rdv.GetRdvById.Cancelation_reason,
	}, 200, nil}
}
