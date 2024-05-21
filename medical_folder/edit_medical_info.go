package medical_folder

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"github.com/edgar-care/edgarlib/treatment"
)

type UpdateMedicalInfoInput struct {
	Name               string                         `json:"name"`
	Firstname          string                         `json:"firstname"`
	Birthdate          int                            `json:"birthdate"`
	Sex                string                         `json:"sex"`
	Weight             int                            `json:"weight"`
	Height             int                            `json:"height"`
	PrimaryDoctorID    string                         `json:"primary_doctor_id,omitempty"`
	MedicalAntecedents []UpdateMedicalAntecedentInput `json:"medical_antecedents"`
}

type UpdateMedicalAntecedentInput struct {
	ID            string                `json:"antedisease_id"`
	Name          string                `json:"name"`
	Medicines     []UpdateMedicineInput `json:"treatments"`
	StillRelevant bool                  `json:"still_relevant"`
}

type UpdateMedicineInput struct {
	ID         string   `json:"treatment_id"`
	MedicineID string   `json:"medicine_id"`
	Period     []string `json:"period"`
	Day        []string `json:"day"`
	Quantity   int      `json:"quantity"`
}

type UpdateMedicalFolderResponse struct {
	MedicalInfo                model.MedicalInfo
	AnteDiseasesWithTreatments []AnteDiseaseWithTreatments
	Code                       int
	Err                        error
}

func UpdateMedicalFolder(input UpdateMedicalInfoInput, medicalInfoID string) UpdateMedicalFolderResponse {
	var antdediseaseids []string
	var res []model.Treatment
	var antediseasesWithTreatments []AnteDiseaseWithTreatments

	gqlClient := graphql.CreateClient()

	control, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, medicalInfoID)
	if err != nil {
		return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to find medical folder by ID: " + err.Error())}
	}

	existingAntecedents := make(map[string]model.AnteDisease)
	for _, antecedentID := range control.GetMedicalFolderById.Antecedent_disease_ids {
		antecedent, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, antecedentID)
		if err != nil {
			return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to find antedisease by ID: " + err.Error())}
		}
		ante := model.AnteDisease{
			ID:            antecedent.GetAnteDiseaseByID.GetId(),
			Name:          antecedent.GetAnteDiseaseByID.GetName(),
			Chronicity:    antecedent.GetAnteDiseaseByID.GetChronicity(),
			SurgeryIds:    antecedent.GetAnteDiseaseByID.GetSurgery_ids(),
			Symptoms:      antecedent.GetAnteDiseaseByID.GetSymptoms(),
			TreatmentIds:  antecedent.GetAnteDiseaseByID.GetTreatment_ids(),
			StillRelevant: antecedent.GetAnteDiseaseByID.GetStill_relevant(),
		}

		existingAntecedents[antecedentID] = ante
	}

	// Suppression des traitements et des antécédents non présents dans le body
	for id, antecedent := range existingAntecedents {
		found := false
		for _, inputAntecedent := range input.MedicalAntecedents {
			if inputAntecedent.ID == id {
				found = true
				break
			}
		}

		if !found {
			for _, treatmentID := range antecedent.TreatmentIds {
				deleteResp := treatment.DeleteTreatment(treatmentID, medicalInfoID)
				if !deleteResp.Deleted {
					return UpdateMedicalFolderResponse{Code: deleteResp.Code, Err: deleteResp.Err}
				}
			}

			// Supprimer l'antécédent
			_, err := graphql.DeleteAnteDisease(context.Background(), gqlClient, id)
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 500, Err: errors.New("unable to delete antedisease: " + err.Error())}
			}
			_, err = graphql.UpdateMedicalFolder(context.Background(), gqlClient, medicalInfoID, control.GetMedicalFolderById.Name, control.GetMedicalFolderById.Firstname, control.GetMedicalFolderById.Birthdate, string(control.GetMedicalFolderById.Sex), control.GetMedicalFolderById.Height, control.GetMedicalFolderById.Weight, control.GetMedicalFolderById.Primary_doctor_id, remElement(control.GetMedicalFolderById.Antecedent_disease_ids, id), "DONE")
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to update medical folder: " + err.Error())}
			}
		}
	}

	// Création et édition des antécédents et traitements
	for _, antecedent := range input.MedicalAntecedents {
		var treatmentIDsPerAnte []string
		var antediseaseTreatments []model.Treatment

		if antecedent.ID == "" {
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
					return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
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
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create antedisease: " + err.Error())}
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

		} else {
			getAnte, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, antecedent.ID)
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to get antedisease: " + err.Error())}
			}

			_, err = graphql.UpdateAnteDisease(context.Background(), gqlClient, antecedent.ID, antecedent.Name, 0, []string{""}, []string{""}, getAnte.GetAnteDiseaseByID.Treatment_ids, antecedent.StillRelevant)
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to update antedisease: " + err.Error())}
			}

			for _, medicine := range antecedent.Medicines {
				periods := make([]graphql.Period, len(medicine.Period))
				for i, p := range medicine.Period {
					periods[i] = graphql.Period(p)
				}

				days := make([]graphql.Day, len(medicine.Day))
				for i, d := range medicine.Day {
					days[i] = graphql.Day(d)
				}

				if medicine.ID == "" {
					treatmentRes, err := graphql.CreateTreatment(context.Background(), gqlClient, periods, days, medicine.Quantity, medicine.MedicineID)
					if err != nil {
						return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
					}
					treatmentPeriods, treatmentDays := convertTreatmentPeriodsAndDays(treatmentRes.CreateTreatment.Period, treatmentRes.CreateTreatment.Day)

					treatmentToAdd := model.Treatment{
						ID:         treatmentRes.CreateTreatment.Id,
						Period:     treatmentPeriods,
						Day:        treatmentDays,
						Quantity:   treatmentRes.CreateTreatment.Quantity,
						MedicineID: treatmentRes.CreateTreatment.Medicine_id,
					}

					antediseaseTreatments = append(antediseaseTreatments, treatmentToAdd)
					res = append(res, treatmentToAdd)

				} else {
					treatmentRes, err := graphql.UpdateTreatment(context.Background(), gqlClient, medicine.ID, periods, days, medicine.Quantity, medicine.MedicineID)
					if err != nil {
						return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to update treatment: " + err.Error())}
					}
					treatmentPeriods, treatmentDays := convertTreatmentPeriodsAndDays(treatmentRes.UpdateTreatment.Period, treatmentRes.UpdateTreatment.Day)

					treatmentToAdd := model.Treatment{
						ID:         treatmentRes.UpdateTreatment.Id,
						Period:     treatmentPeriods,
						Day:        treatmentDays,
						Quantity:   treatmentRes.UpdateTreatment.Quantity,
						MedicineID: treatmentRes.UpdateTreatment.Medicine_id,
					}

					antediseaseTreatments = append(antediseaseTreatments, treatmentToAdd)
					res = append(res, treatmentToAdd)

				}
			}

			antediseaseWithTreatments := AnteDiseaseWithTreatments{
				AnteDisease: model.AnteDisease{
					ID:            getAnte.GetAnteDiseaseByID.Id,
					Name:          antecedent.Name,
					Chronicity:    getAnte.GetAnteDiseaseByID.Chronicity,
					TreatmentIds:  getAnte.GetAnteDiseaseByID.Treatment_ids,
					StillRelevant: antecedent.StillRelevant,
				},
				Treatments: antediseaseTreatments,
			}
			antediseasesWithTreatments = append(antediseasesWithTreatments, antediseaseWithTreatments)
		}
	}

	medical, err := graphql.UpdateMedicalFolder(context.Background(), gqlClient, medicalInfoID, input.Name, input.Firstname, input.Birthdate, input.Sex, input.Height, input.Weight, input.PrimaryDoctorID, append(control.GetMedicalFolderById.Antecedent_disease_ids, antdediseaseids...), "DONE")
	if err != nil {
		return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create medical folder: " + err.Error())}
	}

	return UpdateMedicalFolderResponse{
		MedicalInfo: model.MedicalInfo{
			ID:                   control.GetMedicalFolderById.Id,
			Name:                 medical.UpdateMedicalFolder.Name,
			Firstname:            medical.UpdateMedicalFolder.Firstname,
			Birthdate:            medical.UpdateMedicalFolder.Birthdate,
			Sex:                  model.Sex(medical.UpdateMedicalFolder.Sex),
			Weight:               medical.UpdateMedicalFolder.Weight,
			Height:               medical.UpdateMedicalFolder.Height,
			PrimaryDoctorID:      medical.UpdateMedicalFolder.Primary_doctor_id,
			OnboardingStatus:     model.OnboardingStatus(control.GetMedicalFolderById.Onboarding_status),
			AntecedentDiseaseIds: antdediseaseids,
		},
		AnteDiseasesWithTreatments: antediseasesWithTreatments,
		Code:                       200,
		Err:                        nil,
	}
}

func UpdateMedicalFolderFromDoctor(input UpdateMedicalInfoInput, PatientID string) UpdateMedicalFolderResponse {
	gqlClient := graphql.CreateClient()
	if PatientID == "" {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 400, Err: errors.New("medical info ID is required")}
	}
	patient, err := graphql.GetPatientById(context.Background(), gqlClient, PatientID)
	if err != nil {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	control, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id)
	if err != nil {
		return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to find medical folder by ID: " + err.Error())}
	}

	update := UpdateMedicalFolder(input, control.GetMedicalFolderById.Id)

	return UpdateMedicalFolderResponse{update.MedicalInfo, update.AnteDiseasesWithTreatments, update.Code, update.Err}
}

func convertTreatmentPeriodsAndDays(periods []graphql.Period, days []graphql.Day) ([]model.Period, []model.Day) {
	treatmentPeriods := make([]model.Period, len(periods))
	for i, p := range periods {
		treatmentPeriods[i] = model.Period(p)
	}

	treatmentDays := make([]model.Day, len(days))
	for i, d := range days {
		treatmentDays[i] = model.Day(d)
	}

	return treatmentPeriods, treatmentDays
}
