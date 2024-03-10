package dashboard

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetPatientByIdResponse struct {
	Patient     model.Patient
	MedicalInfo model.MedicalInfo
	Code        int
	Err         error
}

type PatientWithMedicalInfo struct {
	ID            string            `json:"id"`
	Email         string            `json:"email"`
	MedicalInfo   model.MedicalInfo `json:"medical_info"`
	RendezVousIds []string          `json:"rendez_vous_ids"`
	DocumentsIds  []string          `json:"documents_ids"`
}

type GetPatientsResponse struct {
	PatientsInfo []PatientWithMedicalInfo `json:"patients_info"`
	Code         int
	Err          error
}

func GetPatientById(id string, doctorid string) GetPatientByIdResponse {
	gqlClient := graphql.CreateClient()

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorid)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}
	for _, item := range doctor.GetDoctorById.Patient_ids {
		if item != id {
			return GetPatientByIdResponse{Code: 400, Err: errors.New("unauthorized to access to this account")}
		}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err != nil {
		return GetPatientByIdResponse{Code: 401, Err: errors.New("id does not correspond to a patient")}
	}

	if patient.GetPatientById.Medical_info_id == "" {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("onboarding not started")}
	}

	medicalInfo, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to an medical table")}
	}

	medicalAntecedentsResp := make([]*model.MedicalAntecedents, len(medicalInfo.GetMedicalFolderById.Medical_antecedents))
	for i, antecedent := range medicalInfo.GetMedicalFolderById.Medical_antecedents {
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

	return GetPatientByIdResponse{
		Patient: model.Patient{
			ID:            patient.GetPatientById.Id,
			Email:         patient.GetPatientById.Email,
			Password:      patient.GetPatientById.Password,
			MedicalInfoID: &patient.GetPatientById.Medical_info_id,
			RendezVousIds: graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Rendez_vous_ids),
			DocumentIds:   graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Document_ids),
		},
		MedicalInfo: model.MedicalInfo{
			ID:                 medicalInfo.GetMedicalFolderById.Id,
			Name:               medicalInfo.GetMedicalFolderById.Name,
			Firstname:          medicalInfo.GetMedicalFolderById.Firstname,
			Birthdate:          medicalInfo.GetMedicalFolderById.Birthdate,
			Sex:                model.Sex(medicalInfo.GetMedicalFolderById.Sex),
			Weight:             medicalInfo.GetMedicalFolderById.Weight,
			Height:             medicalInfo.GetMedicalFolderById.Height,
			PrimaryDoctorID:    medicalInfo.GetMedicalFolderById.Primary_doctor_id,
			OnboardingStatus:   model.OnboardingStatus(medicalInfo.GetMedicalFolderById.Onboarding_status),
			MedicalAntecedents: medicalAntecedentsResp,
		},
		Code: 200,
		Err:  nil,
	}
}

func GetPatients(doctorId string) GetPatientsResponse {
	gqlClient := graphql.CreateClient()

	patientDoctor, err := graphql.GetPatientsFromDoctorById(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetPatientsResponse{Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	var patients []PatientWithMedicalInfo
	for _, patient := range patientDoctor.GetPatientsFromDoctorById {
		patientId := patient.Medical_info_id
		medicalInfo, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patientId)
		if err != nil {
			return GetPatientsResponse{Code: 400, Err: errors.New("id does not correspond to an medical table")}
		}

		medicalAntecedentsResp := make([]*model.MedicalAntecedents, len(medicalInfo.GetMedicalFolderById.Medical_antecedents))
		for i, antecedent := range medicalInfo.GetMedicalFolderById.Medical_antecedents {
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
		patients = append(patients, PatientWithMedicalInfo{
			ID:            patient.Id,
			Email:         patient.Email,
			RendezVousIds: patient.Rendez_vous_ids,
			DocumentsIds:  patient.Document_ids,
			MedicalInfo: model.MedicalInfo{
				ID:                 medicalInfo.GetMedicalFolderById.Id,
				Name:               medicalInfo.GetMedicalFolderById.Name,
				Firstname:          medicalInfo.GetMedicalFolderById.Firstname,
				Birthdate:          medicalInfo.GetMedicalFolderById.Birthdate,
				Sex:                model.Sex(medicalInfo.GetMedicalFolderById.Sex),
				Height:             medicalInfo.GetMedicalFolderById.Height,
				Weight:             medicalInfo.GetMedicalFolderById.Weight,
				PrimaryDoctorID:    medicalInfo.GetMedicalFolderById.Primary_doctor_id,
				OnboardingStatus:   model.OnboardingStatus(medicalInfo.GetMedicalFolderById.Onboarding_status),
				MedicalAntecedents: medicalAntecedentsResp,
			},
		})
	}
	return GetPatientsResponse{
		PatientsInfo: patients,
	}
}
