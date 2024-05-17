package dashboard

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/auth"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"github.com/edgar-care/edgarlib/medical_folder"
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

func CreatePatientFormDoctor(newPatient CreatePatientInput, doctorID string) PatientByIdResponse {
	gqlClient := graphql.CreateClient()

	patient := auth.CreatePatientAccount(newPatient.Email)
	if patient.Err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("Unable to create a new account")}
	}

	getPatient, err := graphql.GetPatientById(context.Background(), gqlClient, patient.Id)
	if err != nil {
		if err != nil {
			return PatientByIdResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
		}
	}
	medicalInfo := medical_folder.CreateMedicalInfo(newPatient.MedicalInfo, patient.Id)
	if medicalInfo.Err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("Unable to create a medical Information")}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorID)
	if err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctorID, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, append(doctor.GetDoctorById.Patient_ids, patient.Id), graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids)
	if err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	return PatientByIdResponse{
		Patient: model.Patient{
			ID:                   getPatient.GetPatientById.Id,
			Email:                getPatient.GetPatientById.Email,
			MedicalInfoID:        &medicalInfo.MedicalInfo.ID,
			RendezVousIds:        graphql.ConvertStringSliceToPointerSlice(getPatient.GetPatientById.Rendez_vous_ids),
			DocumentIds:          graphql.ConvertStringSliceToPointerSlice(getPatient.GetPatientById.Document_ids),
			TreatmentFollowUpIds: graphql.ConvertStringSliceToPointerSlice(getPatient.GetPatientById.Treatment_follow_up_ids),
		},
		MedicalInfo:                medicalInfo.MedicalInfo,
		AnteDiseasesWithTreatments: medicalInfo.AnteDiseasesWithTreatments,
		Code:                       201,
		Err:                        nil,
	}
}
