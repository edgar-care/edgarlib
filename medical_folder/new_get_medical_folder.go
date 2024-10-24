package medical_folder

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type MedicalFolderResponse struct {
	MedicalInfo        model.MedicalInfo
	MedicalAntecedents []model.MedicalAntecedents
	Code               int
	Err                error
}

type GetMedicalAntecedentsResponse struct {
	MedicalAntecedents []model.MedicalAntecedents
	Code               int
	Err                error
}

type GetMedicalAntecedentResponse struct {
	MedicalAntecedent model.MedicalAntecedents
	Code              int
	Err               error
}

func GetMedicalFolder(patientID string) MedicalFolderResponse {
	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return MedicalFolderResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return MedicalFolderResponse{Code: 404, Err: errors.New("medical folder not found for the patient")}
	}

	medicalInfo, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return MedicalFolderResponse{Code: 400, Err: errors.New("unable to retrieve medical info: " + err.Error())}
	}

	medicalAntecedents := []model.MedicalAntecedents{}
	for _, id := range medicalInfo.AntecedentDiseaseIds {
		antecedent, err := graphql.GetMedicalAntecedentsById(id)
		if err != nil {
			return MedicalFolderResponse{Code: 400, Err: errors.New("unable to retrieve antecedent: " + err.Error())}
		}
		medicalAntecedents = append(medicalAntecedents, antecedent)
	}

	return MedicalFolderResponse{
		MedicalInfo:        medicalInfo,
		MedicalAntecedents: medicalAntecedents,
		Code:               200,
		Err:                nil,
	}
}

func getMedicalAntecedents(patientID string) GetMedicalAntecedentsResponse {
	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetMedicalAntecedentsResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetMedicalAntecedentsResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	medicalFolder, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return GetMedicalAntecedentsResponse{Code: 400, Err: errors.New("unable to retrieve medical folder: " + err.Error())}
	}

	var medicalAntecedents []model.MedicalAntecedents
	for _, id := range medicalFolder.AntecedentDiseaseIds {
		antecedent, err := graphql.GetMedicalAntecedentsById(id)
		if err != nil {
			return GetMedicalAntecedentsResponse{Code: 400, Err: errors.New("unable to retrieve medical antecedents: " + err.Error())}
		}
		medicalAntecedents = append(medicalAntecedents, antecedent)
	}

	return GetMedicalAntecedentsResponse{
		MedicalAntecedents: medicalAntecedents,
		Code:               200,
		Err:                nil,
	}
}

func getMedicalAntecedentById(antecdentID string, patientID string) GetMedicalAntecedentResponse {
	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetMedicalAntecedentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetMedicalAntecedentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	medicalFolder, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return GetMedicalAntecedentResponse{Code: 400, Err: errors.New("unable to retrieve medical folder: " + err.Error())}
	}

	antecedentFound := false
	for _, id := range medicalFolder.AntecedentDiseaseIds {
		if id == antecdentID {
			antecedentFound = true
			break
		}
	}

	if !antecedentFound {
		return GetMedicalAntecedentResponse{Code: 400, Err: errors.New("antecedent id not found in medical folder")}
	}

	medicalAntecedent, err := graphql.GetMedicalAntecedentsById(antecdentID)
	if err != nil {
		return GetMedicalAntecedentResponse{Code: 400, Err: errors.New("unable to retrieve medical antecedent: " + err.Error())}
	}

	return GetMedicalAntecedentResponse{
		MedicalAntecedent: medicalAntecedent,
		Code:              200,
		Err:               nil,
	}
}
