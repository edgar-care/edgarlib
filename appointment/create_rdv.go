package appointment

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type CreateRdvResponse struct {
	Rdv    model.Rdv
	Doctor model.Doctor
	Code   int
	Err    error
}

func CreateRdv(patientId string, doctorId string, startDate int, endDate int, session_id string) CreateRdvResponse {

	var appointment_status model.AppointmentStatus = "WAITING_FOR_REVIEW"
	rdv, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         patientId,
		DoctorID:          doctorId,
		StartDate:         startDate,
		EndDate:           endDate,
		AppointmentStatus: appointment_status,
		SessionID:         session_id,
	})
	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	doctor, err := graphql.GetDoctorById(doctorId)
	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedDoctor, err := graphql.UpdateDoctor(doctorId, model.UpdateDoctorInput{
		RendezVousIds: append(doctor.RendezVousIds, &rdv.ID),
	})

	if err != nil {
		return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 500, Err: errors.New("update failed" + err.Error())}
	}

	if patientId != "" {
		patient, err := graphql.GetPatientById(patientId)
		if err != nil {
			return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to an patient")}
		}
		_, err = graphql.UpdatePatient(patientId, model.UpdatePatientInput{
			RendezVousIds: append(patient.RendezVousIds, &rdv.ID),
		})
		if err != nil {
			return CreateRdvResponse{Rdv: model.Rdv{}, Doctor: model.Doctor{}, Code: 500, Err: errors.New("unable to update patient")}
		}
	}

	return CreateRdvResponse{
		Rdv:    rdv,
		Doctor: updatedDoctor,
		Code:   200,
		Err:    nil,
	}
}
