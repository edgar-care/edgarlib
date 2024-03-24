package medical_folder

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	control, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return MedicalInfoResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if control.GetPatientById.Medical_info_id == "" {
		return MedicalInfoResponse{Code: 404, Err: errors.New("medical folder not found for the patient")}
	}

	medical, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, control.GetPatientById.Medical_info_id)
	if err != nil {
		return MedicalInfoResponse{Code: 400, Err: errors.New("unable to fetch medical folder: " + err.Error())}
	}

	var treatments []model.Treatment
	var antediseasesWithTreatments []AnteDiseaseWithTreatments
	for _, antediseaseID := range medical.GetMedicalFolderById.Antecedent_disease_ids {
		antedisease, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, antediseaseID)
		if err != nil {
			return MedicalInfoResponse{Code: 400, Err: errors.New("unable to fetch antedisease: " + err.Error())}
		}

		var antediseaseTreatments []model.Treatment
		for _, treatmentID := range antedisease.GetAnteDiseaseByID.Treatment_ids {
			treatment, err := graphql.GetTreatmentByID(context.Background(), gqlClient, treatmentID)
			if err != nil {
				return MedicalInfoResponse{Code: 400, Err: errors.New("unable to fetch treatment: " + err.Error())}
			}

			var periods []model.Period
			for _, period := range treatment.GetTreatmentByID.Period {
				periods = append(periods, model.Period(period))
			}

			var days []model.Day
			for _, day := range treatment.GetTreatmentByID.Day {
				days = append(days, model.Day(day))
			}

			treatmentToAdd := model.Treatment{
				ID:         treatment.GetTreatmentByID.Id,
				Period:     periods,
				Day:        days,
				Quantity:   treatment.GetTreatmentByID.Quantity,
				MedicineID: treatment.GetTreatmentByID.Medicine_id,
			}

			antediseaseTreatments = append(antediseaseTreatments, treatmentToAdd)
			treatments = append(treatments, treatmentToAdd)
		}

		antediseaseWithTreatments := AnteDiseaseWithTreatments{
			AnteDisease: model.AnteDisease{
				ID:            antedisease.GetAnteDiseaseByID.Id,
				Name:          antedisease.GetAnteDiseaseByID.Name,
				Chronicity:    antedisease.GetAnteDiseaseByID.Chronicity,
				SurgeryIds:    antedisease.GetAnteDiseaseByID.Surgery_ids,
				Symptoms:      antedisease.GetAnteDiseaseByID.Symptoms,
				TreatmentIds:  antedisease.GetAnteDiseaseByID.Treatment_ids,
				StillRelevant: antedisease.GetAnteDiseaseByID.Still_relevant,
			},
			Treatments: antediseaseTreatments,
		}

		antediseasesWithTreatments = append(antediseasesWithTreatments, antediseaseWithTreatments)
	}

	return MedicalInfoResponse{
		MedicalInfo: model.MedicalInfo{
			ID:                   medical.GetMedicalFolderById.Id,
			Name:                 medical.GetMedicalFolderById.Name,
			Firstname:            medical.GetMedicalFolderById.Firstname,
			Birthdate:            medical.GetMedicalFolderById.Birthdate,
			Sex:                  model.Sex(medical.GetMedicalFolderById.Sex),
			Weight:               medical.GetMedicalFolderById.Weight,
			Height:               medical.GetMedicalFolderById.Height,
			PrimaryDoctorID:      medical.GetMedicalFolderById.Primary_doctor_id,
			OnboardingStatus:     model.OnboardingStatus(medical.GetMedicalFolderById.Onboarding_status),
			AntecedentDiseaseIds: medical.GetMedicalFolderById.Antecedent_disease_ids,
		},
		AnteDiseasesWithTreatments: antediseasesWithTreatments,
		Code:                       200,
		Err:                        nil,
	}
}
