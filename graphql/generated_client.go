package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/edgar-care/edgarlib/graphql/model"
)

func CreateAnteDisease(input model.CreateAnteDiseaseInput) (model.AnteDisease, error) {
	query := `mutation CreateAnteDisease($input: CreateAnteDiseaseInput!){
	    createAnteDisease(input: $input){
	        id
	        name
	        chronicity
	        surgery_ids
	        symptoms
	        treatment_ids
	        still_relevant
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AnteDisease{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteDisease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteDisease{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteDisease{}, err
	}

	var result struct {
		Data struct {
			CreateAnteDisease model.AnteDisease `json:"createAnteDisease"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteDisease{}, err
	}

	return result.Data.CreateAnteDisease, nil
}

func CreateSymptom(input model.CreateSymptomInput) (model.Symptom, error) {
	query := `mutation CreateSymptom($input: CreateSymptomInput!){
	    createSymptom(input: $input){
	        id
	        code
	        name
	        chronic
	        symptom
	        advice
	        question
	        question_basic
	        question_duration
	        question_ante
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Symptom{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Symptom{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Symptom{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Symptom{}, err
	}

	var result struct {
		Data struct {
			CreateSymptom model.Symptom `json:"createSymptom"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Symptom{}, err
	}

	return result.Data.CreateSymptom, nil
}

func CreateDocument(input model.CreateDocumentInput) (model.Document, error) {
	query := `mutation CreateDocument($input: CreateDocumentInput!){
	    createDocument(input: $input){
	        id
	        owner_id
	        name
	        document_type
	        category
	        is_favorite
	        download_url
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Document{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Document{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Document{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Document{}, err
	}

	var result struct {
		Data struct {
			CreateDocument model.Document `json:"createDocument"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Document{}, err
	}

	return result.Data.CreateDocument, nil
}

func CreateAdmin(input model.CreateAdminInput) (model.Admin, error) {
	query := `mutation CreateAdmin($input: CreateAdminInput!){
	    createAdmin(input: $input){
	        id
	        email
	        password
	        name
	        last_name
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Admin{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Admin{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Admin{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Data struct {
			CreateAdmin model.Admin `json:"createAdmin"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	return result.Data.CreateAdmin, nil
}

func UpdateMedicalFolder(id string, input model.UpdateMedicalFolderInput) (model.MedicalInfo, error) {
	query := `mutation UpdateMedicalFolder($id: String!, $input: UpdateMedicalFolderInput!){
	    updateMedicalFolder(id: $id, input: $input){
	        id
	        name
	        firstname
	        birthdate
	        sex
	        height
	        weight
	        primary_doctor_id
	        onboarding_status
	        antecedent_disease_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.MedicalInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Data struct {
			UpdateMedicalFolder model.MedicalInfo `json:"updateMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	return result.Data.UpdateMedicalFolder, nil
}

func CreateMedicine(input model.CreateMedicineInput) (model.Medicine, error) {
	query := `mutation CreateMedicine($input: CreateMedicineInput!){
	    createMedicine(input: $input){
	        id
	        name
	        unit
	        target_diseases
	        treated_symptoms
	        side_effects
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Medicine{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Medicine{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Medicine{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Medicine{}, err
	}

	var result struct {
		Data struct {
			CreateMedicine model.Medicine `json:"createMedicine"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Medicine{}, err
	}

	return result.Data.CreateMedicine, nil
}

func UpdateTreatment(id string, input model.UpdateTreatmentInput) (model.Treatment, error) {
	query := `mutation UpdateTreatment($id: String!, $input: UpdateTreatmentInput!){
	    updateTreatment(id: $id, input: $input){
	        id
	        period
	        day
	        quantity
	        medicine_id
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Treatment{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Treatment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Treatment{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Treatment{}, err
	}

	var result struct {
		Data struct {
			UpdateTreatment model.Treatment `json:"updateTreatment"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Treatment{}, err
	}

	return result.Data.UpdateTreatment, nil
}

func CreateAnteFamily(input model.CreateAnteFamilyInput) (model.AnteFamily, error) {
	query := `mutation CreateAnteFamily($input: CreateAnteFamilyInput!){
	    createAnteFamily(input: $input){
	        id
	        name
	        disease
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AnteFamily{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteFamily{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteFamily{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteFamily{}, err
	}

	var result struct {
		Data struct {
			CreateAnteFamily model.AnteFamily `json:"createAnteFamily"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteFamily{}, err
	}

	return result.Data.CreateAnteFamily, nil
}

func UpdateTreatmentsFollowUp(id string, input model.UpdateTreatmentsFollowUpInput) (model.TreatmentsFollowUp, error) {
	query := `mutation UpdateTreatmentsFollowUp($id: String!, $input: UpdateTreatmentsFollowUpInput!){
	    updateTreatmentsFollowUp(id: $id, input: $input){
	        id
	        treatment_id
	        date
	        period
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.TreatmentsFollowUp{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	var result struct {
		Data struct {
			UpdateTreatmentsFollowUp model.TreatmentsFollowUp `json:"updateTreatmentsFollowUp"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	return result.Data.UpdateTreatmentsFollowUp, nil
}

func UpdateDocument(id string, input model.UpdateDocumentInput) (model.Document, error) {
	query := `mutation UpdateDocument($id: String!, $input: UpdateDocumentInput!){
	    updateDocument(id: $id, input: $input){
	        id
	        owner_id
	        name
	        document_type
	        category
	        is_favorite
	        download_url
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Document{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Document{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Document{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Document{}, err
	}

	var result struct {
		Data struct {
			UpdateDocument model.Document `json:"updateDocument"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Document{}, err
	}

	return result.Data.UpdateDocument, nil
}

func UpdateRdv(id string, input model.UpdateRdvInput) (model.Rdv, error) {
	query := `mutation UpdateRdv($id: String!, $input: UpdateRdvInput!){
	    updateRdv(id: $id, input: $input){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Rdv{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Rdv{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Data struct {
			UpdateRdv model.Rdv `json:"updateRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	return result.Data.UpdateRdv, nil
}

func CreateAlert(input model.CreateAlertInput) (model.Alert, error) {
	query := `mutation CreateAlert($input: CreateAlertInput!){
	    createAlert(input: $input){
	        id
	        name
	        sex
	        height
	        weight
	        symptoms
	        comment
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Alert{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Alert{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Alert{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Alert{}, err
	}

	var result struct {
		Data struct {
			CreateAlert model.Alert `json:"createAlert"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Alert{}, err
	}

	return result.Data.CreateAlert, nil
}

func CreatePatient(input model.CreatePatientInput) (model.Patient, error) {
	query := `mutation CreatePatient($input: CreatePatientInput!){
	    createPatient(input: $input){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Patient{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Patient{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Patient{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Data struct {
			CreatePatient model.Patient `json:"createPatient"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	return result.Data.CreatePatient, nil
}

func UpdatePatient(id string, input model.UpdatePatientInput) (model.Patient, error) {
	query := `mutation UpdatePatient($id: String!, $input: UpdatePatientInput!){
	    updatePatient(id: $id, input: $input){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Patient{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Patient{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Patient{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Data struct {
			UpdatePatient model.Patient `json:"updatePatient"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	return result.Data.UpdatePatient, nil
}

func CreateMedicalFolder(input model.CreateMedicalFolderInput) (model.MedicalInfo, error) {
	query := `mutation CreateMedicalFolder($input: CreateMedicalFolderInput!){
	    createMedicalFolder(input: $input){
	        id
	        name
	        firstname
	        birthdate
	        sex
	        height
	        weight
	        primary_doctor_id
	        onboarding_status
	        antecedent_disease_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.MedicalInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Data struct {
			CreateMedicalFolder model.MedicalInfo `json:"createMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	return result.Data.CreateMedicalFolder, nil
}

func UpdateAlert(id string, input model.UpdateAlertInput) (model.Alert, error) {
	query := `mutation UpdateAlert($id: String!, $input: UpdateAlertInput!){
	    updateAlert(id: $id, input: $input){
	        id
	        name
	        sex
	        height
	        weight
	        symptoms
	        comment
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Alert{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Alert{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Alert{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Alert{}, err
	}

	var result struct {
		Data struct {
			UpdateAlert model.Alert `json:"updateAlert"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Alert{}, err
	}

	return result.Data.UpdateAlert, nil
}

func CreateTreatment(input model.CreateTreatmentInput) (model.Treatment, error) {
	query := `mutation CreateTreatment($input: CreateTreatmentInput!){
	    createTreatment(input: $input){
	        id
	        period
	        day
	        quantity
	        medicine_id
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Treatment{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Treatment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Treatment{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Treatment{}, err
	}

	var result struct {
		Data struct {
			CreateTreatment model.Treatment `json:"createTreatment"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Treatment{}, err
	}

	return result.Data.CreateTreatment, nil
}

func UpdateAdmin(id string, input model.UpdateAdminInput) (model.Admin, error) {
	query := `mutation UpdateAdmin($id: String!, $input: UpdateAdminInput!){
	    updateAdmin(id: $id, input: $input){
	        id
	        email
	        password
	        name
	        last_name
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Admin{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Admin{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Admin{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Data struct {
			UpdateAdmin model.Admin `json:"updateAdmin"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	return result.Data.UpdateAdmin, nil
}

func UpdateSymptom(id string, input model.UpdateSymptomInput) (model.Symptom, error) {
	query := `mutation UpdateSymptom($id: String!, $input: UpdateSymptomInput!){
	    updateSymptom(id: $id, input: $input){
	        id
	        code
	        name
	        chronic
	        symptom
	        advice
	        question
	        question_basic
	        question_duration
	        question_ante
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Symptom{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Symptom{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Symptom{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Symptom{}, err
	}

	var result struct {
		Data struct {
			UpdateSymptom model.Symptom `json:"updateSymptom"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Symptom{}, err
	}

	return result.Data.UpdateSymptom, nil
}

func CreateTreatmentsFollowUp(input model.CreateTreatmentsFollowUpInput) (model.TreatmentsFollowUp, error) {
	query := `mutation CreateTreatmentsFollowUp($input: CreateTreatmentsFollowUpInput!){
	    createTreatmentsFollowUp(input: $input){
	        id
	        treatment_id
	        date
	        period
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.TreatmentsFollowUp{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	var result struct {
		Data struct {
			CreateTreatmentsFollowUp model.TreatmentsFollowUp `json:"createTreatmentsFollowUp"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	return result.Data.CreateTreatmentsFollowUp, nil
}

func UpdateNotification(id string, input model.UpdateNotificationInput) (model.Notification, error) {
	query := `mutation UpdateNotification($id: String!, $input: UpdateNotificationInput!){
	    updateNotification(id: $id, input: $input){
	        id
	        token
	        title
	        message
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Notification{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Notification{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Notification{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Notification{}, err
	}

	var result struct {
		Data struct {
			UpdateNotification model.Notification `json:"updateNotification"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Notification{}, err
	}

	return result.Data.UpdateNotification, nil
}

func CreateNotification(input model.CreateNotificationInput) (model.Notification, error) {
	query := `mutation CreateNotification($input: CreateNotificationInput!){
	    createNotification(input: $input){
	        id
	        token
	        title
	        message
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Notification{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Notification{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Notification{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Notification{}, err
	}

	var result struct {
		Data struct {
			CreateNotification model.Notification `json:"createNotification"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Notification{}, err
	}

	return result.Data.CreateNotification, nil
}

func UpdateAnteFamily(id string, input model.UpdateAnteFamilyInput) (model.AnteFamily, error) {
	query := `mutation UpdateAnteFamily($id: String!, $input: UpdateAnteFamilyInput!){
	    updateAnteFamily(id: $id, input: $input){
	        id
	        name
	        disease
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AnteFamily{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteFamily{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteFamily{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteFamily{}, err
	}

	var result struct {
		Data struct {
			UpdateAnteFamily model.AnteFamily `json:"updateAnteFamily"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteFamily{}, err
	}

	return result.Data.UpdateAnteFamily, nil
}

func UpdateAnteDisease(id string, input model.UpdateAnteDiseaseInput) (model.AnteDisease, error) {
	query := `mutation UpdateAnteDisease($id: String!, $input: UpdateAnteDiseaseInput!){
	    updateAnteDisease(id: $id, input: $input){
	        id
	        name
	        chronicity
	        surgery_ids
	        symptoms
	        treatment_ids
	        still_relevant
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AnteDisease{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteDisease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteDisease{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteDisease{}, err
	}

	var result struct {
		Data struct {
			UpdateAnteDisease model.AnteDisease `json:"updateAnteDisease"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteDisease{}, err
	}

	return result.Data.UpdateAnteDisease, nil
}

func CreateRdv(input model.CreateRdvInput) (model.Rdv, error) {
	query := `mutation CreateRdv($input: CreateRdvInput!){
	    createRdv(input: $input){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Rdv{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Rdv{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Data struct {
			CreateRdv model.Rdv `json:"createRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	return result.Data.CreateRdv, nil
}

func GetWaitingRdv(doctor_id string, option *model.Options) ([]model.Rdv, error) {
	query := `query GetWaitingRdv($doctor_id: String!, $option: Options){
	    getWaitingRdv(doctor_id: $doctor_id, option: $option){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"doctor_id": doctor_id,
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetWaitingRdv []model.Rdv `json:"getWaitingRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetWaitingRdv, nil
}

func GetTreatmentByID(id string) (model.Treatment, error) {
	query := `query GetTreatmentByID($id: String!){
	    getTreatmentByID(id: $id){
	        id
	        period
	        day
	        quantity
	        medicine_id
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Treatment{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Treatment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Treatment{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Treatment{}, err
	}

	var result struct {
		Data struct {
			GetTreatmentByID model.Treatment `json:"getTreatmentByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Treatment{}, err
	}

	return result.Data.GetTreatmentByID, nil
}

func GetAlerts(option *model.Options) ([]model.Alert, error) {
	query := `query GetAlerts($option: Options){
	    getAlerts(option: $option){
	        id
	        name
	        sex
	        height
	        weight
	        symptoms
	        comment
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetAlerts []model.Alert `json:"getAlerts"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetAlerts, nil
}

func GetMedicineByID(id string) (model.Medicine, error) {
	query := `query GetMedicineByID($id: String!){
	    getMedicineByID(id: $id){
	        id
	        name
	        unit
	        target_diseases
	        treated_symptoms
	        side_effects
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Medicine{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Medicine{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Medicine{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Medicine{}, err
	}

	var result struct {
		Data struct {
			GetMedicineByID model.Medicine `json:"getMedicineByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Medicine{}, err
	}

	return result.Data.GetMedicineByID, nil
}

func GetRdvById(id string) (model.Rdv, error) {
	query := `query GetRdvById($id: String!){
	    getRdvById(id: $id){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Rdv{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Rdv{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Data struct {
			GetRdvById model.Rdv `json:"getRdvById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	return result.Data.GetRdvById, nil
}

func GetPatientDocument(id string, option *model.Options) ([]model.Document, error) {
	query := `query GetPatientDocument($id: String!, $option: Options){
	    getPatientDocument(id: $id, option: $option){
	        id
	        owner_id
	        name
	        document_type
	        category
	        is_favorite
	        download_url
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetPatientDocument []model.Document `json:"getPatientDocument"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetPatientDocument, nil
}

func GetNotifications(option *model.Options) ([]model.Notification, error) {
	query := `query GetNotifications($option: Options){
	    getNotifications(option: $option){
	        id
	        token
	        title
	        message
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetNotifications []model.Notification `json:"getNotifications"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetNotifications, nil
}

func GetDoctorRdv(doctor_id string, option *model.Options) ([]model.Rdv, error) {
	query := `query GetDoctorRdv($doctor_id: String!, $option: Options){
	    getDoctorRdv(doctor_id: $doctor_id, option: $option){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"doctor_id": doctor_id,
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetDoctorRdv []model.Rdv `json:"getDoctorRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetDoctorRdv, nil
}

func GetTreatments(option *model.Options) ([]model.Treatment, error) {
	query := `query GetTreatments($option: Options){
	    getTreatments(option: $option){
	        id
	        period
	        day
	        quantity
	        medicine_id
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetTreatments []model.Treatment `json:"getTreatments"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetTreatments, nil
}

func GetNotificationById(id string) (model.Notification, error) {
	query := `query GetNotificationById($id: String!){
	    getNotificationById(id: $id){
	        id
	        token
	        title
	        message
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Notification{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Notification{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Notification{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Notification{}, err
	}

	var result struct {
		Data struct {
			GetNotificationById model.Notification `json:"getNotificationById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Notification{}, err
	}

	return result.Data.GetNotificationById, nil
}

func GetDocumentById(id string) (model.Document, error) {
	query := `query GetDocumentById($id: String!){
	    getDocumentById(id: $id){
	        id
	        owner_id
	        name
	        document_type
	        category
	        is_favorite
	        download_url
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Document{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Document{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Document{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Document{}, err
	}

	var result struct {
		Data struct {
			GetDocumentById model.Document `json:"getDocumentById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Document{}, err
	}

	return result.Data.GetDocumentById, nil
}

func GetAdminById(id string) (model.Admin, error) {
	query := `query GetAdminById($id: String!){
	    getAdminById(id: $id){
	        id
	        email
	        password
	        name
	        last_name
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Admin{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Admin{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Admin{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Data struct {
			GetAdminById model.Admin `json:"getAdminById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	return result.Data.GetAdminById, nil
}

func GetPatientByEmail(email string) (model.Patient, error) {
	query := `query GetPatientByEmail($email: String!){
	    getPatientByEmail(email: $email){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"email": email,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Patient{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Patient{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Patient{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Data struct {
			GetPatientByEmail model.Patient `json:"getPatientByEmail"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	return result.Data.GetPatientByEmail, nil
}

func GetAnteDiseaseByID(id string) (model.AnteDisease, error) {
	query := `query GetAnteDiseaseByID($id: String!){
	    getAnteDiseaseByID(id: $id){
	        id
	        name
	        chronicity
	        surgery_ids
	        symptoms
	        treatment_ids
	        still_relevant
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AnteDisease{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteDisease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteDisease{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteDisease{}, err
	}

	var result struct {
		Data struct {
			GetAnteDiseaseByID model.AnteDisease `json:"getAnteDiseaseByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteDisease{}, err
	}

	return result.Data.GetAnteDiseaseByID, nil
}

func GetAlertById(id string) (model.Alert, error) {
	query := `query GetAlertById($id: String!){
	    getAlertById(id: $id){
	        id
	        name
	        sex
	        height
	        weight
	        symptoms
	        comment
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Alert{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Alert{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Alert{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Alert{}, err
	}

	var result struct {
		Data struct {
			GetAlertById model.Alert `json:"getAlertById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Alert{}, err
	}

	return result.Data.GetAlertById, nil
}

func GetSymptoms(option *model.Options) ([]model.Symptom, error) {
	query := `query GetSymptoms($option: Options){
	    getSymptoms(option: $option){
	        id
	        code
	        name
	        chronic
	        symptom
	        advice
	        question
	        question_basic
	        question_duration
	        question_ante
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetSymptoms []model.Symptom `json:"getSymptoms"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetSymptoms, nil
}

func GetPatientsFromDoctorById(id string, option *model.Options) ([]model.Patient, error) {
	query := `query GetPatientsFromDoctorById($id: String!, $option: Options){
	    getPatientsFromDoctorById(id: $id, option: $option){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetPatientsFromDoctorById []model.Patient `json:"getPatientsFromDoctorById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetPatientsFromDoctorById, nil
}

func GetMedicalFolderById(id string) (model.MedicalInfo, error) {
	query := `query GetMedicalFolderById($id: String!){
	    getMedicalFolderById(id: $id){
	        id
	        name
	        firstname
	        birthdate
	        sex
	        height
	        weight
	        primary_doctor_id
	        onboarding_status
	        antecedent_disease_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.MedicalInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Data struct {
			GetMedicalFolderById model.MedicalInfo `json:"getMedicalFolderById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	return result.Data.GetMedicalFolderById, nil
}

func GetSymptomById(id string) (model.Symptom, error) {
	query := `query GetSymptomById($id: String!){
	    getSymptomById(id: $id){
	        id
	        code
	        name
	        chronic
	        symptom
	        advice
	        question
	        question_basic
	        question_duration
	        question_ante
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Symptom{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Symptom{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Symptom{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Symptom{}, err
	}

	var result struct {
		Data struct {
			GetSymptomById model.Symptom `json:"getSymptomById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Symptom{}, err
	}

	return result.Data.GetSymptomById, nil
}

func GetPatientById(id string) (model.Patient, error) {
	query := `query GetPatientById($id: String!){
	    getPatientById(id: $id){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Patient{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Patient{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Patient{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Data struct {
			GetPatientById model.Patient `json:"getPatientById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	return result.Data.GetPatientById, nil
}

func GetMedicalFolder(option *model.Options) ([]model.MedicalInfo, error) {
	query := `query GetMedicalFolder($option: Options){
	    getMedicalFolder(option: $option){
	        id
	        name
	        firstname
	        birthdate
	        sex
	        height
	        weight
	        primary_doctor_id
	        onboarding_status
	        antecedent_disease_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetMedicalFolder []model.MedicalInfo `json:"getMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetMedicalFolder, nil
}

func GetSlots(id string, option *model.Options) ([]model.Rdv, error) {
	query := `query GetSlots($id: String!, $option: Options){
	    getSlots(id: $id, option: $option){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetSlots []model.Rdv `json:"getSlots"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetSlots, nil
}

func GetPatientRdv(id_patient string, option *model.Options) ([]model.Rdv, error) {
	query := `query GetPatientRdv($id_patient: String!, $option: Options){
	    getPatientRdv(id_patient: $id_patient, option: $option){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id_patient": id_patient,
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetPatientRdv []model.Rdv `json:"getPatientRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetPatientRdv, nil
}

func GetAdminByEmail(email string) (model.Admin, error) {
	query := `query GetAdminByEmail($email: String!){
	    getAdminByEmail(email: $email){
	        id
	        email
	        password
	        name
	        last_name
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"email": email,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Admin{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Admin{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Admin{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Data struct {
			GetAdminByEmail model.Admin `json:"getAdminByEmail"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	return result.Data.GetAdminByEmail, nil
}

func GetSlotById(id string) (model.Rdv, error) {
	query := `query GetSlotById($id: String!){
	    getSlotById(id: $id){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
	        health_method
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Rdv{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Rdv{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Data struct {
			GetSlotById model.Rdv `json:"getSlotById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	return result.Data.GetSlotById, nil
}

func GetAdmins(option *model.Options) ([]model.Admin, error) {
	query := `query GetAdmins($option: Options){
	    getAdmins(option: $option){
	        id
	        email
	        password
	        name
	        last_name
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetAdmins []model.Admin `json:"getAdmins"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetAdmins, nil
}

func GetAnteFamilyByID(id string) (model.AnteFamily, error) {
	query := `query GetAnteFamilyByID($id: String!){
	    getAnteFamilyByID(id: $id){
	        id
	        name
	        disease
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AnteFamily{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteFamily{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteFamily{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteFamily{}, err
	}

	var result struct {
		Data struct {
			GetAnteFamilyByID model.AnteFamily `json:"getAnteFamilyByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteFamily{}, err
	}

	return result.Data.GetAnteFamilyByID, nil
}

func GetAnteDiseases(option *model.Options) ([]model.AnteDisease, error) {
	query := `query GetAnteDiseases($option: Options){
	    getAnteDiseases(option: $option){
	        id
	        name
	        chronicity
	        surgery_ids
	        symptoms
	        treatment_ids
	        still_relevant
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetAnteDiseases []model.AnteDisease `json:"getAnteDiseases"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetAnteDiseases, nil
}

func GetMedicines(option *model.Options) ([]model.Medicine, error) {
	query := `query GetMedicines($option: Options){
	    getMedicines(option: $option){
	        id
	        name
	        unit
	        target_diseases
	        treated_symptoms
	        side_effects
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetMedicines []model.Medicine `json:"getMedicines"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetMedicines, nil
}

func GetPatients(option *model.Options) ([]model.Patient, error) {
	query := `query GetPatients($option: Options){
	    getPatients(option: $option){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetPatients []model.Patient `json:"getPatients"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetPatients, nil
}

func GetTreatmentsFollowUpById(id string) (model.TreatmentsFollowUp, error) {
	query := `query GetTreatmentsFollowUpById($id: String!){
	    getTreatmentsFollowUpById(id: $id){
	        id
	        treatment_id
	        date
	        period
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.TreatmentsFollowUp{}, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	var result struct {
		Data struct {
			GetTreatmentsFollowUpById model.TreatmentsFollowUp `json:"getTreatmentsFollowUpById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	return result.Data.GetTreatmentsFollowUpById, nil
}

func GetDocuments(option *model.Options) ([]model.Document, error) {
	query := `query GetDocuments($option: Options){
	    getDocuments(option: $option){
	        id
	        owner_id
	        name
	        document_type
	        category
	        is_favorite
	        download_url
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetDocuments []model.Document `json:"getDocuments"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetDocuments, nil
}

func GetTreatmentsFollowUps(id string, option *model.Options) ([]model.TreatmentsFollowUp, error) {
	query := `query GetTreatmentsFollowUps($id: String!, $option: Options){
	    getTreatmentsFollowUps(id: $id, option: $option){
	        id
	        treatment_id
	        date
	        period
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetTreatmentsFollowUps []model.TreatmentsFollowUp `json:"getTreatmentsFollowUps"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetTreatmentsFollowUps, nil
}

func GetAnteFamilies(option *model.Options) ([]model.AnteFamily, error) {
	query := `query GetAnteFamilies($option: Options){
	    getAnteFamilies(option: $option){
	        id
	        name
	        disease
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"option": option,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %!v(MISSING)", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			GetAnteFamilies []model.AnteFamily `json:"getAnteFamilies"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.GetAnteFamilies, nil
}

