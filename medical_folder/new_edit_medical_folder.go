package medical_folder

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type UpdateMedicalFolderPatientInput struct {
	MedicalAntecedentInput model.UpdateMedicalAntecedentsInput `json:"medical_antecedent"`
}

type UpdateMedicalFolderPatientResponse struct {
	MedicalFolder model.MedicalInfo
	Code          int
	Err           error
}

type UpdateMedicalAntecedentResponse struct {
	UpdatedAntecedent model.MedicalAntecedents
	Code              int
	Err               error
}

func UpdateMedicalFolderPatient(patientID string, input model.UpdateMedicalFolderInput) UpdateMedicalFolderPatientResponse {
	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return UpdateMedicalFolderPatientResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return UpdateMedicalFolderPatientResponse{Code: 404, Err: errors.New("medical folder not found for the patient")}
	}

	updatedMedicalFolder, err := graphql.UpdateMedicalFolder(*patient.MedicalInfoID, input)
	if err != nil {
		return UpdateMedicalFolderPatientResponse{Code: 500, Err: errors.New("unable to update medical folder: " + err.Error())}
	}

	return UpdateMedicalFolderPatientResponse{
		MedicalFolder: updatedMedicalFolder,
		Code:          200,
		Err:           nil,
	}
}

func UpdateMedicalAntecedent(patientID string, input UpdateMedicalFolderPatientInput, MedicalAntecedentID string) UpdateMedicalAntecedentResponse {
	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return UpdateMedicalAntecedentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return UpdateMedicalAntecedentResponse{Code: 404, Err: errors.New("medical folder not found for the patient")}
	}

	updatedAntecedent, err := graphql.UpdateMedicalAntecedents(MedicalAntecedentID, input.MedicalAntecedentInput)
	if err != nil {
		return UpdateMedicalAntecedentResponse{Code: 500, Err: errors.New("unable to update antecedent: " + err.Error())}
	}

	return UpdateMedicalAntecedentResponse{
		UpdatedAntecedent: updatedAntecedent,
		Code:              200,
		Err:               nil,
	}
}
