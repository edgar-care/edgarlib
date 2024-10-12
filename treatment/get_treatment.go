package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type GetTreatmentByIdResponse struct {
	Treatment   model.AntecedentTreatment
	Antedisease model.AnteDisease
	Code        int
	Err         error
}

type GetTreatmentsResponse struct {
	Antedisease []model.AnteDisease
	Treatments  []model.AntecedentTreatment
	Code        int
	Err         error
}

func GetTreatmentById(id string, patientID string) GetTreatmentByIdResponse {
	treatment, err := graphql.GetAntecedentTreatmentByID(id)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.AntecedentTreatment{}, Code: 400, Err: errors.New("id does not correspond to a Treatment")}
	}

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetTreatmentByIdResponse{Treatment: model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get medical_info from patient")}
	}

	return GetTreatmentByIdResponse{Treatment: treatment, Code: 200, Err: nil}
}

func GetTreatments(patientID string, treatmentID string) GetTreatmentsResponse {
	var res []model.AntecedentTreatment

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get medical_info from patient")}
	}

	medicalInfo, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get medical_info with id: " + err.Error())}
	}

	treatmentIDs := medicalInfo.AntecedentTreatmentIds
	if treatmentID != "" {
		treatmentIDs = []string{treatmentID}
	}

	for _, treatmentID := range treatmentIDs {
		treatment, err := graphql.GetAntecedentTreatments(treatmentID)
		if err != nil {
			return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("invalid input: " + err.Error())}
		}

		res = append(res, treatment)
	}

	return GetTreatmentsResponse{Treatments: res, Code: 200, Err: nil}
}
