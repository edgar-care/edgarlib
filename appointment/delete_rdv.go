package appointment

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type DeleteRdvResponse struct {
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func DeleteRdv(rdvId string, patientId string) DeleteRdvResponse {
	if rdvId == "" {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("rdv id is required")}
	}

	rdv, err := graphql.GetRdvById(rdvId)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}
	var appointment_status = model.AppointmentStatusCanceled
	_, err = graphql.UpdateRdv(rdv.ID, model.UpdateRdvInput{
		AppointmentStatus: &appointment_status,
	})
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("could not update appointment")}
	}

	var new_appointment_status model.AppointmentStatus = "OPENED"
	new_slot, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          rdv.DoctorID,
		StartDate:         rdv.StartDate,
		EndDate:           rdv.EndDate,
		AppointmentStatus: new_appointment_status,
		SessionID:         "",
	})
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	doctor, err := graphql.GetDoctorById(rdv.DoctorID)
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdateDoctor(rdv.DoctorID, model.UpdateDoctorInput{
		RendezVousIds: append(doctor.RendezVousIds, &new_slot.ID),
	})
	if err != nil {
		return DeleteRdvResponse{UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}
	return DeleteRdvResponse{
		UpdatedPatient: patient,
		Code:           200,
		Err:            nil,
	}
}

func removeElement(slice []*string, element *string) []*string {
	var result []*string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
