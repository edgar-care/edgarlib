package medical_folder

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type DeleteMedicalInfoResponse struct {
	Deleted        bool
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func DeleteMedicalInfo(medicalId string, patientId string) DeleteMedicalInfoResponse {
	if medicalId == "" {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("medical id is required")}
	}

	_, err := graphql.GetMedicalFolderByID(medicalId)
	if err != nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a medical folder")}
	}

	deleted, err := graphql.DeleteMedicalFolder(medicalId)
	if err != nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	emptyMedical := ""
	patient, err := graphql.UpdatePatient(patientId, model.UpdatePatientInput{MedicalInfoID: &emptyMedical})

	if err != nil {
		return DeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	return DeleteMedicalInfoResponse{
		Deleted:        deleted,
		UpdatedPatient: patient,
		Code:           200,
		Err:            nil,
	}
}
