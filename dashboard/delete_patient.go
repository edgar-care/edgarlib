package dashboard

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type DeletePatientResponse struct {
	UpdatedDoctor model.Doctor
	Code          int
	Err           error
}

func remElement(slice []*string, element string) []*string {
	var result []*string
	for _, v := range slice {
		if *v != element {
			result = append(result, v)
		}
	}
	return result
}

func DeletePatient(PatientId string, doctorId string) DeletePatientResponse {
	if PatientId == "" {
		return DeletePatientResponse{UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("patient id is required")}
	}

	_, err := graphql.GetPatientById(PatientId)
	if err != nil {
		return DeletePatientResponse{UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	doctor, err := graphql.GetDoctorById(doctorId)
	if err != nil {
		return DeletePatientResponse{UpdatedDoctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	updatedDoctor, err := graphql.UpdateDoctorsPatientIDs(doctorId, model.UpdateDoctorsPatientIDsInput{
		PatientIds: remElement(doctor.PatientIds, PatientId),
	})

	if err != nil {
		return DeletePatientResponse{UpdatedDoctor: model.Doctor{}, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	return DeletePatientResponse{
		UpdatedDoctor: updatedDoctor,
		Code:          200,
		Err:           nil}
}
