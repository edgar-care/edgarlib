package dashboard

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
)

type CreatePatientInput struct {
	Email       string                                `json:"email"`
	MedicalInfo medical_folder.CreateMedicalInfoInput `json:"medical_info"`
}

type PatientByIdResponse struct {
	Patient                    model.Patient
	MedicalInfo                model.MedicalInfo
	AnteDiseasesWithTreatments []medical_folder.AnteDiseaseWithTreatments
	Code                       int
	Err                        error
}

func ConvertCreateToUpdateMedicalInfoInput(createInput medical_folder.CreateMedicalInfoInput) medical_folder.UpdateMedicalInfoInput {
	var updateAntecedents []medical_folder.UpdateMedicalAntecedentInput
	for _, antecedent := range createInput.MedicalAntecedents {
		var updateMedicines []medical_folder.UpdateMedicineInput
		for _, medicine := range antecedent.Medicines {
			updateMedicines = append(updateMedicines, medical_folder.UpdateMedicineInput{
				MedicineID: medicine.MedicineID,
				Period:     medicine.Period,
				Day:        medicine.Day,
				Quantity:   medicine.Quantity,
				StartDate:  medicine.StartDate,
				EndDate:    medicine.EndDate,
			})
		}
		updateAntecedents = append(updateAntecedents, medical_folder.UpdateMedicalAntecedentInput{
			Name:          antecedent.Name,
			Medicines:     updateMedicines,
			StillRelevant: antecedent.StillRelevant,
		})
	}

	return medical_folder.UpdateMedicalInfoInput{
		Name:               createInput.Name,
		Firstname:          createInput.Firstname,
		Birthdate:          createInput.Birthdate,
		Sex:                createInput.Sex,
		Weight:             createInput.Weight,
		Height:             createInput.Height,
		PrimaryDoctorID:    createInput.PrimaryDoctorID,
		MedicalAntecedents: updateAntecedents,
	}
}

func CreatePatientFormDoctor(newPatient CreatePatientInput, doctorID string) PatientByIdResponse {
	existingPatient, err := graphql.GetPatientByEmail(newPatient.Email)
	if err == nil {
		updateMedicalInfoInput := ConvertCreateToUpdateMedicalInfoInput(newPatient.MedicalInfo)
		if existingPatient.MedicalInfoID == nil || *existingPatient.MedicalInfoID == "" {
			return PatientByIdResponse{Code: 404, Err: errors.New("medical folder not found")}
		}
		medicalInfo := medical_folder.UpdateMedicalFolder(updateMedicalInfoInput, *existingPatient.MedicalInfoID)
		if medicalInfo.Err != nil {
			return PatientByIdResponse{Code: 400, Err: errors.New("unable to update medical information" + medicalInfo.Err.Error())}
		}

		doctor, err := graphql.GetDoctorById(doctorID)
		if err != nil {
			return PatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
		}

		_, err = graphql.UpdateDoctor(doctorID, model.UpdateDoctorInput{
			PatientIds: append(doctor.PatientIds, &existingPatient.ID),
		})
		if err != nil {
			return PatientByIdResponse{Code: 400, Err: errors.New("update failed: " + err.Error())}
		}

		return PatientByIdResponse{
			Patient: model.Patient{
				ID:                   existingPatient.ID,
				Email:                existingPatient.Email,
				MedicalInfoID:        &medicalInfo.MedicalInfo.ID,
				RendezVousIds:        existingPatient.RendezVousIds,
				DocumentIds:          existingPatient.DocumentIds,
				TreatmentFollowUpIds: existingPatient.TreatmentFollowUpIds,
			},
			MedicalInfo:                medicalInfo.MedicalInfo,
			AnteDiseasesWithTreatments: medicalInfo.AnteDiseasesWithTreatments,
			Code:                       200,
			Err:                        nil,
		}
	}

	patient := auth.CreatePatientAccount(newPatient.Email)
	if patient.Err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("unable to create a new account" + patient.Err.Error())}
	}

	getPatient, err := graphql.GetPatientById(patient.Id)
	if err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	medicalInfo := medical_folder.CreateMedicalInfo(newPatient.MedicalInfo, patient.Id)
	if medicalInfo.Err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("unable to create medical information" + medicalInfo.Err.Error())}
	}

	doctor, err := graphql.GetDoctorById(doctorID)
	if err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdateDoctor(doctorID, model.UpdateDoctorInput{
		PatientIds: append(doctor.PatientIds, &patient.Id),
	})
	if err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("update failed: " + err.Error())}
	}

	return PatientByIdResponse{
		Patient: model.Patient{
			ID:                   getPatient.ID,
			Email:                getPatient.Email,
			MedicalInfoID:        &medicalInfo.MedicalInfo.ID,
			RendezVousIds:        getPatient.RendezVousIds,
			DocumentIds:          getPatient.DocumentIds,
			TreatmentFollowUpIds: getPatient.TreatmentFollowUpIds,
		},
		MedicalInfo:                medicalInfo.MedicalInfo,
		AnteDiseasesWithTreatments: medicalInfo.AnteDiseasesWithTreatments,
		Code:                       201,
		Err:                        nil,
	}
}
