package diagnostic

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/diagnostic/utils"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type InitiateResponse struct {
	Id   string
	Code int
	Err  error
}

func Initiate(id string) InitiateResponse {
	gqlClient := graphql.CreateClient()
	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err != nil {
		return InitiateResponse{"", 400, errors.New("unable to get patient")}
	}

	patientInfos, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id)
	if err != nil {
		return InitiateResponse{"", 400, errors.New("unable to get patient medical infos")}
	}

	var input model.Session
	input.Age = patientInfos.GetMedicalFolderById.Birthdate
	input.Height = patientInfos.GetMedicalFolderById.Height
	input.Weight = patientInfos.GetMedicalFolderById.Weight
	if patientInfos.GetMedicalFolderById.Sex == graphql.SexMale {
		input.Sex = "M"
	} else if patientInfos.GetMedicalFolderById.Sex == graphql.SexFemale {
		input.Sex = "F"
	} else {
		input.Sex = "O"
	}
	input.AnteDiseases = patientInfos.GetMedicalFolderById.Antecedent_disease_ids

	if input.AnteDiseases != nil && len(input.AnteDiseases) > 0 {
		for _, anteDiseaseID := range input.AnteDiseases {
			if anteDiseaseID != "" {
				ante, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, anteDiseaseID)
				if err != nil {
					return InitiateResponse{"", 500, errors.New("problem with anteDisease ID")}
				}
				if ante.GetAnteDiseaseByID.Surgery_ids != nil && len(ante.GetGetAnteDiseaseByID().Surgery_ids) > 0 && ante.GetAnteDiseaseByID.Still_relevant == true {
					for _, anteChirId := range ante.GetAnteDiseaseByID.Surgery_ids {
						input.AnteChirs = append(input.AnteChirs, anteChirId)
					}
				} else {
					input.AnteChirs = []string{}
				}
			} else {
				input.AnteChirs = []string{}
			}
		}
	}

	input.Medicine = []string{}
	input.Medicine = append(input.Medicine, "CanonFlesh")

	for _, antecedentDiseaseId := range input.AnteDiseases {
		{
			if antecedentDiseaseId != "" {
				antecedentDisease, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, antecedentDiseaseId)
				if err != nil {
					return InitiateResponse{"", 500, errors.New("problem with antedisease ID")}
				}
				if antecedentDisease.GetAnteDiseaseByID.Still_relevant == true {
					for _, treatmentIds := range antecedentDisease.GetAnteDiseaseByID.Treatment_ids {
						treatment, err := graphql.GetTreatmentByID(context.Background(), gqlClient, treatmentIds)
						if err != nil {
							return InitiateResponse{"", 500, errors.New("problem with treatment ID")}
						}
						if treatment.GetTreatmentByID.Medicine_id != "" {
							input.Medicine = append(input.Medicine, treatment.GetTreatmentByID.Medicine_id)
						}
					}
				}
			}
		}
	}

	utils.WakeNlpUp()

	session, err := graphql.CreateSession(context.Background(), gqlClient, []graphql.SessionDiseasesInput{}, []graphql.SessionSymptomInput{}, input.Age, input.Height, input.Weight, input.Sex, input.AnteChirs, input.AnteDiseases, input.Medicine, "", []graphql.LogsInput{}, []string{})

	if err != nil {
		return InitiateResponse{"", 500, errors.New("unable to create session")}
	}
	return InitiateResponse{session.CreateSession.Id, 200, nil}
}
