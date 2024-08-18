package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type GetTreatmentByIdResponse struct {
	Treatment   model.Treatment
	Antedisease model.AnteDisease
	Code        int
	Err         error
}

type GetTreatmentsResponse struct {
	Antedisease []model.AnteDisease
	Treatments  []model.Treatment
	Code        int
	Err         error
}

func GetTreatmentById(id string, patientID string) GetTreatmentByIdResponse {
	treatment, err := graphql.GetTreatmentByID(id)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("id does not correspond to a Treatment")}
	}

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("unable to get medical_info from patient")}
	}
	medicalInfo, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("unable to get medical_info with id: " + err.Error())}
	}

	antediseasesIDs := medicalInfo.AntecedentDiseaseIds
	antecedentTreatment := model.AnteDisease{}

	for _, antediseaseID := range antediseasesIDs {
		antedisease, err := graphql.GetAnteDiseaseByID(antediseaseID)
		if err != nil {
			return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("unable to get antedisease with id: " + err.Error())}
		}
		for _, treatmentID := range antedisease.TreatmentIds {
			if treatmentID == id {
				antecedentTreatment = antedisease
				break
			}
		}
	}

	return GetTreatmentByIdResponse{Treatment: treatment, Antedisease: antecedentTreatment, Code: 200, Err: nil}
}

func GetTreatments(patientID string) GetTreatmentsResponse {
	var res []model.Treatment
	var anteDisease []model.AnteDisease

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("unable to get medical_info from patient")}
	}
	medicalInfo, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("unable to get medical_info with id: " + err.Error())}
	}

	antediseasesIDs := medicalInfo.AntecedentDiseaseIds

	for _, antediseaseID := range antediseasesIDs {
		antedisease, err := graphql.GetAnteDiseaseByID(antediseaseID)
		if err != nil {
			return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("unable to get antedisease with id: " + err.Error())}
		}

		treatmentIDs := antedisease.TreatmentIds

		for _, treatmentID := range treatmentIDs {
			treatment, err := graphql.GetTreatmentByID(treatmentID)
			if err != nil {
				return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("invalid input: " + err.Error())}
			}

			res = append(res, treatment)
		}
		anteDisease = append(anteDisease, antedisease)
	}
	return GetTreatmentsResponse{Treatments: res, Antedisease: anteDisease, Code: 200, Err: nil}
}
