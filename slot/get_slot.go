package slot

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
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
	slot, err := graphql.GetSlotById(id)
	if err != nil {
		return GetSlotByIdResponse{model.Rdv{}, 400, errors.New("id does not correspond to a slot")}
	}
	if slot.DoctorID != doctorId {
		return GetSlotByIdResponse{model.Rdv{}, 403, errors.New("you cannot access to this appointment")}
	}
	return GetSlotByIdResponse{slot, 200, nil}
}

func GetSlots(doctorId string) GetSlotsResponse {
	_, err := graphql.GetDoctorById(doctorId)
	if err != nil {
		return GetSlotsResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a doctor")}
	}

	slots, err := graphql.GetSlots(doctorId, nil)
	if err != nil {
		return GetSlotsResponse{[]model.Rdv{}, 400, errors.New("invalid input: " + err.Error())}
	}

	return GetSlotsResponse{slots, 200, nil}
}
