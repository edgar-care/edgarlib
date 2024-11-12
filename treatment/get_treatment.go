package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"log"
)

type GetTreatmentByIdResponse struct {
	Treatment model.AntecedentTreatment
	Code      int
	Err       error
}

type GetTreatmentsResponse struct {
	Treatments []model.AntecedentTreatment
	Code       int
	Err        error
}

func GetTreatmentById(id string, patientID string) GetTreatmentByIdResponse {
	log.Printf("GetTreatmentById called with id: %s, patientID: %s", id, patientID)

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetTreatmentByIdResponse{Treatment: model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get medical_info from patient")}
	}
	treatment, err := graphql.GetAntecedentTreatmentByID(id)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.AntecedentTreatment{}, Code: 400, Err: errors.New("id does not correspond to a Treatment")}
	}

	return GetTreatmentByIdResponse{Treatment: treatment, Code: 200, Err: nil}
}

func GetTreatments(patientID string) GetTreatmentsResponse {
	var res []model.AntecedentTreatment

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get medical_info from patient")}
	}

	medicalFolder, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get medical_info with id: " + err.Error())}
	}

	for _, id := range medicalFolder.AntecedentDiseaseIds {
		medicalAntecedent, err := graphql.GetMedicalAntecedentsById(id)
		if err != nil {
			return GetTreatmentsResponse{Treatments: []model.AntecedentTreatment{}, Code: 400, Err: errors.New("unable to get medical antecedent with id: " + err.Error())}
		}

		res = make([]model.AntecedentTreatment, len(medicalAntecedent.Treatments))
		for i, treatment := range medicalAntecedent.Treatments {
			res[i] = *treatment
		}
	}
	return GetTreatmentsResponse{Treatments: res, Code: 200, Err: nil}
}
