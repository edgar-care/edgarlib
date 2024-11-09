package dashboard

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
)

type CreatePatientInput struct {
	Email       string                                   `json:"email"`
	MedicalInfo medical_folder.CreateNewMedicalInfoInput `json:"medical_info"`
}

type PatientByIdResponse struct {
	Patient                    model.Patient
	MedicalInfo                model.MedicalInfo
	AnteDiseasesWithTreatments []model.MedicalAntecedents
	Code                       int
	Err                        error
}

func CreatePatientFromDoctor(doctorID string, newPatient CreatePatientInput) PatientByIdResponse {
	_, err := graphql.GetPatientByEmail(newPatient.Email)
	if err == nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("patient already exists")}
	}

	patient := auth.CreatePatientAccount(newPatient.Email)
	if patient.Err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("unable to create a new account: " + patient.Err.Error())}
	}

	getPatient, err := graphql.GetPatientById(patient.Id)
	if err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	medicalInfo := medical_folder.NewMedicalFolder(newPatient.MedicalInfo, patient.Id)
	if medicalInfo.Err != nil {
		return PatientByIdResponse{Code: 400, Err: errors.New("unable to create medical information: " + medicalInfo.Err.Error())}
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
		AnteDiseasesWithTreatments: medicalInfo.MedicalAntecedents,
		Code:                       201,
		Err:                        nil,
	}

}
