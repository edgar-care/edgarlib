package medical_folder

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetMedicalInfoByIdResponse struct {
	MedicalInfo model.MedicalInfo
	Code        int
	Err         error
}

func GetMedicalInfosById(id string) GetMedicalInfoByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.MedicalInfo

	id_med, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err != nil {
		return GetMedicalInfoByIdResponse{model.MedicalInfo{}, 400, errors.New("ID does not correspond to patient")}
	}

	if id_med.GetPatientById.Medical_info_id == "" {
		return GetMedicalInfoByIdResponse{model.MedicalInfo{}, 400, errors.New("ID not found")}
	}

	medical, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, id_med.GetPatientById.Medical_info_id)
	if err != nil {
		return GetMedicalInfoByIdResponse{model.MedicalInfo{}, 400, errors.New("ID does not correspond to any medical information")}
	}

	medicalAntecedentsResp := make([]*model.MedicalAntecedents, len(medical.GetMedicalFolderById.Medical_antecedents))
	for i, antecedent := range medical.GetMedicalFolderById.Medical_antecedents {
		medicines := make([]*model.Medicines, len(antecedent.Medicines))
		for j, med := range antecedent.Medicines {
			periods := make([]*model.Period, len(med.Period))
			for k, p := range med.Period {
				period := model.Period(p)
				periods[k] = &period
			}
			days := make([]*model.Day, len(med.Day))
			for k, d := range med.Day {
				day := model.Day(d)
				days[k] = &day
			}
			medicines[j] = &model.Medicines{
				Period:   periods,
				Day:      days,
				Quantity: med.Quantity,
			}
		}
		medicalAntecedentsResp[i] = &model.MedicalAntecedents{
			Name:          antecedent.Name,
			Medicines:     medicines,
			StillRelevant: antecedent.Still_relevant,
		}
	}

	res = model.MedicalInfo{
		ID:                 medical.GetMedicalFolderById.Id,
		Name:               medical.GetMedicalFolderById.Name,
		Firstname:          medical.GetMedicalFolderById.Firstname,
		Birthdate:          medical.GetMedicalFolderById.Birthdate,
		Sex:                model.Sex(medical.GetMedicalFolderById.Sex),
		Weight:             medical.GetMedicalFolderById.Weight,
		Height:             medical.GetMedicalFolderById.Height,
		PrimaryDoctorID:    medical.GetMedicalFolderById.Primary_doctor_id,
		MedicalAntecedents: medicalAntecedentsResp,
		OnboardingStatus:   model.OnboardingStatus(medical.GetMedicalFolderById.Onboarding_status),
	}
	return GetMedicalInfoByIdResponse{res, 200, nil}
}
