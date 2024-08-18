package medical_folder

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type AnteDiseaseWithTreatments struct {
	AnteDisease model.AnteDisease
	Treatments  []model.Treatment
}

type MedicalInfoResponse struct {
	MedicalInfo                model.MedicalInfo
	AnteDiseasesWithTreatments []AnteDiseaseWithTreatments
	Code                       int
	Err                        error
}

func GetMedicalInfo(patientID string) MedicalInfoResponse {
	control, err := graphql.GetPatientById(patientID)
	if err != nil {
		return MedicalInfoResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if control.MedicalInfoID == nil || *control.MedicalInfoID == "" {
		return MedicalInfoResponse{Code: 404, Err: errors.New("medical folder not found for the patient")}
	}

	medical, err := graphql.GetMedicalFolderByID(*control.MedicalInfoID)
	if err != nil {
		return MedicalInfoResponse{Code: 400, Err: errors.New("unable to fetch medical folder: " + err.Error())}
	}

	if len(medical.AntecedentDiseaseIds) == 1 && medical.AntecedentDiseaseIds[0] == "" {
		return MedicalInfoResponse{
			MedicalInfo: medical,
			Code:        200,
			Err:         nil,
		}
	}

	var treatments []model.Treatment
	var antediseasesWithTreatments []AnteDiseaseWithTreatments
	for _, antediseaseID := range medical.AntecedentDiseaseIds {
		antedisease, err := graphql.GetAnteDiseaseByID(antediseaseID)
		if err != nil {
			return MedicalInfoResponse{Code: 400, Err: errors.New("unable to fetch antedisease: " + err.Error())}
		}

		var antediseaseTreatments []model.Treatment
		for _, treatmentID := range antedisease.TreatmentIds {
			treatment, err := graphql.GetTreatmentByID(treatmentID)
			if err != nil {
				return MedicalInfoResponse{Code: 400, Err: errors.New("unable to fetch treatment: " + err.Error())}
			}

			antediseaseTreatments = append(antediseaseTreatments, treatment)
			treatments = append(treatments, treatment)
		}

		antediseaseWithTreatments := AnteDiseaseWithTreatments{
			AnteDisease: antedisease,
			Treatments:  antediseaseTreatments,
		}

		antediseasesWithTreatments = append(antediseasesWithTreatments, antediseaseWithTreatments)
	}

	return MedicalInfoResponse{
		MedicalInfo:                medical,
		AnteDiseasesWithTreatments: antediseasesWithTreatments,
		Code:                       200,
		Err:                        nil,
	}
}
