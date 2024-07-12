package slot

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type DeleteSlotResponse struct {
	Deleted       bool
	UpdatedDoctor model.Doctor
	Code          int
	Err           error
}

func remElement(slice []string, element string) []string {
	var result []string
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
	gqlClient := graphql.CreateClient()

	check_id, err := graphql.GetSlotById(context.Background(), gqlClient, slotId)
	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a slot")}
	}
	if check_id.GetSlotById.Id_patient != "" {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("this slot is booked, you cannot delete it")}
	}

	deleted, err := graphql.DeleteRdv(context.Background(), gqlClient, slotId)
	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorId)
	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedDoctor, err := graphql.UpdateDoctor(context.Background(), gqlClient, doctorId, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, remElement(doctor.GetDoctorById.Rendez_vous_ids, slotId), doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, doctor.GetDoctorById.Device_connect, doctor.GetDoctorById.Double_auth_methods_id, doctor.GetDoctorById.Trust_devices)

	if err != nil {
		return DeleteSlotResponse{Deleted: false, UpdatedDoctor: model.Doctor{}, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	return DeleteSlotResponse{
		Deleted: deleted.DeleteRdv,
		UpdatedDoctor: model.Doctor{
			ID:            updatedDoctor.UpdateDoctor.Id,
			Email:         updatedDoctor.UpdateDoctor.Email,
			Password:      updatedDoctor.UpdateDoctor.Password,
			Name:          updatedDoctor.UpdateDoctor.Name,
			Firstname:     updatedDoctor.UpdateDoctor.Firstname,
			RendezVousIds: graphql.ConvertStringSliceToPointerSlice(updatedDoctor.UpdateDoctor.Rendez_vous_ids),
			PatientIds:    graphql.ConvertStringSliceToPointerSlice(updatedDoctor.UpdateDoctor.Patient_ids),
		},
		Code: 200,
		Err:  nil}
}
