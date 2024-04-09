package medical_folder

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateMedicalInfoInput struct {
	Name               string                         `json:"name"`
	Firstname          string                         `json:"firstname"`
	Birthdate          int                            `json:"birthdate"`
	Sex                string                         `json:"sex"`
	Weight             int                            `json:"weight"`
	Height             int                            `json:"height"`
	PrimaryDoctorID    string                         `json:"primary_doctor_id,omitempty"`
	MedicalAntecedents []CreateMedicalAntecedentInput `json:"medical_antecedents"`
}

type CreateMedicalAntecedentInput struct {
	Name          string                `json:"name"`
	Medicines     []CreateMedicineInput `json:"treatments"`
	StillRelevant bool                  `json:"still_relevant"`
}

type CreateMedicineInput struct {
	MedicineID string   `json:"medicine_id"`
	Period     []string `json:"period"`
	Day        []string `json:"day"`
	Quantity   int      `json:"quantity"`
}

type CreateMedicalInfoResponse struct {
	MedicalInfo                model.MedicalInfo
	AnteDiseasesWithTreatments []AnteDiseaseWithTreatments
	Code                       int
	Err                        error
}

func CreateMedicalInfo(input CreateMedicalInfoInput, patientID string) CreateMedicalInfoResponse {
	var antdediseaseids []string
	var res []model.Treatment
	var antediseasesWithTreatments []AnteDiseaseWithTreatments

	gqlClient := graphql.CreateClient()

	control, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if control.GetPatientById.Medical_info_id != "" {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("medical folder has already been created")}
	}

	if len(input.MedicalAntecedents) == 0 {
		medical, err := graphql.CreateMedicalFolder(context.Background(), gqlClient, input.Name, input.Firstname, input.Birthdate, input.Sex, input.Height, input.Weight, input.PrimaryDoctorID, []string{""}, "DONE")
		if err != nil {
			return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create medical folder: " + err.Error())}
		}

		_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, control.GetPatientById.Email, control.GetPatientById.Password, medical.CreateMedicalFolder.Id, control.GetPatientById.Rendez_vous_ids, control.GetPatientById.Document_ids, control.GetPatientById.Treatment_follow_up_ids, control.GetPatientById.Chat_ids)
		if err != nil {
			return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to update patient: " + err.Error())}
		}

		return CreateMedicalInfoResponse{
			MedicalInfo: model.MedicalInfo{
				ID:                   medical.CreateMedicalFolder.Id,
				Name:                 medical.CreateMedicalFolder.Name,
				Firstname:            medical.CreateMedicalFolder.Firstname,
				Birthdate:            medical.CreateMedicalFolder.Birthdate,
				Sex:                  model.Sex(medical.CreateMedicalFolder.Sex),
				Weight:               medical.CreateMedicalFolder.Weight,
				Height:               medical.CreateMedicalFolder.Height,
				PrimaryDoctorID:      medical.CreateMedicalFolder.Primary_doctor_id,
				OnboardingStatus:     model.OnboardingStatus(medical.CreateMedicalFolder.Onboarding_status),
				AntecedentDiseaseIds: medical.CreateMedicalFolder.Antecedent_disease_ids,
			},
			Code: 201,
			Err:  nil,
		}
	}

	for _, antecedent := range input.MedicalAntecedents {
		var treatmentIDsPerAnte []string
		var antediseaseTreatments []model.Treatment

		for _, medicine := range antecedent.Medicines {
			periods := make([]graphql.Period, len(medicine.Period))
			for i, p := range medicine.Period {
				periods[i] = graphql.Period(p)
			}

			days := make([]graphql.Day, len(medicine.Day))
			for i, d := range medicine.Day {
				days[i] = graphql.Day(d)
			}

			treatmentRes, err := graphql.CreateTreatment(context.Background(), gqlClient, periods, days, medicine.Quantity, medicine.MedicineID)
			if err != nil {
				return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
			}

			treatmentIDsPerAnte = append(treatmentIDsPerAnte, treatmentRes.CreateTreatment.Id)

			treatmentPeriods := make([]model.Period, len(treatmentRes.CreateTreatment.Period))
			for i, p := range treatmentRes.CreateTreatment.Period {
				treatmentPeriods[i] = model.Period(p)
			}

			treatmentDays := make([]model.Day, len(treatmentRes.CreateTreatment.Day))
			for i, d := range treatmentRes.CreateTreatment.Day {
				treatmentDays[i] = model.Day(d)
			}

			treatmentToAdd := model.Treatment{
				ID:         treatmentRes.CreateTreatment.Id,
				Period:     treatmentPeriods,
				Day:        treatmentDays,
				Quantity:   treatmentRes.CreateTreatment.Quantity,
				MedicineID: treatmentRes.CreateTreatment.Medicine_id,
			}

			antediseaseTreatments = append(antediseaseTreatments, treatmentToAdd)
			res = append(res, treatmentToAdd)
		}

		antedisease, err := graphql.CreateAnteDisease(context.Background(), gqlClient, antecedent.Name, 0, []string{""}, nil, treatmentIDsPerAnte, antecedent.StillRelevant)
		if err != nil {
			return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create antedisease: " + err.Error())}
		}

		antdediseaseids = append(antdediseaseids, antedisease.CreateAnteDisease.Id)

		antediseaseWithTreatments := AnteDiseaseWithTreatments{
			AnteDisease: model.AnteDisease{
				ID:            antedisease.CreateAnteDisease.Id,
				Name:          antedisease.CreateAnteDisease.Name,
				Chronicity:    antedisease.CreateAnteDisease.Chronicity,
				SurgeryIds:    antedisease.CreateAnteDisease.Surgery_ids,
				Symptoms:      antedisease.CreateAnteDisease.Symptoms,
				TreatmentIds:  antedisease.CreateAnteDisease.Treatment_ids,
				StillRelevant: antedisease.CreateAnteDisease.Still_relevant,
			},
			Treatments: antediseaseTreatments,
		}

		antediseasesWithTreatments = append(antediseasesWithTreatments, antediseaseWithTreatments)
	}

	medical, err := graphql.CreateMedicalFolder(context.Background(), gqlClient, input.Name, input.Firstname, input.Birthdate, input.Sex, input.Height, input.Weight, input.PrimaryDoctorID, antdediseaseids, "DONE")
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create medical folder: " + err.Error())}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, control.GetPatientById.Email, control.GetPatientById.Password, medical.CreateMedicalFolder.Id, control.GetPatientById.Rendez_vous_ids, control.GetPatientById.Document_ids, control.GetPatientById.Treatment_follow_up_ids, control.GetPatientById.Chat_ids)
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to update patient: " + err.Error())}
	}

	return CreateMedicalInfoResponse{
		MedicalInfo: model.MedicalInfo{
			ID:                   medical.CreateMedicalFolder.Id,
			Name:                 medical.CreateMedicalFolder.Name,
			Firstname:            medical.CreateMedicalFolder.Firstname,
			Birthdate:            medical.CreateMedicalFolder.Birthdate,
			Sex:                  model.Sex(medical.CreateMedicalFolder.Sex),
			Weight:               medical.CreateMedicalFolder.Weight,
			Height:               medical.CreateMedicalFolder.Height,
			PrimaryDoctorID:      medical.CreateMedicalFolder.Primary_doctor_id,
			OnboardingStatus:     model.OnboardingStatus(medical.CreateMedicalFolder.Onboarding_status),
			AntecedentDiseaseIds: antdediseaseids,
		},
		AnteDiseasesWithTreatments: antediseasesWithTreatments,
		Code:                       201,
		Err:                        nil,
	}
}
