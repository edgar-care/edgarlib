package auth

import (
	"crypto/rand"
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

type CreateSaveCodeResponse struct {
	SaveCode model.SaveCode
	Code     int
	Err      error
}

func generateRandomCode(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func generateBackupCodes() ([]string, error) {
	const codeLength = 8
	const numCodes = 10
	codes := make([]string, numCodes)

	for i := 0; i < numCodes; i++ {
		code, err := generateRandomCode(codeLength)
		if err != nil {
			return nil, err
		}
		codes[i] = code
	}

	return codes, nil
}

func hashCode(code string) string {
	salt, _ := strconv.Atoi(os.Getenv("SALT"))
	bytes, _ := bcrypt.GenerateFromPassword([]byte(code), salt)
	return string(bytes)
}

func CreateBackupCodes(id string, r *http.Request) CreateSaveCodeResponse {
	var doubleAuthId *string

	codes, err := generateBackupCodes()
	if err != nil {
		return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
	}

	hashedCodes := make([]string, len(codes))
	for i, code := range codes {
		hashedCodes[i] = hashCode(code)
	}

	saveCode, err := graphql.CreateSaveCode(model.CreateSaveCodeInput{Code: hashedCodes})
	if err != nil {
		return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 400, Err: errors.New("unable to create save code")}
	}

	token := GetBearerToken(r)
	accountType, err := GetAccountType(token)
	if accountType == "" {
		return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 400, Err: errors.New("no account type found")}
	}

	if accountType == "patient" {
		patient, err := graphql.GetPatientById(id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient")}
		}
		doubleAuthId = patient.DoubleAuthMethodsID
	} else {
		doctor, err := graphql.GetDoctorById(id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a doctor")}
		}
		doubleAuthId = doctor.DoubleAuthMethodsID
	}

	if doubleAuthId != nil {
		test, err := graphql.GetDoubleAuthById(*doubleAuthId)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
		}

		_, err = graphql.UpdateDoubleAuth(*doubleAuthId, model.UpdateDoubleAuthInput{
			Methods:       test.Methods,
			Secret:        &saveCode.ID,
			TrustDeviceID: test.TrustDeviceID,
		})
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
		}
	} else {
		newDoubleAuth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
			Methods:       []string{},
			Secret:        saveCode.ID,
			TrustDeviceID: []string{},
		})
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
		}

		if accountType == "patient" {
			_, err = graphql.UpdatePatient(id, model.UpdatePatientInput{DoubleAuthMethodsID: &newDoubleAuth.ID})
			if err != nil {
				return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
			}
		} else {
			_, err = graphql.UpdateDoctor(id, model.UpdateDoctorInput{DoubleAuthMethodsID: &newDoubleAuth.ID})
			if err != nil {
				return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
			}
		}
	}

	return CreateSaveCodeResponse{
		SaveCode: model.SaveCode{Code: codes},
		Code:     201,
		Err:      nil,
	}
}
