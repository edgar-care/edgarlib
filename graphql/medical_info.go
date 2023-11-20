package graphql

import (
	"github.com/jinzhu/copier"
)

/********** Types ***********/

type MedicalInfo struct {
	Info   InfoInput   `json:"onboarding_info"`
	Health HealthInput `json:"onboarding_health"`
}

/*************** Implementations *****************/

func UpdateMedicalInfo(id_info string, newInfo InfoInput) (Info, error) {
	var info updateInfoResponse
	var resp Info
	query := `mutation updateInfo($id: String! $name: String!, $surname: String!, $birthdate: String!, $height: Int!, $weight: Int!, $sex: String!) {
                updateInfo(id:$id, name:$name, surname:$surname, birthdate:$birthdate, height:$height, weight:$weight, sex:$sex) {
                    id,
					name,
					surname,
					birthdate,
					weight,
					sex,
					height
                }
            }`

	err := Query(query, map[string]interface{}{
		"id":        id_info,
		"name":      newInfo.Name,
		"surname":   newInfo.Surname,
		"birthdate": newInfo.BirthDate,
		"weight":    newInfo.Weight,
		"sex":       newInfo.Sex,
		"height":    newInfo.Height,
	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}

func UpdateMedicalHealth(id_health string, newHealth HealthInput) (Health, error) {
	var info updateHealthResponse
	var resp Health
	query := `mutation updateHealth($id: String!, $patients_allergies: [String!], $patients_illness: [String!], $patients_treatments: [String!], $patients_primary_doctor: String!) {
                updateHealth(id:$id, patients_allergies:$patients_allergies, patients_illness:$patients_illness, patients_treatments:$patients_treatments, patients_primary_doctor:$patients_primary_doctor) {
                    id,
					patients_allergies,
					patients_illness,
					patients_primary_doctor,
					patients_treatments
                }
            }`

	err := Query(query, map[string]interface{}{
		"id":                      id_health,
		"patients_allergies":      newHealth.PatientsAllergies,
		"patients_illness":        newHealth.PatientsIllness,
		"patients_treatments":     newHealth.PatientsTreatments,
		"patients_primary_doctor": newHealth.PatientsPrimaryDoctor,
	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}
