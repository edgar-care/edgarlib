package slot

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type CreateSlotInput struct {
	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`
}

type CreateSlotResponse struct {
	Rdv    model.Rdv
	Doctor model.Doctor
	Code   int
	Err    error
}

func CreateSlot(input CreateSlotInput, doctorID string) CreateSlotResponse {

	var appointment_status model.AppointmentStatus = "OPENED"

	rdv, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          doctorID,
		StartDate:         input.StartDate,
		EndDate:           input.EndDate,
		AppointmentStatus: appointment_status,
		SessionID:         "",
	})
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	doctor, err := graphql.GetDoctorById(doctorID)
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdateDoctor(doctorID, model.UpdateDoctorInput{RendezVousIds: append(doctor.RendezVousIds, &rdv.ID)})
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	return CreateSlotResponse{
		Rdv:  rdv,
		Code: 201,
		Err:  nil,
	}
}
