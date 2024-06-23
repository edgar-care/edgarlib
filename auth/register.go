package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/edgar-care/edgarlib/auth/utils"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()
	password = utils.HashPassword(password)
	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, email, password, firstname, name, graphql.AddressInput{Street: address.Street, Zip_code: address.ZipCode, Country: address.Country, City: address.City})
	if err != nil {
		return model.Doctor{}, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	doctorAddress := &model.Address{
		Street:  address.Street,
		ZipCode: address.ZipCode,
		Country: address.Country,
		City:    address.City,
	}
	return model.Doctor{
		ID:            doctor.CreateDoctor.Id,
		Email:         doctor.CreateDoctor.Email,
		Password:      doctor.CreateDoctor.Password,
		Name:          doctor.CreateDoctor.Name,
		Firstname:     doctor.CreateDoctor.Firstname,
		Address:       doctorAddress,
		RendezVousIds: graphql.ConvertStringSliceToPointerSlice(doctor.CreateDoctor.Rendez_vous_ids),
		PatientIds:    graphql.ConvertStringSliceToPointerSlice(doctor.CreateDoctor.Patient_ids),
	}, nil
}

func RegisterAndLoginDoctor(email string, password string, name string, firstname string, address AddressInput) RegisterAndLoginResponse {
	doctor, err := RegisterDoctor(email, password, name, firstname, address)
	if err != nil {
		return RegisterAndLoginResponse{"", http.StatusBadRequest, err}
	}
	token, err := utils.CreateToken(map[string]interface{}{
		"doctor": doctor,
	})
	return RegisterAndLoginResponse{token, 200, err}
}

func RegisterPatient(email string, password string) (model.Patient, error) {
	gqlClient := graphql.CreateClient()
	password = utils.HashPassword(password)
	patient, err := graphql.CreatePatient(context.Background(), gqlClient, email, password)
	if err != nil {
		return model.Patient{}, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return model.Patient{
		ID:                   patient.CreatePatient.Id,
		Email:                patient.CreatePatient.Email,
		Password:             patient.CreatePatient.Password,
		MedicalInfoID:        &patient.CreatePatient.Medical_info_id,
		RendezVousIds:        graphql.ConvertStringSliceToPointerSlice(patient.CreatePatient.Rendez_vous_ids),
		DocumentIds:          graphql.ConvertStringSliceToPointerSlice(patient.CreatePatient.Document_ids),
		TreatmentFollowUpIds: graphql.ConvertStringSliceToPointerSlice(patient.CreatePatient.Treatment_follow_up_ids),
	}, nil
}

func RegisterAndLoginPatient(email string, password string, ip string) RegisterAndLoginResponse {
	patient, err := RegisterPatient(email, password)
	if err != nil {
		return RegisterAndLoginResponse{"", http.StatusBadRequest, err}
	}
	token, err := utils.CreateToken(map[string]interface{}{
		"patient":   patient.Email,
		"id":        patient.ID,
		"ip_device": ip,
	})
	return RegisterAndLoginResponse{token, 200, err}
}

func RegisterAdmin(email string, password string, firstName string, lastName string, token string) (model.Admin, error) {
	if utils.VerifyToken(token) == false {
		return model.Admin{}, fmt.Errorf("Unable to create account: Invalid Token")
	}
	gqlClient := graphql.CreateClient()
	password = utils.HashPassword(password)
	admin, err := graphql.CreateAdmin(context.Background(), gqlClient, email, password, firstName, lastName)
	if err != nil {
		return model.Admin{}, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return model.Admin{
		ID:       admin.CreateAdmin.Id,
		Email:    admin.CreateAdmin.Email,
		Password: admin.CreateAdmin.Password,
		Name:     admin.CreateAdmin.Name,
		LastName: admin.CreateAdmin.Last_name,
	}, nil
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
