package treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()
	var res model.Treatment

	treatment, err := graphql.GetTreatmentByID(context.Background(), gqlClient, id)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("id does not correspond to a Treatment")}
	}

	treatmentPeriods := make([]model.Period, len(treatment.GetTreatmentByID.Period))
	for i, p := range treatment.GetTreatmentByID.Period {
		treatmentPeriods[i] = model.Period(p)
	}

	treatmentDays := make([]model.Day, len(treatment.GetTreatmentByID.Day))
	for i, d := range treatment.GetTreatmentByID.Day {
		treatmentDays[i] = model.Day(d)
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	medicalInfo, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id)
	if err != nil {
		return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("unable to get medical_info with id: " + err.Error())}
	}

	antediseasesIDs := medicalInfo.GetMedicalFolderById.Antecedent_disease_ids
	antecedentTreatment := model.AnteDisease{}

	for _, antediseaseID := range antediseasesIDs {
		antedisease, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, antediseaseID)
		if err != nil {
			return GetTreatmentByIdResponse{Treatment: model.Treatment{}, Code: 400, Err: errors.New("unable to get antedisease with id: " + err.Error())}
		}
		for _, treatmentID := range antedisease.GetAnteDiseaseByID.Treatment_ids {
			if treatmentID == id {
				antecedentTreatment = model.AnteDisease{
					ID:            antedisease.GetAnteDiseaseByID.Id,
					Name:          antedisease.GetAnteDiseaseByID.Name,
					StillRelevant: antedisease.GetAnteDiseaseByID.Still_relevant,
				}
				break
			}
		}
	}

	res = model.Treatment{
		ID:         treatment.GetTreatmentByID.Id,
		Period:     treatmentPeriods,
		Day:        treatmentDays,
		Quantity:   treatment.GetTreatmentByID.Quantity,
		MedicineID: treatment.GetTreatmentByID.Medicine_id,
	}
	return GetTreatmentByIdResponse{Treatment: res, Antedisease: antecedentTreatment, Code: 200, Err: nil}
}

func GetTreatments(patientID string) GetTreatmentsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Treatment
	var anteDisease []model.AnteDisease

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("unable to get patient with id: " + err.Error())}
	}

	medicalInfo, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id)
	if err != nil {
		return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("unable to get medical_info with id: " + err.Error())}
	}

	antediseasesIDs := medicalInfo.GetMedicalFolderById.Antecedent_disease_ids

	for _, antediseaseID := range antediseasesIDs {
		antedisease, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, antediseaseID)
		if err != nil {
			return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("unable to get antedisease with id: " + err.Error())}
		}

		treatmentIDs := antedisease.GetAnteDiseaseByID.Treatment_ids

		for _, treatmentID := range treatmentIDs {
			treatment, err := graphql.GetTreatmentByID(context.Background(), gqlClient, treatmentID)
			if err != nil {
				return GetTreatmentsResponse{Treatments: []model.Treatment{}, Code: 400, Err: errors.New("invalid input: " + err.Error())}
			}

			treatmentPeriods := make([]model.Period, len(treatment.GetTreatmentByID.Period))
			for i, p := range treatment.GetTreatmentByID.Period {
				treatmentPeriods[i] = model.Period(p)
			}

			treatmentDays := make([]model.Day, len(treatment.GetTreatmentByID.Day))
			for i, d := range treatment.GetTreatmentByID.Day {
				treatmentDays[i] = model.Day(d)
			}

			res = append(res, model.Treatment{
				ID:         treatment.GetTreatmentByID.Id,
				Period:     treatmentPeriods,
				Day:        treatmentDays,
				Quantity:   treatment.GetTreatmentByID.Quantity,
				MedicineID: treatment.GetTreatmentByID.Medicine_id,
			})
		}
		anteDisease = append(anteDisease, model.AnteDisease{
			ID:            antedisease.GetAnteDiseaseByID.Id,
			Name:          antedisease.GetAnteDiseaseByID.Name,
			StillRelevant: antedisease.GetAnteDiseaseByID.Still_relevant,
		})
	}
	return GetTreatmentsResponse{Treatments: res, Antedisease: anteDisease, Code: 200, Err: nil}
}
