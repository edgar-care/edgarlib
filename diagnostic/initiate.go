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
	patientInfos, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id)

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
	input.AnteChirs = []string{}
	input.AnteDiseases = []string{}
	input.Treatments = []string{}

	utils.WakeNlpUp()

	session, err := graphql.CreateSession(context.Background(), gqlClient, []graphql.SessionDiseasesInput{}, []graphql.SessionSymptomInput{}, input.Age, input.Height, input.Weight, input.Sex, input.AnteChirs, input.AnteDiseases, input.Treatments, "", []graphql.LogsInput{}, []string{})

	if err != nil {
		return InitiateResponse{"", 500, errors.New("unable to create session")}
	}
	return InitiateResponse{session.CreateSession.Id, 200, nil}
}
