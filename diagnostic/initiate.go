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
	input.MedicalAntecedents = patientInfos.AntecedentDiseaseIds

	input.Medicine = []string{}
	input.Medicine = append(input.Medicine, "CanonFlesh")

	for _, medicalAntecedentsId := range input.MedicalAntecedents {
		{
			if medicalAntecedentsId != "" {
				medicalAntecedent, err := graphql.GetMedicalAntecedentsById(medicalAntecedentsId)
				if err != nil {
					return InitiateResponse{"", 500, errors.New("problem with MedicalAntecedents ID")}
				}
				for _, AntecedentTreatment := range medicalAntecedent.Treatments {
					for _, AntecedentsMedicine := range AntecedentTreatment.Medicines {
						input.Medicine = append(input.Medicine, AntecedentsMedicine.MedicineID)
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
				familyMemberAnte, _ := graphql.GetMedicalAntecedentsById(familyMemberAnteId)
				for _, disease := range diseases {
					if familyMemberAnte.Name == disease.Name && disease.HeredityFactor != 0 {
						input.HereditaryDisease = append(input.HereditaryDisease, disease.Name)
					}
				}
			}
		}
	}

	utils.WakeNlpUp()

	session, err := graphql.CreateSession(model.CreateSessionInput{
		Age:                input.Age,
		Height:             input.Height,
		Weight:             input.Weight,
		Sex:                input.Sex,
		MedicalAntecedents: input.MedicalAntecedents,
		Medicine:           input.Medicine,
		HereditaryDisease:  input.HereditaryDisease,
	})

	if err != nil {
		return InitiateResponse{"", 500, errors.New("unable to create session")}
	}
	return InitiateResponse{session.ID, 200, nil}
}
