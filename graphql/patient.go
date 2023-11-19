package graphql

import "github.com/jinzhu/copier"

/********** Types ***********/

type Patient struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type PatientOutput struct {
	Id       *string `json:"id"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

type PatientInput struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

/**************** GraphQL types *****************/

type getPatientByEmailResponse struct {
	Content PatientOutput `json:"getPatientByEmail"`
}

type getPatientByIdResponse struct {
	Content PatientOutput `json:"getPatientById"`
}

type createPatientResponse struct {
	Content PatientOutput `json:"createPatient"`
}

/*************** Implementations *****************/

func GetPatientById(id string) (Patient, error) {
	var patient getPatientByIdResponse
	var resp Patient
	query := `query getPatientByID($id: String!) {
                getPatientByID(id: $id) {
                    id,
                    password,
                    email,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func GetPatientByEmail(email string) (Patient, error) {
	var patient getPatientByEmailResponse
	var resp Patient
	query := `query getPatientByEmail($email: String!) {
                getPatientByEmail(email: $email) {
                    id,
                    password,
                    email,
                    }
                }`

	err := Query(query, map[string]interface{}{
		"email": email,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func CreatePatient(newPatient PatientInput) (Patient, error) {
	var patient createPatientResponse
	var resp Patient
	query := `mutation createPatient($email: String!, $password: String!) {
            createPatient(email:$email, password:$password) {
                    id,
                    password,
                    email,
                }
            }`
	err := Query(query, map[string]interface{}{
		"email":    newPatient.Email,
		"password": newPatient.Password,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}
