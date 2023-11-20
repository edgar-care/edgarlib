package graphql

import (
	"github.com/jinzhu/copier"
)

/********** Types ***********/

type Info struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birthdate"`
	Sex       string `json:"sex"`
	Weight    int    `json:"weight"`
	Height    int    `json:"height"`
}

type InfoOutput struct {
	Id        string  `json:"id"`
	Name      *string `json:"name"`
	Surname   *string `json:"surname"`
	BirthDate *string `json:"birthdate"`
	Sex       *string `json:"sex"`
	Weight    *int    `json:"weight"`
	Height    *int    `json:"height"`
}

type InfoInput struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birthdate"`
	Sex       string `json:"sex"`
	Weight    int    `json:"weight"`
	Height    int    `json:"height"`
}

/**************** GraphQL types *****************/

type createInfoResponse struct {
	Content InfoOutput `json:"createInfo"`
}

type getInfoByIdResponse struct {
	Content InfoOutput `json:"getInfoById"`
}

type updateInfoResponse struct {
	Content InfoOutput `json:"updateInfo"`
}

func CreateInfo(newInfo InfoInput) (Info, error) {
	var info createInfoResponse
	var resp Info
	query := `mutation createInfo($name: String!, $surname: String!, $birthdate: String!, $height: Int!, $weight: Int!, $sex: String!) {
            createInfo(name:$name, surname:$surname, birthdate:$birthdate, height:$height, weight:$weight, sex:$sex) {
                    id,
					name,
					birthdate,
					height,
					weight,
					sex,
					surname
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":      newInfo.Name,
		"birthdate": newInfo.BirthDate,
		"height":    newInfo.Height,
		"weight":    newInfo.Weight,
		"sex":       newInfo.Sex,
		"surname":   newInfo.Surname,
	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}

func GetInfoById(id string) (Info, error) {
	var info getInfoByIdResponse
	var resp Info
	query := `query getInfoById($id: String!) {
				getInfoById(id: $id) {
                    id,
					name,
					birthdate,
					height,
					weight,
					sex,
					surname
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}

func AddOnboardingInfoID(updatePatient UpdatePatientOnboardingInput) (UpdatePatientOnboarding, error) {
	var patient updateOnboardingResponse
	var resp UpdatePatientOnboarding
	query := `mutation updatePatient($id: String!, $onboarding_info_id: String) {
		updatePatient(id:$id, onboarding_info_id:$onboarding_info_id) {
                    id,
					onboarding_info_id
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":                 updatePatient.Id,
		"onboarding_info_id": updatePatient.OnboardingInfoID,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}
