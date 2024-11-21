package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Login2faSaveCodeInput struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	BackupCode string `json:"backup_code"`
}

type Login2faSaveCodeesponse struct {
	Token string
	Code  int
	Err   error
}

func Login2faSaveCode(input Login2faSaveCodeInput, nameDevice string, accountID string) Login2faSaveCodeesponse {
	var token string
	var doubleAuthId *string

	patient, patientErr := graphql.GetPatientById(accountID)
	if patientErr == nil {
		doubleAuthId = patient.DoubleAuthMethodsID
		if doubleAuthId == nil {
			return Login2faSaveCodeesponse{Token: "", Code: 400, Err: errors.New("no 2FA method associated with")}
		}
		doubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthId)
		if err != nil {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("get double_auth failed: " + err.Error()),
			}
		}

		saveCodeTable, err := graphql.GetSaveCodeById(doubleAuth.Secret)
		if err != nil {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("get save code failed: " + err.Error()),
			}
		}

		isValidBackupCode := false
		for _, hashedCode := range saveCodeTable.Code {
			err := bcrypt.CompareHashAndPassword([]byte(hashedCode), []byte(input.BackupCode))
			if err == nil {
				isValidBackupCode = true
				break
			}
		}

		if !isValidBackupCode {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("invalid backup code"),
			}
		}

		NewCodeList := removeHashedCode(saveCodeTable.Code, input.BackupCode)
		_, err = graphql.UpdateSaveCode(saveCodeTable.ID, model.UpdateSaveCodeInput{
			Code: NewCodeList,
		})
		if err != nil {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusInternalServerError,
				Err:   errors.New("update save code failed: " + err.Error()),
			}
		}

		check := CheckPassword(input.Password, patient.Password)
		if !check {
			return Login2faSaveCodeesponse{Token: "", Code: 401, Err: errors.New("invalid password")}
		}
		token, _ := CreateToken(map[string]interface{}{
			"patient":     patient.Email,
			"id":          patient.ID,
			"name_device": nameDevice,
		})

		return Login2faSaveCodeesponse{Token: token, Code: 200, Err: nil}
	}

	doctor, doctorErr := graphql.GetDoctorById(accountID)
	if doctorErr == nil {
		doubleAuthId = doctor.DoubleAuthMethodsID
		if doubleAuthId == nil {
			return Login2faSaveCodeesponse{Token: "", Code: 400, Err: errors.New("no 2FA method associated with")}
		}
		doubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthId)
		if err != nil {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("get double_auth failed: " + err.Error()),
			}
		}

		saveCodeTable, err := graphql.GetSaveCodeById(doubleAuth.Secret)
		if err != nil {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("get save code failed: " + err.Error()),
			}
		}

		isValidBackupCode := false
		for _, hashedCode := range saveCodeTable.Code {
			err := bcrypt.CompareHashAndPassword([]byte(hashedCode), []byte(input.BackupCode))
			if err == nil {
				isValidBackupCode = true
				break
			}
		}

		if !isValidBackupCode {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("invalid backup code"),
			}
		}

		NewCodeList := removeHashedCode(saveCodeTable.Code, input.BackupCode)

		_, err = graphql.UpdateSaveCode(saveCodeTable.ID, model.UpdateSaveCodeInput{
			Code: NewCodeList,
		})
		if err != nil {
			return Login2faSaveCodeesponse{
				Token: "",
				Code:  http.StatusInternalServerError,
				Err:   errors.New("update save code failed: " + err.Error()),
			}
		}

		check := CheckPassword(input.Password, doctor.Password)
		if !check {
			return Login2faSaveCodeesponse{Token: "", Code: 401, Err: errors.New("invalid password")}
		}
		token, _ = CreateToken(map[string]interface{}{
			"doctor":      doctor.Email,
			"id":          doctor.ID,
			"name_device": nameDevice,
		})

		return Login2faSaveCodeesponse{Token: token, Code: 200, Err: nil}
	}

	return Login2faSaveCodeesponse{Token: "", Code: 400, Err: errors.New("email does not correspond to a valid patient or doctor")}
}

func removeHashedCode(hashedCodes []string, inputCode string) []string {
	for i, hashedCode := range hashedCodes {
		err := bcrypt.CompareHashAndPassword([]byte(hashedCode), []byte(inputCode))
		if err == nil {
			return append(hashedCodes[:i], hashedCodes[i+1:]...)
		}
	}
	return hashedCodes
}
