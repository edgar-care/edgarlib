package medical_folder

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type CreateMedicalInfoInput struct {
	Name                   string                         `json:"name"`
	Firstname              string                         `json:"firstname"`
	Birthdate              int                            `json:"birthdate"`
	Sex                    string                         `json:"sex"`
	Weight                 int                            `json:"weight"`
	Height                 int                            `json:"height"`
	PrimaryDoctorID        string                         `json:"primary_doctor_id,omitempty"`
	MedicalAntecedents     []CreateMedicalAntecedentInput `json:"medical_antecedents"`
	FamilyMembersMedInfoId []string                       `json:"family_members_med_info_id"`
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

func ConvertPeriodsAndDays(periods []string, days []string) ([]model.Period, []model.Day) {
	convertedPeriods := make([]model.Period, len(periods))
	for i, p := range periods {
		convertedPeriods[i] = model.Period(p)
	}

	convertedDays := make([]model.Day, len(days))
	for i, d := range days {
		convertedDays[i] = model.Day(d)
	}
	return convertedPeriods, convertedDays
}

func CreateMedicalInfo(input CreateMedicalInfoInput, patientID string) CreateMedicalInfoResponse {
	var antdediseaseids []string
	var antediseasesWithTreatments []AnteDiseaseWithTreatments

	control, err := graphql.GetPatientById(patientID)
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if control.MedicalInfoID != nil && *control.MedicalInfoID != "" {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("medical folder has already been created")}
	}

	if len(input.MedicalAntecedents) == 0 {

		medical, err := graphql.CreateMedicalFolder(model.CreateMedicalFolderInput{
			Name:                   input.Name,
			Firstname:              input.Firstname,
			Birthdate:              input.Birthdate,
			Sex:                    input.Sex,
			Height:                 input.Height,
			Weight:                 input.Weight,
			PrimaryDoctorID:        input.PrimaryDoctorID,
			OnboardingStatus:       "DONE",
			FamilyMembersMedInfoID: input.FamilyMembersMedInfoId,
		})
		if err != nil {
			return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create medical folder: " + err.Error())}
		}

		_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{
			MedicalInfoID: &medical.ID,
		})
		if err != nil {
			return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to update patient: " + err.Error())}
		}

		return CreateMedicalInfoResponse{
			MedicalInfo: medical,
			Code:        201,
			Err:         nil,
		}
	}

	for _, antecedent := range input.MedicalAntecedents {
		var treatmentIDsPerAnte = []string{}
		var antediseaseTreatments []model.Treatment

		for _, medicine := range antecedent.Medicines {
			periods, days := ConvertPeriodsAndDays(medicine.Period, medicine.Day)
			treatmentRes, err := graphql.CreateTreatment(model.CreateTreatmentInput{Period: periods, Day: days, Quantity: medicine.Quantity, MedicineID: medicine.MedicineID})
			if err != nil {
				return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
			}

			treatmentIDsPerAnte = append(treatmentIDsPerAnte, treatmentRes.ID)
			antediseaseTreatments = append(antediseaseTreatments, treatmentRes)
		}

		chronicity := 0.0
		antedisease, err := graphql.CreateAnteDisease(model.CreateAnteDiseaseInput{
			Name:          antecedent.Name,
			Chronicity:    &chronicity,
			TreatmentIds:  treatmentIDsPerAnte,
			SurgeryIds:    []string{},
			Symptoms:      []string{},
			StillRelevant: antecedent.StillRelevant,
		})
		if err != nil {
			return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create antedisease: " + err.Error())}
		}

		antdediseaseids = append(antdediseaseids, antedisease.ID)

		antediseaseWithTreatments := AnteDiseaseWithTreatments{
			AnteDisease: antedisease,
			Treatments:  antediseaseTreatments,
		}

		antediseasesWithTreatments = append(antediseasesWithTreatments, antediseaseWithTreatments)
	}

	medical, err := graphql.CreateMedicalFolder(model.CreateMedicalFolderInput{
		Name:                   input.Name,
		Firstname:              input.Firstname,
		Birthdate:              input.Birthdate,
		Sex:                    input.Sex,
		Height:                 input.Height,
		Weight:                 input.Weight,
		AntecedentDiseaseIds:   antdediseaseids,
		PrimaryDoctorID:        input.PrimaryDoctorID,
		OnboardingStatus:       "DONE",
		FamilyMembersMedInfoID: input.FamilyMembersMedInfoId,
	})
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create medical folder: " + err.Error())}
	}

	_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{MedicalInfoID: &medical.ID})
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to update patient: " + err.Error())}
	}

	return CreateMedicalInfoResponse{
		MedicalInfo:                medical,
		AnteDiseasesWithTreatments: antediseasesWithTreatments,
		Code:                       201,
		Err:                        nil,
	}
}
