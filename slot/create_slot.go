package slot

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateSlotInput struct {
	DoctorID  string `json:"doctor_id"`
	IdPatient string `json:"id_patient"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

type CreateSlotResponse struct {
	Rdv    model.Rdv
	Doctor model.Doctor
	Code   int
	Err    error
}

func CreateSlot(input CreateSlotInput) CreateSlotResponse {
	gqlClient := graphql.CreateClient()

	rdv, err := graphql.CreateRdv(context.Background(), gqlClient, input.IdPatient, input.DoctorID, input.StartDate, input.EndDate)
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, input.DoctorID)
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedDoctor, err := graphql.UpdateDoctor(context.Background(), gqlClient, input.DoctorID, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, append(doctor.GetDoctorById.Rendez_vous_ids, rdv.CreateRdv.Id), doctor.GetDoctorById.Patient_ids)
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	return CreateSlotResponse{
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
