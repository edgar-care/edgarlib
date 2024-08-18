package auth

import (
	"fmt"
	"net/http"

	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type DoctorInput struct {
	Password  string       `jon:"password"`
	Email     string       `json:"email"`
	Name      string       `json:"name"`
	Firstname string       `json:"firstname"`
	Address   AddressInput `json:"address"`
}

type AddressInput struct {
	Street  string `json:"street"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	City    string `json:"city"`
}

type PatientInput struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AdminInput struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type RegisterAndLoginResponse struct {
	Token string
	Code  int
	Err   error
}

func RegisterDoctor(email string, password string, name string, firstname string, address AddressInput) (model.Doctor, error) {
	password = utils.HashPassword(password)
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     email,
		Password:  password,
		Name:      name,
		Firstname: firstname,
		Address: &model.AddressInput{
			Street:  address.Street,
			ZipCode: address.ZipCode,
			Country: address.Country,
			City:    address.City,
		},
		Status: true,
	})
	if err != nil {
		return model.Doctor{}, fmt.Errorf("Unable to create account: %s", err.Error())
	}

	return doctor, nil
}

func RegisterAndLoginDoctor(email string, password string, name string, firstname string, address AddressInput, nameDevice string) RegisterAndLoginResponse {
	doctor, err := RegisterDoctor(email, password, name, firstname, address)
	if err != nil {
		return RegisterAndLoginResponse{"", http.StatusBadRequest, err}
	}
	token, err := utils.CreateToken(map[string]interface{}{
		"doctor":      doctor.Email,
		"id":          doctor.ID,
		"name_device": nameDevice,
	})
	return RegisterAndLoginResponse{token, 200, err}
}

func RegisterPatient(email string, password string) (model.Patient, error) {
	password = utils.HashPassword(password)
	patient, err := graphql.CreatePatient(model.CreatePatientInput{Email: email, Password: password, Status: true})
	if err != nil {
		return model.Patient{}, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return patient, nil
}

func RegisterAndLoginPatient(email string, password string, nameDevice string) RegisterAndLoginResponse {
	patient, err := RegisterPatient(email, password)
	if err != nil {
		return RegisterAndLoginResponse{"", http.StatusBadRequest, err}
	}
	token, err := utils.CreateToken(map[string]interface{}{
		"patient":     patient.Email,
		"id":          patient.ID,
		"name_device": nameDevice,
	})
	return RegisterAndLoginResponse{token, 200, err}
}

func RegisterAdmin(email string, password string, firstName string, lastName string, token string) (model.Admin, error) {
	if utils.VerifyToken(token) == false {
		return model.Admin{}, fmt.Errorf("Unable to create account: Invalid Token")
	}
	password = utils.HashPassword(password)
	admin, err := graphql.CreateAdmin(model.CreateAdminInput{
		Email:    email,
		Password: password,
		Name:     firstName,
		LastName: lastName,
	})
	if err != nil {
		return model.Admin{}, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return admin, nil
}

func RegisterAndLoginAdmin(email string, password string, firstName string, lastName string, token string) RegisterAndLoginResponse {
	admin, err := RegisterAdmin(email, password, firstName, lastName, token)
	if err != nil {
		return RegisterAndLoginResponse{"", 400, err}
	}
	authToken, err := utils.CreateToken(map[string]interface{}{
		"admin": admin,
	})
	return RegisterAndLoginResponse{authToken, 200, err}
}
