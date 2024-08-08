package medical_folder

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type UpdateMedicalInfoInput struct {
	Name                   string                         `json:"name"`
	Firstname              string                         `json:"firstname"`
	Birthdate              int                            `json:"birthdate"`
	Sex                    string                         `json:"sex"`
	Weight                 int                            `json:"weight"`
	Height                 int                            `json:"height"`
	PrimaryDoctorID        string                         `json:"primary_doctor_id,omitempty"`
	MedicalAntecedents     []UpdateMedicalAntecedentInput `json:"medical_antecedents"`
	FamilyMembersMedInfoId []string                       `json:"family_members_med_info_id"`
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
	var antediseasesWithTreatments []AnteDiseaseWithTreatments
	control, err := graphql.GetMedicalFolderByID(medicalInfoID)
	if err != nil {
		return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to find medical folder by ID: " + err.Error())}
	}

	var updatedAntecedentDiseaseIDs []string

	for _, antecedent := range input.MedicalAntecedents {
		var treatmentIDsPerAnte []string
		var antediseaseTreatments []model.Treatment
		if antecedent.ID == "" {
			for _, medicine := range antecedent.Medicines {
				periods, days := ConvertPeriodsAndDays(medicine.Period, medicine.Day)
				treatmentRes, err := graphql.CreateTreatment(model.CreateTreatmentInput{
					Period:     periods,
					Day:        days,
					Quantity:   medicine.Quantity,
					MedicineID: medicine.MedicineID,
				})
				if err != nil {
					return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
				}
				treatmentIDsPerAnte = append(treatmentIDsPerAnte, treatmentRes.ID)
				antediseaseTreatments = append(antediseaseTreatments, treatmentRes)
			}

			antedisease, err := graphql.CreateAnteDisease(model.CreateAnteDiseaseInput{
				Name:          antecedent.Name,
				TreatmentIds:  treatmentIDsPerAnte,
				StillRelevant: antecedent.StillRelevant,
			})
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create antedisease: " + err.Error())}
			}
			antdediseaseids = append(antdediseaseids, antedisease.ID)
			antediseaseWithTreatments := AnteDiseaseWithTreatments{
				AnteDisease: antedisease,
				Treatments:  antediseaseTreatments,
			}
			antediseasesWithTreatments = append(antediseasesWithTreatments, antediseaseWithTreatments)
		} else {

			medicalAntecedentIDs := control.AntecedentDiseaseIds

			for _, antecedentID := range medicalAntecedentIDs {
				found := false
				for _, ante := range input.MedicalAntecedents {
					if ante.ID == antecedentID {
						found = true
						break
					}
				}
				if found || antecedentID == "" {
					updatedAntecedentDiseaseIDs = append(updatedAntecedentDiseaseIDs, antecedentID)
				} else {
					_, err = graphql.DeleteAnteDisease(antecedentID)
					if err != nil {
						return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to delete antedisease: " + err.Error())}
					}
				}
			}

			status := model.OnboardingStatusDone

			_, err = graphql.UpdateMedicalFolder(medicalInfoID, model.UpdateMedicalFolderInput{AntecedentDiseaseIds: updatedAntecedentDiseaseIDs, OnboardingStatus: &status})
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to update medical folder: " + err.Error())}
			}

			getAnte, err := graphql.GetAnteDiseaseByID(antecedent.ID)
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to get antedisease: " + err.Error())}
			}

			updatedTreatmentIDs := make(map[string]bool)
			for _, tid := range getAnte.TreatmentIds {
				updatedTreatmentIDs[tid] = true
			}

			existingTreatmentIDs := make(map[string]bool)
			for _, tid := range getAnte.TreatmentIds {
				existingTreatmentIDs[tid] = true
			}

			for _, medicine := range antecedent.Medicines {
				periods, days := ConvertPeriodsAndDays(medicine.Period, medicine.Day)
				if medicine.ID == "" {
					treatmentRes, err := graphql.CreateTreatment(model.CreateTreatmentInput{
						Period:     periods,
						Day:        days,
						Quantity:   medicine.Quantity,
						MedicineID: medicine.MedicineID,
					})
					if err != nil {
						return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
					}

					updatedTreatmentIDs[treatmentRes.ID] = true

					antediseaseTreatments = append(antediseaseTreatments, treatmentRes)
				} else {
					treatmentRes, err := graphql.UpdateTreatment(medicine.ID, model.UpdateTreatmentInput{
						Period:     periods,
						Day:        days,
						Quantity:   &medicine.Quantity,
						MedicineID: &medicine.MedicineID,
					})
					if err != nil {
						return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to update treatment: " + err.Error())}
					}

					antediseaseTreatments = append(antediseaseTreatments, treatmentRes)

					updatedTreatmentIDs[medicine.ID] = true
					delete(existingTreatmentIDs, medicine.ID)
				}
			}

			for treatmentID := range existingTreatmentIDs {
				_, err := graphql.DeleteTreatment(treatmentID)
				if err != nil {
					return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to delete treatment: " + err.Error())}
				}
				delete(updatedTreatmentIDs, treatmentID)
			}

			finalTreatmentIDs := make([]string, 0, len(updatedTreatmentIDs))
			for tid := range updatedTreatmentIDs {
				finalTreatmentIDs = append(finalTreatmentIDs, tid)
			}

			updatedAnteDisease, err := graphql.UpdateAnteDisease(antecedent.ID, model.UpdateAnteDiseaseInput{TreatmentIds: finalTreatmentIDs})
			if err != nil {
				return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to update antedisease: " + err.Error())}
			}

			antediseasesWithTreatments = append(antediseasesWithTreatments, AnteDiseaseWithTreatments{
				AnteDisease: updatedAnteDisease,
				Treatments:  antediseaseTreatments,
			})
		}

	}

	updatedAntecedentDiseaseIDs = append(updatedAntecedentDiseaseIDs, antdediseaseids...)

	status := model.OnboardingStatusDone

	medical, err := graphql.UpdateMedicalFolder(medicalInfoID, model.UpdateMedicalFolderInput{
		Name:                   &input.Name,
		Firstname:              &input.Firstname,
		Birthdate:              &input.Birthdate,
		Sex:                    &input.Sex,
		Height:                 &input.Height,
		Weight:                 &input.Weight,
		PrimaryDoctorID:        &input.PrimaryDoctorID,
		AntecedentDiseaseIds:   updatedAntecedentDiseaseIDs,
		OnboardingStatus:       &status,
		FamilyMembersMedInfoID: input.FamilyMembersMedInfoId,
	})
	if err != nil {
		return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to create medical folder: " + err.Error())}
	}

	return UpdateMedicalFolderResponse{
		MedicalInfo:                medical,
		AnteDiseasesWithTreatments: antediseasesWithTreatments,
		Code:                       200,
		Err:                        nil,
	}
}

func UpdateMedicalFolderFromDoctor(input UpdateMedicalInfoInput, PatientID string) UpdateMedicalFolderResponse {
	if PatientID == "" {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 400, Err: errors.New("medical info ID is required")}
	}
	patient, err := graphql.GetPatientById(PatientID)
	if err != nil {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}
	if patient.MedicalInfoID == nil {
		return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("patient doesn't have a medical folder")}
	}
	control, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return UpdateMedicalFolderResponse{Code: 400, Err: errors.New("unable to find medical folder by ID: " + err.Error())}
	}

	update := UpdateMedicalFolder(input, control.ID)

	return UpdateMedicalFolderResponse{update.MedicalInfo, update.AnteDiseasesWithTreatments, update.Code, update.Err}
}
