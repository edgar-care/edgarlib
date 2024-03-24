package treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetTreatmentByIdResponse struct {
	Treatment model.Treatment
	Code      int
	Err       error
}

type GetTreatmentsResponse struct {
	Treatments []model.Treatment
	Code       int
	Err        error
}

func GetTreatmentById(id string) GetTreatmentByIdResponse {
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

	res = model.Treatment{
		ID:         treatment.GetTreatmentByID.Id,
		Period:     treatmentPeriods,
		Day:        treatmentDays,
		Quantity:   treatment.GetTreatmentByID.Quantity,
		MedicineID: treatment.GetTreatmentByID.Medicine_id,
	}
	return GetTreatmentByIdResponse{Treatment: res, Code: 200, Err: nil}
}

func GetTreatments(patientID string) GetTreatmentsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Treatment

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
	}

	return GetTreatmentsResponse{Treatments: res, Code: 200, Err: nil}
}
