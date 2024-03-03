package slot

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetSlotByIdResponse struct {
	Slot model.Rdv
	Code int
	Err  error
}

type GetSlotsResponse struct {
	Slots []model.Rdv
	Code  int
	Err   error
}

func GetSlotById(id string, doctorId string) GetSlotByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.Rdv

	slot, err := graphql.GetRdvById(context.Background(), gqlClient, id)
	if err != nil {
		return GetSlotByIdResponse{model.Rdv{}, 400, errors.New("id does not correspond to a slot")}
	}
	if slot.GetRdvById.Doctor_id != doctorId {
		return GetSlotByIdResponse{model.Rdv{}, 403, errors.New("you cannot access to this appointment")}
	}
	res = model.Rdv{
		ID:        slot.GetRdvById.Id,
		DoctorID:  slot.GetRdvById.Doctor_id,
		IDPatient: slot.GetRdvById.Id_patient,
		StartDate: slot.GetRdvById.Start_date,
		EndDate:   slot.GetRdvById.End_date,
	}
	return GetSlotByIdResponse{res, 200, nil}
}

func GetSlots(doctorId string) GetSlotsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Rdv

	_, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetSlotsResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a doctor")}
	}

	slots, err := graphql.GetDoctorRdv(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetSlotsResponse{[]model.Rdv{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, slot := range slots.GetDoctorRdv {
		temp := slot.Cancelation_reason
		res = append(res, model.Rdv{
			ID:                slot.Id,
			DoctorID:          slot.Doctor_id,
			IDPatient:         slot.Id_patient,
			StartDate:         slot.Start_date,
			EndDate:           slot.End_date,
			CancelationReason: &temp,
		})
	}
	return GetSlotsResponse{res, 200, nil}
}
