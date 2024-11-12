package medical_folder

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type NewDeleteMedicalInfoResponse struct {
	Deleted        bool
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

type DeleteMedicalAntecedentResponse struct {
	Deleted        bool
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func NewDeleteMedicalInfo(medicalId string, patientId string) NewDeleteMedicalInfoResponse {
	if medicalId == "" {
		return NewDeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("medical id is required")}
	}

	_, err := graphql.GetMedicalFolderByID(medicalId)
	if err != nil {
		return NewDeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a medical folder")}
	}

	deleted, err := graphql.DeleteMedicalFolder(medicalId)
	if err != nil {
		return NewDeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	emptyMedical := ""
	patient, err := graphql.UpdatePatient(patientId, model.UpdatePatientInput{MedicalInfoID: &emptyMedical})

	if err != nil {
		return NewDeleteMedicalInfoResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	return NewDeleteMedicalInfoResponse{
		Deleted:        deleted,
		UpdatedPatient: patient,
		Code:           200,
		Err:            nil,
	}
}

func DeleteMedicalAntecedent(antecedentId string, patientId string) DeleteMedicalAntecedentResponse {
	if antecedentId == "" {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("antecedent id is required")}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("medical folder has not been created")}
	}

	medicalFolder, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("unable to retrieve medical folder: " + err.Error())}
	}

	antecedentFound := false
	for _, id := range medicalFolder.AntecedentDiseaseIds {
		if id == antecedentId {
			antecedentFound = true
			break
		}
	}

	if !antecedentFound {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("antecedent id not found in medical folder")}
	}

	_, err = graphql.GetMedicalAntecedentsById(antecedentId)
	if err != nil {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a medical antecedent")}
	}

	deleted, err := graphql.DeleteMedicalAntecedents(antecedentId)
	if err != nil {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error while deleting antecedent: " + err.Error())}
	}

	var updatedAntecedentDiseaseIds []*string
	for _, id := range medicalFolder.AntecedentDiseaseIds {
		if id != antecedentId {
			idCopy := id
			updatedAntecedentDiseaseIds = append(updatedAntecedentDiseaseIds, &idCopy)
		}
	}

	_, err = graphql.UpdateMedicalFolderdAntedisease(*patient.MedicalInfoID, model.UpdateMedicalFOlderAntedisease{
		AntecedentDiseaseIds: updatedAntecedentDiseaseIds,
	})
	if err != nil {
		return DeleteMedicalAntecedentResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("unable to update medical folder: " + err.Error())}
	}

	return DeleteMedicalAntecedentResponse{
		Deleted:        deleted,
		UpdatedPatient: patient,
		Code:           200,
	}
}
