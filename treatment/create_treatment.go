package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateNewTreatmentInput struct {
	Name          string             `json:"name"`
	DiseaseId     string             `json:"disease_id"`
	Treatments    []CreateTreatInput `json:"treatments"`
	StillRelevant bool               `json:"still_relevant"`
}

type CreateTreatInput struct {
	MedicineId string   `json:"medicine_id"`
	Period     []string `json:"period"`
	Day        []string `json:"day"`
	Quantity   int      `json:"quantity"`
	StartDate  int      `json:"start_date"`
	EndDate    int      `json:"end_date"`
}

type CreateTreatmentResponse struct {
	AnteDisease model.AnteDisease
	Treatment   []model.Treatment
	Code        int
	Err         error
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

func CreateTreatment(input CreateNewTreatmentInput, patientID string) CreateTreatmentResponse {
	var res []model.Treatment

	control, err := graphql.GetPatientById(patientID)
	if err != nil {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if control.MedicalInfoID == nil || *control.MedicalInfoID == "" {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	for _, treat := range input.Treatments {
		periods, days := ConvertPeriodsAndDays(treat.Period, treat.Day)

		treatment, err := graphql.CreateTreatment(model.CreateTreatmentInput{
			Period:     periods,
			Day:        days,
			Quantity:   treat.Quantity,
			StartDate:  treat.StartDate,
			EndDate:    treat.EndDate,
			MedicineID: treat.MedicineId,
		})
		if err != nil {
			return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
		}
		res = append(res, treatment)
	}
	var treatmentIds []string
	for _, treatment := range res {
		treatmentIds = append(treatmentIds, treatment.ID)
	}

	if input.DiseaseId == "" {
		anteDisease, err := graphql.CreateAnteDisease(model.CreateAnteDiseaseInput{
			Name:          input.Name,
			TreatmentIds:  treatmentIds,
			StillRelevant: input.StillRelevant,
		})
		if err != nil {
			return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to create ante disease: " + err.Error())}
		}
		getMedic, err := graphql.GetMedicalFolderByID(*control.MedicalInfoID)
		if err != nil {
			return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to get medical folder: " + err.Error())}
		}

		_, err = graphql.UpdateMedicalFolder(*control.MedicalInfoID, model.UpdateMedicalFolderInput{
			AntecedentDiseaseIds: append(getMedic.AntecedentDiseaseIds, anteDisease.ID),
		})
		if err != nil {
			return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to update medical folder: " + err.Error())}
		}
		return CreateTreatmentResponse{
			AnteDisease: anteDisease,
			Treatment:   res,
			Code:        201,
			Err:         nil,
		}
	} else {
		checkAnte, err := graphql.GetAnteDiseaseByID(input.DiseaseId)
		if err != nil {
			return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to get ante disease: " + err.Error())}
		}
		anteDisease, err := graphql.UpdateAnteDisease(input.DiseaseId, model.UpdateAnteDiseaseInput{
			Name:         &input.Name,
			TreatmentIds: append(checkAnte.TreatmentIds, treatmentIds...),
		})
		if err != nil {
			return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to update ante disease: " + err.Error())}
		}
		return CreateTreatmentResponse{
			AnteDisease: anteDisease,
			Treatment:   res,
			Code:        201,
			Err:         nil,
		}
	}
}
