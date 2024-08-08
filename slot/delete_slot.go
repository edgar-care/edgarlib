package slot

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type DeleteSlotResponse struct {
	Deleted       bool
	UpdatedDoctor model.Doctor
	Code          int
	Err           error
}

func remElement(slice []*string, element *string) []*string {
	var result []*string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}

func DeleteSlot(slotId string, doctorId string) DeleteSlotResponse {
	if slotId == "" {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("slot id is required")}
	}

	check_id, err := graphql.GetSlotById(slotId)
	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a slot")}
	}
	if check_id.IDPatient != "" {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("this slot is booked, you cannot delete it")}
	}

	deleted, err := graphql.DeleteRdv(slotId)
	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	doctor, err := graphql.GetDoctorById(doctorId)
	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedDoctor, err := graphql.UpdateDoctor(doctorId, model.UpdateDoctorInput{RendezVousIds: remElement(doctor.RendezVousIds, &slotId)})

	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	return DeleteSlotResponse{
		Deleted:       deleted,
		UpdatedDoctor: updatedDoctor,
		Code:          200,
		Err:           nil}
}
