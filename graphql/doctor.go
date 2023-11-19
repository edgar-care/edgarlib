package graphql

import (
	"github.com/jinzhu/copier"
)

/********** Types ***********/

type Doctor struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type DoctorOutput struct {
	Id       *string `json:"id"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

type DoctorInput struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

/**************** GraphQL types *****************/

type getDoctorByIdResponse struct {
	Content DoctorOutput `json:"getDoctorById"`
}

type getDoctorByEmailResponse struct {
	Content DoctorOutput `json:"getDoctorByEmail"`
}

type createDoctorResponse struct {
	Content DoctorOutput `json:"createDoctor"`
}

/*************** Implementations *****************/

func GetDoctorById(id string) (Doctor, error) {
	var doctor getDoctorByIdResponse
	var resp Doctor
	query := `query getDoctorByID($id: String!) {
                getDoctorByID(id: $id) {
                    id,
                    password,
                    email,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}

func GetDoctorByEmail(email string) (Doctor, error) {
	var doctor getDoctorByEmailResponse
	var resp Doctor
	query := `query getDoctorByEmail($email: String!) {
                getDoctorByEmail(email: $email) {
                    id,
                    password,
                    email,
                }
            }`

	err := Query(query, map[string]interface{}{
		"email": email,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}

func CreateDoctor(newDoctor DoctorInput) (Doctor, error) {
	var doctor createDoctorResponse
	var resp Doctor
	query := `mutation createDoctor($email: String!, $password: String!) {
        createDoctor(email:$email, password:$password) {
                    id,
                    email,
                    password,
                }
            }`
	err := Query(query, map[string]interface{}{
		"email":    newDoctor.Email,
		"password": newDoctor.Password,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}
