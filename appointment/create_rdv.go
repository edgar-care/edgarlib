package appointment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateRdvResponse struct {
	Rdv    model.Rdv
	Doctor model.Doctor
	Code   int
	Err    error
}

func CreateRdv(patientId string, doctorId string, startDate int, endDate int) CreateRdvResponse {
	gqlClient := graphql.CreateClient()

	rdv, err := graphql.CreateRdv(context.Background(), gqlClient, patientId, doctorId, startDate, endDate)
	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorId)
	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedDoctor, err := graphql.UpdateDoctor(context.Background(), gqlClient, doctorId, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, append(doctor.GetDoctorById.Rendez_vous_ids, rdv.CreateRdv.Id), doctor.GetDoctorById.Patient_ids)
	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	return CreateRdvResponse{
		Rdv: model.Rdv{
			ID:                rdv.CreateRdv.Id,
			DoctorID:          rdv.CreateRdv.Doctor_id,
			IDPatient:         rdv.CreateRdv.Id_patient,
			StartDate:         rdv.CreateRdv.Start_date,
			EndDate:           rdv.CreateRdv.End_date,
			CancelationReason: &rdv.CreateRdv.Cancelation_reason,
		},
		Doctor: model.Doctor{
			ID:            updatedDoctor.UpdateDoctor.Id,
			Email:         updatedDoctor.UpdateDoctor.Email,
			Password:      updatedDoctor.UpdateDoctor.Password,
			RendezVousIds: graphql.ConvertStringSliceToPointerSlice(updatedDoctor.UpdateDoctor.Rendez_vous_ids),
			PatientIds:    graphql.ConvertStringSliceToPointerSlice(updatedDoctor.UpdateDoctor.Patient_ids),
		},
		Code: 200,
		Err:  nil,
	}
}
