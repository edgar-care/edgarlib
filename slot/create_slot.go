package slot

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	var appointment_status graphql.AppointmentStatus = "OPENED"
	rdv, err := graphql.CreateRdv(context.Background(), gqlClient, "", doctorID, input.StartDate, input.EndDate, appointment_status, "")
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorID)
	if err != nil {
		return CreateSlotResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctorID, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, append(doctor.GetDoctorById.Rendez_vous_ids, rdv.CreateRdv.Id), doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, doctor.GetDoctorById.Device_connect, doctor.GetDoctorById.Double_auth_methods_id, doctor.GetDoctorById.Trust_devices)
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
			AppointmentStatus: model.AppointmentStatus(rdv.CreateRdv.Appointment_status),
			SessionID:         rdv.CreateRdv.Session_id,
		},
		Code: 200,
		Err:  nil,
	}
}
