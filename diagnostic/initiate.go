package diagnostic

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/diagnostic/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type InitiateResponse struct {
	Id   string
	Code int
	Err  error
}

func Initiate(id string) InitiateResponse {
	patient, err := graphql.GetPatientById(id)
	if err != nil {
		return InitiateResponse{"", 400, errors.New("unable to get patient")}
	}

	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return InitiateResponse{"", 400, errors.New("unable to get patient medical infos")}
	}
	patientInfos, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return InitiateResponse{"", 400, errors.New("unable to get patient medical infos")}
	}

	diseases, _ := graphql.GetDiseases(nil)

	var input model.Session
	input.Age = patientInfos.Birthdate
	input.Height = patientInfos.Height
	input.Weight = patientInfos.Weight
	if patientInfos.Sex == model.SexMale {
		input.Sex = "M"
	} else if patientInfos.Sex == model.SexFemale {
		input.Sex = "F"
	} else {
		input.Sex = "O"
	}
	input.AnteDiseases = patientInfos.AntecedentDiseaseIds

	input.AnteChirs = []string{}
	if input.AnteDiseases != nil && len(input.AnteDiseases) > 0 {
		for _, anteDiseaseID := range input.AnteDiseases {
			if anteDiseaseID != "" {
				ante, err := graphql.GetAnteDiseaseByID(anteDiseaseID)
				if err != nil {
					return InitiateResponse{"", 500, errors.New("problem with anteDisease ID")}
				}
				if ante.SurgeryIds != nil && len(ante.SurgeryIds) > 0 && ante.StillRelevant == true {
					for _, anteChirId := range ante.SurgeryIds {
						input.AnteChirs = append(input.AnteChirs, anteChirId)
					}
				}
			}
		}
	}

	input.Medicine = []string{}
	input.Medicine = append(input.Medicine, "CanonFlesh")

	for _, antecedentDiseaseId := range input.AnteDiseases {
		{
			if antecedentDiseaseId != "" {
				antecedentDisease, err := graphql.GetAnteDiseaseByID(antecedentDiseaseId)
				if err != nil {
					return InitiateResponse{"", 500, errors.New("problem with antedisease ID")}
				}
				if antecedentDisease.StillRelevant == true {
					for _, treatmentIds := range antecedentDisease.TreatmentIds {
						treatment, err := graphql.GetTreatmentByID(treatmentIds)
						if err != nil {
							return InitiateResponse{"", 500, errors.New("problem with treatment ID")}
						}
						if treatment.MedicineID != "" {
							input.Medicine = append(input.Medicine, treatment.MedicineID)
						}
					}
				}
			}
		}
	}

	input.HereditaryDisease = []string{}
	for _, familyMemberInfoId := range patientInfos.FamilyMembersMedInfoID {
		if familyMemberInfoId != "" {
			familyMemberInfo, _ := graphql.GetMedicalFolderByID(familyMemberInfoId)
			for _, familyMemberAnteId := range familyMemberInfo.AntecedentDiseaseIds {
				familyMemberAnte, _ := graphql.GetAnteDiseaseByID(familyMemberAnteId)
				if familyMemberAnte.StillRelevant == true {
					for _, disease := range diseases {
						if familyMemberAnte.Name == disease.Name && disease.HeredityFactor != 0 {
							input.HereditaryDisease = append(input.HereditaryDisease, disease.Name)
						}
					}
				}
			}
		}
	}

	utils.WakeNlpUp()

	session, err := graphql.CreateSession(model.CreateSessionInput{
		Age:               input.Age,
		Height:            input.Height,
		Weight:            input.Weight,
		Sex:               input.Sex,
		AnteChirs:         input.AnteChirs,
		AnteDiseases:      input.AnteDiseases,
		Medicine:          input.Medicine,
		HereditaryDisease: input.HereditaryDisease,
	})

	if err != nil {
		return InitiateResponse{"", 500, errors.New("unable to create session")}
	}
	return InitiateResponse{session.ID, 200, nil}
}
