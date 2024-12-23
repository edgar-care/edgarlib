package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

func DeleteAntecdentTreatment(id string) (bool, error) {
	query := `mutation DeleteAntecdentTreatment($id: String!){
	    deleteAntecdentTreatment(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteAntecdentTreatment bool `json:"deleteAntecdentTreatment"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteAntecdentTreatment, nil
}

func CreateAnteChir(input model.CreateAnteChirInput) (model.AnteChir, error) {
	query := `mutation CreateAnteChir($input: CreateAnteChirInput!){
	    createAnteChir(input: $input){
	        id
	        name
	        induced_symptoms{
	            symptom
	            factor
	        }
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
		return model.AnteChir{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteChir{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteChir{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteChir{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateAnteChir model.AnteChir `json:"createAnteChir"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteChir{}, err
	}

	if len(result.Errors) > 0 {
		return model.AnteChir{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateAnteChir, nil
}

func CreateSymptom(input model.CreateSymptomInput) (model.Symptom, error) {
	query := `mutation CreateSymptom($input: CreateSymptomInput!){
	    createSymptom(input: $input){
	        id
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
		return model.Symptom{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Symptom{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateSymptom model.Symptom `json:"createSymptom"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Symptom{}, err
	}

	if len(result.Errors) > 0 {
		return model.Symptom{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateSymptom, nil
}

func DeleteDisease(id string) (bool, error) {
	query := `mutation DeleteDisease($id: String!){
	    deleteDisease(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteDisease bool `json:"deleteDisease"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteDisease, nil
}

func DeleteSymptom(id string) (bool, error) {
	query := `mutation DeleteSymptom($id: String!){
	    deleteSymptom(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteSymptom bool `json:"deleteSymptom"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteSymptom, nil
}

func DeleteAdmin(id string) (bool, error) {
	query := `mutation DeleteAdmin($id: String!){
	    deleteAdmin(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteAdmin bool `json:"deleteAdmin"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteAdmin, nil
}

func DeleteDoctor(id string) (bool, error) {
	query := `mutation DeleteDoctor($id: String!){
	    deleteDoctor(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteDoctor bool `json:"deleteDoctor"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteDoctor, nil
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
	        uploader_id
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
		return model.Document{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Document{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateDocument model.Document `json:"createDocument"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Document{}, err
	}

	if len(result.Errors) > 0 {
		return model.Document{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateDocument, nil
}

func CreateChat(input model.CreateChatInput) (model.Chat, error) {
	query := `mutation CreateChat($input: CreateChatInput!){
	    createChat(input: $input){
	        id
	        participants{
	            participant_id
	            last_seen
	        }
	        messages{
	            owner_id
	            message
	            sended_time
	        }
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
		return model.Chat{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Chat{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Chat{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Chat{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateChat model.Chat `json:"createChat"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Chat{}, err
	}

	if len(result.Errors) > 0 {
		return model.Chat{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateChat, nil
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
		return model.Admin{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateAdmin model.Admin `json:"createAdmin"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	if len(result.Errors) > 0 {
		return model.Admin{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateAdmin, nil
}

func DeletePatient(id string) (bool, error) {
	query := `mutation DeletePatient($id: String!){
	    deletePatient(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeletePatient bool `json:"deletePatient"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeletePatient, nil
}

func CreateNlpReport(input model.CreateNlpReportInput) (model.NlpReport, error) {
	query := `mutation CreateNlpReport($input: CreateNlpReportInput!){
	    createNlpReport(input: $input){
	        id
	        version
	        input_symptoms
	        input_sentence
	        output{
	            symptom
	            present
	            days
	        }
	        computation_time
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
		return model.NlpReport{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.NlpReport{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.NlpReport{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.NlpReport{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateNlpReport model.NlpReport `json:"createNlpReport"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.NlpReport{}, err
	}

	if len(result.Errors) > 0 {
		return model.NlpReport{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateNlpReport, nil
}

func UpdateAnteChir(id string, input model.UpdateAnteChirInput) (model.AnteChir, error) {
	query := `mutation UpdateAnteChir($id: String!, $input: UpdateAnteChirInput!){
	    updateAnteChir(id: $id, input: $input){
	        id
	        name
	        induced_symptoms{
	            symptom
	            factor
	        }
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
		return model.AnteChir{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteChir{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteChir{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteChir{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateAnteChir model.AnteChir `json:"updateAnteChir"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteChir{}, err
	}

	if len(result.Errors) > 0 {
		return model.AnteChir{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateAnteChir, nil
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
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateMedicalFolder model.MedicalInfo `json:"updateMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalInfo{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateMedicalFolder, nil
}

func UpdateMedicalFolderdAntedisease(id string, input model.UpdateMedicalFOlderAntedisease) (model.MedicalInfo, error) {
	query := `mutation UpdateMedicalFolderdAntedisease($id: String!, $input: UpdateMedicalFOlderAntedisease!){
	    updateMedicalFolderdAntedisease(id: $id, input: $input){
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
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateMedicalFolderdAntedisease model.MedicalInfo `json:"updateMedicalFolderdAntedisease"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalInfo{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateMedicalFolderdAntedisease, nil
}

func CreateMedicine(input model.CreateMedicineInput) (model.Medicine, error) {
	query := `mutation CreateMedicine($input: CreateMedicineInput!){
	    createMedicine(input: $input){
	        id
	        dci
	        target_diseases
	        treated_symptoms
	        side_effects
	        dosage
	        dosage_unit
	        container
	        name
	        dosage_form
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
		return model.Medicine{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Medicine{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateMedicine model.Medicine `json:"createMedicine"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Medicine{}, err
	}

	if len(result.Errors) > 0 {
		return model.Medicine{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateMedicine, nil
}

func DeleteAlert(id string) (bool, error) {
	query := `mutation DeleteAlert($id: String!){
	    deleteAlert(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteAlert bool `json:"deleteAlert"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteAlert, nil
}

func UpdateAntecedentTreatment(id string, antecedentID string, input model.UpdateAntecedentTreatmentInput) (model.AntecedentTreatment, error) {
	query := `mutation UpdateAntecedentTreatment($id: String!, $antecedentID: String!, $input: UpdateAntecedentTreatmentInput!){
	    updateAntecdentTreatment(id: $id, antecedentID:$antecedentID, input: $input){
	        id
	        created_by
	        start_date
	        end_date
	        medicines {
	            medicine_id
	            comment
	            period {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
	    }
	}`
	variables := map[string]interface{}{
		"id": id,
		"antecedentID": antecedentID,
		"input": input,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AntecedentTreatment{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AntecedentTreatment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AntecedentTreatment{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AntecedentTreatment{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateAntecedentTreatment model.AntecedentTreatment `json:"updateAntecdentTreatment"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AntecedentTreatment{}, err
	}

	if len(result.Errors) > 0 {
		return model.AntecedentTreatment{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateAntecedentTreatment, nil
}

func UpdateTreatmentsMedicalAntecedents(id string, input model.UpdateTreatmentMedicalAntecedentsInput) (model.MedicalAntecedents, error) {
	query := `mutation UpdateTreatmentsMedicalAntecedents($id: String!, $input: UpdateTreatmentMedicalAntecedentsInput!){
	    updateTreatmentsMedicalAntecedents(id: $id, input: $input){
	        id
	        name
	        symptoms
	        treatments {
	            id
	            created_by
	            start_date
	            end_date
	            medicines {
	                medicine_id
	                comment
	                period {
	                    quantity
	                    frequency
	                    frequency_ratio
	                    frequency_unit
	                    period_length
	                    period_unit
	                }
	            }
	        }
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
		return model.MedicalAntecedents{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.MedicalAntecedents{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.MedicalAntecedents{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateTreatmentsMedicalAntecedents model.MedicalAntecedents `json:"updateTreatmentsMedicalAntecedents"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalAntecedents{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateTreatmentsMedicalAntecedents, nil
}

func CreateDoctor(input model.CreateDoctorInput) (model.Doctor, error) {
	query := `mutation CreateDoctor($input: CreateDoctorInput!){
	    createDoctor(input: $input){
	        id
	        email
	        password
	        name
	        firstname
	        address {
	            street
	            zip_code
	            country
	        }
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
	        status
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
		return model.Doctor{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Doctor{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Doctor{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Doctor{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateDoctor model.Doctor `json:"createDoctor"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Doctor{}, err
	}

	if len(result.Errors) > 0 {
		return model.Doctor{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateDoctor, nil
}

func DeleteChat(id string) (bool, error) {
	query := `mutation DeleteChat($id: String!){
	    deleteChat(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteChat bool `json:"deleteChat"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteChat, nil
}

func UpdateChat(id string, input model.UpdateChatInput) (model.Chat, error) {
	query := `mutation UpdateChat($id: String!, $input: UpdateChatInput!){
	    updateChat(id: $id, input: $input){
	        id
	        participants{
	            participant_id
	            last_seen
	        }
	        messages{
	            owner_id
	            message
	            sended_time
	        }
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
		return model.Chat{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Chat{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Chat{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Chat{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateChat model.Chat `json:"updateChat"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Chat{}, err
	}

	if len(result.Errors) > 0 {
		return model.Chat{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateChat, nil
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
		return model.AnteFamily{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteFamily{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateAnteFamily model.AnteFamily `json:"createAnteFamily"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteFamily{}, err
	}

	if len(result.Errors) > 0 {
		return model.AnteFamily{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.TreatmentsFollowUp{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateTreatmentsFollowUp model.TreatmentsFollowUp `json:"updateTreatmentsFollowUp"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	if len(result.Errors) > 0 {
		return model.TreatmentsFollowUp{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
	        uploader_id
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
		return model.Document{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Document{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDocument model.Document `json:"updateDocument"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Document{}, err
	}

	if len(result.Errors) > 0 {
		return model.Document{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDocument, nil
}

func DeleteSlot(id string) (bool, error) {
	query := `mutation DeleteSlot($id: String!){
	    deleteSlot(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteSlot bool `json:"deleteSlot"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteSlot, nil
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
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateRdv model.Rdv `json:"updateRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	if len(result.Errors) > 0 {
		return model.Rdv{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateRdv, nil
}

func DeleteAnteFamily(id string) (bool, error) {
	query := `mutation DeleteAnteFamily($id: String!){
	    deleteAnteFamily(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteAnteFamily bool `json:"deleteAnteFamily"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteAnteFamily, nil
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
		return model.Alert{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Alert{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateAlert model.Alert `json:"createAlert"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Alert{}, err
	}

	if len(result.Errors) > 0 {
		return model.Alert{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateAlert, nil
}

func DeleteMedicalFolder(id string) (bool, error) {
	query := `mutation DeleteMedicalFolder($id: String!){
	    deleteMedicalFolder(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteMedicalFolder bool `json:"deleteMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteMedicalFolder, nil
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
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreatePatient model.Patient `json:"createPatient"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreatePatient, nil
}

func DeleteAnteChir(id string) (bool, error) {
	query := `mutation DeleteAnteChir($id: String!){
	    deleteAnteChir(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteAnteChir bool `json:"deleteAnteChir"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteAnteChir, nil
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
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdatePatient model.Patient `json:"updatePatient"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdatePatient, nil
}

func UpdatePatientsDeviceConnect(id string, input model.UpdatePatientsDeviceConnectInput) (model.Patient, error) {
	query := `mutation UpdatePatientsDeviceConnect($id: String!, $input: UpdatePatientsDeviceConnectInput!){
	    updatePatientsDeviceConnect(id: $id, input: $input){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdatePatientsDeviceConnect model.Patient `json:"updatePatientsDeviceConnect"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdatePatientsDeviceConnect, nil
}

func UpdatePatientTrustDevice(id string, input model.UpdatePatientTrustDeviceInput) (model.Patient, error) {
	query := `mutation UpdatePatientTrustDevice($id: String!, $input: UpdatePatientTrustDeviceInput!){
	    updatePatientTrustDevice(id: $id, input: $input){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdatePatientTrustDevice model.Patient `json:"updatePatientTrustDevice"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdatePatientTrustDevice, nil
}

func UpdatePatientsRendezVousIds(id string, input model.UpdatePatientRendezVousIdsInput) (model.Patient, error) {
	query := `mutation UpdatePatientsRendezVousIds($id: String!, $input: UpdatePatientRendezVousIdsInput!){
	    updatePatientsRendezVousIds(id: $id, input: $input){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdatePatientsRendezVousIds model.Patient `json:"updatePatientsRendezVousIds"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdatePatientsRendezVousIds, nil
}

func UpdatePatientFollowTreatment(id string, input model.UpdatePatientFollowTreatmentInput) (model.Patient, error) {
	query := `mutation UpdatePatientFollowTreatment($id: String!, $input: UpdatePatientFollowTreatmentInput!){
	    updatePatientFollowTreatment(id: $id, input: $input){
	        id
	        email
	        password
	        rendez_vous_ids
	        medical_info_id
	        document_ids
	        treatment_follow_up_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdatePatientFollowTreatment model.Patient `json:"updatePatientFollowTreatment"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdatePatientFollowTreatment, nil
}

func UpdateAccountsMedicalFolder(id string, input model.UpdateAccountMedicalFolder) (model.MedicalInfo, error) {
	query := `mutation UpdateAccountsMedicalFolder($id: String!, $input: UpdateAccountMedicalFolder!){
	    updateAccountsMedicalFolder(id: $id, input: $input){
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
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateAccountsMedicalFolder model.MedicalInfo `json:"updateAccountsMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalInfo{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateAccountsMedicalFolder, nil
}

func CreateSession(input model.CreateSessionInput) (model.Session, error) {
	query := `mutation CreateSession($input: CreateSessionInput!){
	    createSession(input: $input){
	        id
	        diseases{
	            name
	            presence
	        }
	        symptoms{
	            name
	            presence
	            duration
	            treated
	        }
	        age
	        height
	        weight
	        sex
	        medical_antecedents
	        medicine
	        last_question
	        logs{
	            question
	            answer
	        }
	        hereditary_disease
	        alerts
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
		return model.Session{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Session{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Session{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Session{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateSession model.Session `json:"createSession"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Session{}, err
	}

	if len(result.Errors) > 0 {
		return model.Session{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateSession, nil
}

func DeleteAnteDisease(id string) (bool, error) {
	query := `mutation DeleteAnteDisease($id: String!){
	    deleteAnteDisease(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteAnteDisease bool `json:"deleteAnteDisease"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteAnteDisease, nil
}

func UpdateDoctor(id string, input model.UpdateDoctorInput) (model.Doctor, error) {
	query := `mutation UpdateDoctor($id: String!, $input: UpdateDoctorInput!){
	    updateDoctor(id: $id, input: $input){
	        id
	        email
	        password
	        name
	        firstname
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        address {
	            street
	            zip_code
	            country
	        }
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
	        status
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
		return model.Doctor{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Doctor{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Doctor{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Doctor{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDoctor model.Doctor `json:"updateDoctor"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Doctor{}, err
	}

	if len(result.Errors) > 0 {
		return model.Doctor{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDoctor, nil
}

func UpdateDoctorsDeviceConnect(id string, input model.UpdateDoctorsDeviceConnectInput) (model.Doctor, error) {
	query := `mutation UpdateDoctorsDeviceConnect($id: String!, $input: UpdateDoctorsDeviceConnectInput!){
	    UpdateDoctorsDeviceConnect(id: $id, input: $input){
	        id
	        email
	        password
	        name
	        firstname
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        address {
	            street
	            zip_code
	            country
	        }
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
	        status
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
		return model.Doctor{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Doctor{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Doctor{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Doctor{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDoctorsDeviceConnect model.Doctor `json:"UpdateDoctorsDeviceConnect"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Doctor{}, err
	}

	if len(result.Errors) > 0 {
		return model.Doctor{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDoctorsDeviceConnect, nil
}

func UpdateDoctorsTrustDevice(id string, input model.UpdateDoctorsTrustDeviceInput) (model.Doctor, error) {
	query := `mutation UpdateDoctorsTrustDevice($id: String!, $input: UpdateDoctorsTrustDeviceInput!){
	    UpdateDoctorsTrustDevice(id: $id, input: $input){
	        id
	        email
	        password
	        name
	        firstname
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        address {
	            street
	            zip_code
	            country
	        }
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
	        status
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
		return model.Doctor{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Doctor{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Doctor{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Doctor{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDoctorsTrustDevice model.Doctor `json:"UpdateDoctorsTrustDevice"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Doctor{}, err
	}

	if len(result.Errors) > 0 {
		return model.Doctor{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDoctorsTrustDevice, nil
}

func UpdateDoctorsPatientIDs(id string, input model.UpdateDoctorsPatientIDsInput) (model.Doctor, error) {
	query := `mutation UpdateDoctorsPatientIDs($id: String!, $input: UpdateDoctorsPatientIDsInput!){
	    updateDoctorsPatientIDs(id: $id, input: $input){
	        id
	        email
	        password
	        name
	        firstname
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        address {
	            street
	            zip_code
	            country
	        }
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
	        status
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
		return model.Doctor{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Doctor{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Doctor{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Doctor{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDoctorsPatientIDs model.Doctor `json:"updateDoctorsPatientIDs"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Doctor{}, err
	}

	if len(result.Errors) > 0 {
		return model.Doctor{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDoctorsPatientIDs, nil
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
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateMedicalFolder model.MedicalInfo `json:"createMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalInfo{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateMedicalFolder, nil
}

func DeleteDocument(id string) (bool, error) {
	query := `mutation DeleteDocument($id: String!){
	    deleteDocument(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteDocument bool `json:"deleteDocument"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteDocument, nil
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
		return model.Alert{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Alert{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateAlert model.Alert `json:"updateAlert"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Alert{}, err
	}

	if len(result.Errors) > 0 {
		return model.Alert{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateAlert, nil
}

func DeleteMedicine(id string) (bool, error) {
	query := `mutation DeleteMedicine($id: String!){
	    deleteMedicine(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteMedicine bool `json:"deleteMedicine"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteMedicine, nil
}

func CreateAntecdentTreatment(id string, input model.CreateAntecedentTreatmentInput) (model.AntecedentTreatment, error) {
	query := `mutation CreateAntecdentTreatment($id: String!, $input: CreateAntecedentTreatmentInput!){
	    createAntecdentTreatment(id:$id, input: $input){
	        id
	        created_by
	        start_date
	        end_date
	        medicines {
	            medicine_id
	            comment
	            period {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
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
		return model.AntecedentTreatment{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AntecedentTreatment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AntecedentTreatment{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AntecedentTreatment{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateAntecdentTreatment model.AntecedentTreatment `json:"createAntecdentTreatment"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AntecedentTreatment{}, err
	}

	if len(result.Errors) > 0 {
		return model.AntecedentTreatment{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateAntecdentTreatment, nil
}

func DeleteRdv(id string) (bool, error) {
	query := `mutation DeleteRdv($id: String!){
	    deleteRdv(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteRdv bool `json:"deleteRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteRdv, nil
}

func UpdateDisease(id string, input model.UpdateDiseaseInput) (model.Disease, error) {
	query := `mutation UpdateDisease($id: String!, $input: UpdateDiseaseInput!){
	    updateDisease(id: $id, input: $input){
	        id
	        name
	        symptoms
	        symptoms_weight{
	            symptom
	            value
	            chronic
	        }
	        overweight_factor
	        heredity_factor
	        advice
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
		return model.Disease{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Disease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Disease{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Disease{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDisease model.Disease `json:"updateDisease"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Disease{}, err
	}

	if len(result.Errors) > 0 {
		return model.Disease{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDisease, nil
}

func DeleteNotification(id string) (bool, error) {
	query := `mutation DeleteNotification($id: String!){
	    deleteNotification(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteNotification bool `json:"deleteNotification"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteNotification, nil
}

func DeleteSession(id string) (bool, error) {
	query := `mutation DeleteSession($id: String!){
	    deleteSession(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteSession bool `json:"deleteSession"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteSession, nil
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
		return model.Admin{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateAdmin model.Admin `json:"updateAdmin"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	if len(result.Errors) > 0 {
		return model.Admin{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateAdmin, nil
}

func UpdateSymptom(id string, input model.UpdateSymptomInput) (model.Symptom, error) {
	query := `mutation UpdateSymptom($id: String!, $input: UpdateSymptomInput!){
	    updateSymptom(id: $id, input: $input){
	        id
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
		return model.Symptom{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Symptom{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateSymptom model.Symptom `json:"updateSymptom"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Symptom{}, err
	}

	if len(result.Errors) > 0 {
		return model.Symptom{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.TreatmentsFollowUp{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateTreatmentsFollowUp model.TreatmentsFollowUp `json:"createTreatmentsFollowUp"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	if len(result.Errors) > 0 {
		return model.TreatmentsFollowUp{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateTreatmentsFollowUp, nil
}

func CreateDisease(input model.CreateDiseaseInput) (model.Disease, error) {
	query := `mutation CreateDisease($input: CreateDiseaseInput!){
	    createDisease(input: $input){
	        id
	        name
	        symptoms
	        symptoms_weight{
	            symptom
	            value
	            chronic
	        }
	        overweight_factor
	        heredity_factor
	        advice
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
		return model.Disease{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Disease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Disease{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Disease{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateDisease model.Disease `json:"createDisease"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Disease{}, err
	}

	if len(result.Errors) > 0 {
		return model.Disease{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateDisease, nil
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
		return model.Notification{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Notification{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateNotification model.Notification `json:"updateNotification"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Notification{}, err
	}

	if len(result.Errors) > 0 {
		return model.Notification{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.Notification{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Notification{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateNotification model.Notification `json:"createNotification"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Notification{}, err
	}

	if len(result.Errors) > 0 {
		return model.Notification{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.AnteFamily{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteFamily{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateAnteFamily model.AnteFamily `json:"updateAnteFamily"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteFamily{}, err
	}

	if len(result.Errors) > 0 {
		return model.AnteFamily{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateAnteFamily, nil
}

func DeleteTreatmentsFollowUp(id string) (bool, error) {
	query := `mutation DeleteTreatmentsFollowUp($id: String!){
	    deleteTreatmentsFollowUp(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteTreatmentsFollowUp bool `json:"deleteTreatmentsFollowUp"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteTreatmentsFollowUp, nil
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
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateRdv model.Rdv `json:"createRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	if len(result.Errors) > 0 {
		return model.Rdv{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateRdv, nil
}

func UpdateSession(id string, input model.UpdateSessionInput) (model.Session, error) {
	query := `mutation UpdateSession($id: String!, $input: UpdateSessionInput!){
	    updateSession(id: $id, input: $input){
	        id
	        diseases{
	            name
	            presence
	        }
	        symptoms{
	            name
	            presence
	            duration
	            treated
	        }
	        age
	        height
	        weight
	        sex
	        medical_antecedents
	        medicine
	        last_question
	        logs{
	            question
	            answer
	        }
	        hereditary_disease
	        alerts
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
		return model.Session{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Session{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Session{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Session{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateSession model.Session `json:"updateSession"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Session{}, err
	}

	if len(result.Errors) > 0 {
		return model.Session{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateSession, nil
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetWaitingRdv []model.Rdv `json:"getWaitingRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetWaitingRdv, nil
}

func GetAntecedentTreatmentByID(id string) (model.AntecedentTreatment, error) {
	query := `query GetAntecedentTreatmentByID($id: String!){
	    getAntecedentTreatmentByID(id: $id){
	        id
	        created_by
	        start_date
	        end_date
	        medicines {
	            medicine_id
	            comment
	            period {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
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
		return model.AntecedentTreatment{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AntecedentTreatment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AntecedentTreatment{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AntecedentTreatment{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAntecedentTreatmentByID model.AntecedentTreatment `json:"getAntecedentTreatmentByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AntecedentTreatment{}, err
	}

	if len(result.Errors) > 0 {
		return model.AntecedentTreatment{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAntecedentTreatmentByID, nil
}

func GetDiseases(option *model.Options) ([]model.Disease, error) {
	query := `query GetDiseases($option: Options){
	    getDiseases(option: $option){
	        id
	        name
	        symptoms
	        symptoms_weight{
	            symptom
	            value
	            chronic
	        }
	        overweight_factor
	        heredity_factor
	        advice
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDiseases []model.Disease `json:"getDiseases"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDiseases, nil
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAlerts []model.Alert `json:"getAlerts"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAlerts, nil
}

func GetMedicineByID(id string) (model.Medicine, error) {
	query := `query GetMedicineByID($id: String!){
	    getMedicineByID(id: $id){
	        id
	        dci
	        target_diseases
	        treated_symptoms
	        side_effects
	        dosage
	        dosage_unit
	        container
	        name
	        dosage_form
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
		return model.Medicine{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Medicine{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetMedicineByID model.Medicine `json:"getMedicineByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Medicine{}, err
	}

	if len(result.Errors) > 0 {
		return model.Medicine{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetMedicineByID, nil
}

func GetMedicineByIDWithSymptoms(medicineId string) (model.Medicine, error) {
	query := `query GetMedicineByIDWithSymptoms($medicineId: String!){
	    getMedicineByIDWithSymptoms(medicineId: $medicineId){
	        id
	        dci
	        target_diseases
	        treated_symptoms
	        side_effects
	        dosage
	        dosage_unit
	        container
	        name
	        dosage_form
	        createdAt
	        updatedAt
	        symptoms {
	            id
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
	    }
	}`
	variables := map[string]interface{}{
		"medicineId": medicineId,
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
		return model.Medicine{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Medicine{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetMedicineByIDWithSymptoms model.Medicine `json:"getMedicineByIDWithSymptoms"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Medicine{}, err
	}

	if len(result.Errors) > 0 {
		return model.Medicine{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetMedicineByIDWithSymptoms, nil
}

func GetAnteChirs(option *model.Options) ([]model.AnteChir, error) {
	query := `query GetAnteChirs($option: Options){
	    getAnteChirs(option: $option){
	        id
	        name
	        induced_symptoms{
	            symptom
	            factor
	        }
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAnteChirs []model.AnteChir `json:"getAnteChirs"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAnteChirs, nil
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
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetRdvById model.Rdv `json:"getRdvById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	if len(result.Errors) > 0 {
		return model.Rdv{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetPatientDocument []model.Document `json:"getPatientDocument"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetNotifications []model.Notification `json:"getNotifications"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDoctorRdv []model.Rdv `json:"getDoctorRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDoctorRdv, nil
}

func GetAntecedentTreatments(option *model.Options) ([]model.AntecedentTreatment, error) {
	query := `query GetAntecedentTreatments($option: Options){
	    getAntecedentTreatments(option: $option){
	        id
	        created_by
	        start_date
	        end_date
	        medicines {
	            medicine_id
	            comment
	            period {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAntecedentTreatments []model.AntecedentTreatment `json:"getAntecedentTreatments"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAntecedentTreatments, nil
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
		return model.Notification{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Notification{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetNotificationById model.Notification `json:"getNotificationById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Notification{}, err
	}

	if len(result.Errors) > 0 {
		return model.Notification{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetNotificationById, nil
}

func GetDiseaseById(id string) (model.Disease, error) {
	query := `query GetDiseaseById($id: String!){
	    getDiseaseById(id: $id){
	        id
	        name
	        symptoms
	        symptoms_weight{
	            symptom
	            value
	            chronic
	        }
	        overweight_factor
	        heredity_factor
	        advice
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
		return model.Disease{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Disease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Disease{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Disease{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDiseaseById model.Disease `json:"getDiseaseById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Disease{}, err
	}

	if len(result.Errors) > 0 {
		return model.Disease{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDiseaseById, nil
}

func GetSymptomsByDiseaseName(name string) (model.Disease, error) {
	query := `query GetSymptomsByDiseaseName($name: String!){
	    getSymptomsByDiseaseName(name: $name) {
	        symptoms
	    }
	}`
	variables := map[string]interface{}{
		"name": name,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.Disease{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Disease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Disease{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Disease{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSymptomsByDiseaseName model.Disease `json:"getSymptomsByDiseaseName"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Disease{}, err
	}

	if len(result.Errors) > 0 {
		return model.Disease{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSymptomsByDiseaseName, nil
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
	        uploader_id
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
		return model.Document{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Document{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDocumentById model.Document `json:"getDocumentById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Document{}, err
	}

	if len(result.Errors) > 0 {
		return model.Document{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.Admin{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAdminById model.Admin `json:"getAdminById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	if len(result.Errors) > 0 {
		return model.Admin{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetPatientByEmail model.Patient `json:"getPatientByEmail"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetPatientByEmail, nil
}

func GetChatById(id string) (model.Chat, error) {
	query := `query GetChatById($id: String!){
	    getChatById(id: $id){
	        id
	        participants{
	            participant_id
	            last_seen
	        }
	        messages{
	            owner_id
	            message
	            sended_time
	        }
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
		return model.Chat{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Chat{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Chat{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Chat{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetChatById model.Chat `json:"getChatById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Chat{}, err
	}

	if len(result.Errors) > 0 {
		return model.Chat{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetChatById, nil
}

func GetDoctorByEmail(email string) (model.Doctor, error) {
	query := `query GetDoctorByEmail($email: String!){
	    getDoctorByEmail(email: $email){
	        id
	        email
	        password
	        name
	        firstname
	        address{
	            street
	            zip_code
	            country
	            city
	        }
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
	        status
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
		return model.Doctor{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Doctor{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Doctor{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Doctor{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDoctorByEmail model.Doctor `json:"getDoctorByEmail"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Doctor{}, err
	}

	if len(result.Errors) > 0 {
		return model.Doctor{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDoctorByEmail, nil
}

func GetNlpReports(option *model.Options) ([]model.NlpReport, error) {
	query := `query GetNlpReports($option: Options){
	    getNlpReports(option: $option){
	        id
	        version
	        input_symptoms
	        input_sentence
	        output{
	            symptom
	            present
	            days
	        }
	        computation_time
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetNlpReports []model.NlpReport `json:"getNlpReports"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetNlpReports, nil
}

func GetNlpReportsByVersion(version int, option *model.Options) ([]model.NlpReport, error) {
	query := `query GetNlpReportsByVersion($version: Int!, $option: Options){
	    getNlpReportsByVersion(version: $version, option: $option){
	        id
	        version
	        input_symptoms
	        input_sentence
	        output{
	            symptom
	            present
	            days
	        }
	        computation_time
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"version": version,
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetNlpReportsByVersion []model.NlpReport `json:"getNlpReportsByVersion"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetNlpReportsByVersion, nil
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
		return model.Alert{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Alert{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAlertById model.Alert `json:"getAlertById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Alert{}, err
	}

	if len(result.Errors) > 0 {
		return model.Alert{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAlertById, nil
}

func GetAnteChirByID(id string) (model.AnteChir, error) {
	query := `query GetAnteChirByID($id: String!){
	    getAnteChirByID(id: $id){
	        id
	        name
	        induced_symptoms{
	            symptom
	            factor
	        }
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
		return model.AnteChir{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AnteChir{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AnteChir{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteChir{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAnteChirByID model.AnteChir `json:"getAnteChirByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteChir{}, err
	}

	if len(result.Errors) > 0 {
		return model.AnteChir{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAnteChirByID, nil
}

func GetSymptoms(option *model.Options) ([]model.Symptom, error) {
	query := `query GetSymptoms($option: Options){
	    getSymptoms(option: $option){
	        id
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSymptoms []model.Symptom `json:"getSymptoms"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSymptoms, nil
}

func GetSessionById(id string) (model.Session, error) {
	query := `query GetSessionById($id: String!){
	    getSessionById(id: $id){
	        id
	        diseases{
	            name
	            presence
	        }
	        symptoms{
	            name
	            presence
	            duration
	            treated
	        }
	        age
	        height
	        weight
	        sex
	        medical_antecedents
	        medicine
	        last_question
	        logs{
	            question
	            answer
	        }
	        hereditary_disease
	        alerts
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
		return model.Session{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Session{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Session{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Session{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSessionById model.Session `json:"getSessionById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Session{}, err
	}

	if len(result.Errors) > 0 {
		return model.Session{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSessionById, nil
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
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetPatientsFromDoctorById []model.Patient `json:"getPatientsFromDoctorById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetPatientsFromDoctorById, nil
}

func GetMedicalFolderByID(id string) (model.MedicalInfo, error) {
	query := `query GetMedicalFolderByID($id: String!){
	    getMedicalFolderById(id: $id){
	        id
	        name
	        firstname
	        birthdate
	        sex
	        height
	        weight
	        primary_doctor_id
	        antecedent_disease_ids
	        onboarding_status
	        family_members_med_info_id
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
		return model.MedicalInfo{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetMedicalFolderByID model.MedicalInfo `json:"getMedicalFolderById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalInfo{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalInfo{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetMedicalFolderByID, nil
}

func GetSymptomById(id string) (model.Symptom, error) {
	query := `query GetSymptomById($id: String!){
	    getSymptomById(id: $id){
	        id
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
		return model.Symptom{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Symptom{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSymptomById model.Symptom `json:"getSymptomById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Symptom{}, err
	}

	if len(result.Errors) > 0 {
		return model.Symptom{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSymptomById, nil
}

func GetSymptomByName(name string) (model.Symptom, error) {
	query := `query GetSymptomByName($name: String!){
	    getSymptomByName(name: $name) {
	        id
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
		"name": name,
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
		return model.Symptom{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Symptom{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSymptomByName model.Symptom `json:"getSymptomByName"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Symptom{}, err
	}

	if len(result.Errors) > 0 {
		return model.Symptom{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSymptomByName, nil
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
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return model.Patient{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Patient{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetPatientById model.Patient `json:"getPatientById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Patient{}, err
	}

	if len(result.Errors) > 0 {
		return model.Patient{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
	        family_members_med_info_id
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetMedicalFolder []model.MedicalInfo `json:"getMedicalFolder"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetMedicalFolder, nil
}

func GetRdvs(option *model.Options) ([]model.Rdv, error) {
	query := `query GetRdvs($option: Options){
	    getRdvs(option: $option){
	        id
	        doctor_id
	        id_patient
	        start_date
	        end_date
	        cancelation_reason
	        appointment_status
	        session_id
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetRdvs []model.Rdv `json:"getRdvs"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetRdvs, nil
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSlots []model.Rdv `json:"getSlots"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetPatientRdv []model.Rdv `json:"getPatientRdv"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.Admin{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Admin{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAdminByEmail model.Admin `json:"getAdminByEmail"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Admin{}, err
	}

	if len(result.Errors) > 0 {
		return model.Admin{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.Rdv{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Rdv{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSlotById model.Rdv `json:"getSlotById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Rdv{}, err
	}

	if len(result.Errors) > 0 {
		return model.Rdv{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSlotById, nil
}

func GetDoctors(option *model.Options) ([]model.Doctor, error) {
	query := `query GetDoctors($option: Options){
	    getDoctors(option: $option){
	        id
	        email
	        password
	        name
	        firstname
	        address{
	            street
	            zip_code
	            country
	            city
	        }
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDoctors []model.Doctor `json:"getDoctors"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDoctors, nil
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAdmins []model.Admin `json:"getAdmins"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.AnteFamily{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AnteFamily{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAnteFamilyByID model.AnteFamily `json:"getAnteFamilyByID"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AnteFamily{}, err
	}

	if len(result.Errors) > 0 {
		return model.AnteFamily{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAnteFamilyByID, nil
}

func GetChats(id string, option *model.Options) ([]model.Chat, error) {
	query := `query GetChats($id: String!, $option: Options){
	    getChats(id: $id, option: $option){
	        id
	        participants{
	            participant_id
	            last_seen
	        }
	        messages{
	            owner_id
	            message
	            sended_time
	        }
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetChats []model.Chat `json:"getChats"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetChats, nil
}

func GetSessions(option *model.Options) ([]model.Session, error) {
	query := `query GetSessions($option: Options){
	    getSessions(option: $option){
	        id
	        diseases{
	            name
	            presence
	        }
	        symptoms{
	            name
	            presence
	            duration
	            treated
	        }
	        age
	        height
	        weight
	        sex
	        medical_antecedents
	        medicine
	        last_question
	        logs{
	            question
	            answer
	        }
	        hereditary_disease
	        alerts
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSessions []model.Session `json:"getSessions"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSessions, nil
}

func GetMedicines(option *model.Options) ([]model.Medicine, error) {
	query := `query GetMedicines($option: Options){
	    getMedicines(option: $option){
	        id
	        dci
	        target_diseases
	        treated_symptoms
	        side_effects
	        dosage
	        dosage_unit
	        container
	        name
	        dosage_form
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetMedicines []model.Medicine `json:"getMedicines"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetMedicines, nil
}

func GetDoctorById(id string) (model.Doctor, error) {
	query := `query GetDoctorById($id: String!){
	    getDoctorById(id: $id){
	        id
	        email
	        password
	        name
	        firstname
	        address{
	            street
	            zip_code
	            country
	            city
	        }
	        rendez_vous_ids
	        patient_ids
	        chat_ids
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        ordonnance_ids
	        status
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
		return model.Doctor{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Doctor{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Doctor{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Doctor{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDoctorById model.Doctor `json:"getDoctorById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Doctor{}, err
	}

	if len(result.Errors) > 0 {
		return model.Doctor{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDoctorById, nil
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
	        device_connect
	        double_auth_methods_id
	        trust_devices
	        status
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetPatients []model.Patient `json:"getPatients"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return model.TreatmentsFollowUp{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetTreatmentsFollowUpById model.TreatmentsFollowUp `json:"getTreatmentsFollowUpById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.TreatmentsFollowUp{}, err
	}

	if len(result.Errors) > 0 {
		return model.TreatmentsFollowUp{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
	        uploader_id
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDocuments []model.Document `json:"getDocuments"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetTreatmentsFollowUps []model.TreatmentsFollowUp `json:"getTreatmentsFollowUps"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAnteFamilies []model.AnteFamily `json:"getAnteFamilies"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAnteFamilies, nil
}

func CreateDeviceConnect(input model.CreateDeviceConnectInput) (model.DeviceConnect, error) {
	query := `mutation CreateDeviceConnect($input: CreateDeviceConnectInput!) {
	    createDeviceConnect(input: $input) {
	        id
	        device_type
	        browser
	        ip_address
	        city
	        country
	        date
	        trust_device
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
		return model.DeviceConnect{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.DeviceConnect{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.DeviceConnect{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateDeviceConnect model.DeviceConnect `json:"createDeviceConnect"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	if len(result.Errors) > 0 {
		return model.DeviceConnect{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateDeviceConnect, nil
}

func UpdateDeviceConnect(id string, input model.UpdateDeviceConnectInput) (model.DeviceConnect, error) {
	query := `mutation UpdateDeviceConnect($id: String!, $input: UpdateDeviceConnectInput!) {
	    updateDeviceConnect(id: $id, input: $input) {
	        id
	        device_type
	        browser
	        ip_address
	        city
	        country
	        date
	        trust_device
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
		return model.DeviceConnect{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.DeviceConnect{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.DeviceConnect{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDeviceConnect model.DeviceConnect `json:"updateDeviceConnect"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	if len(result.Errors) > 0 {
		return model.DeviceConnect{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDeviceConnect, nil
}

func DeleteDeviceConnect(id string) (bool, error) {
	query := `mutation DeleteDeviceConnect($id: String!){
	    deleteDeviceConnect(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteDeviceConnect bool `json:"deleteDeviceConnect"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteDeviceConnect, nil
}

func GetDevicesConnect(id string, option *model.Options) ([]model.DeviceConnect, error) {
	query := `query GetDevicesConnect($id: String!, $option: Options){
	    getDevicesConnect(id:$id, option: $option){
	        id
	        device_type
	        browser
	        ip_address
	        city
	        country
	        date
	        trust_device
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDevicesConnect []model.DeviceConnect `json:"getDevicesConnect"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDevicesConnect, nil
}

func GetDeviceConnectById(id string) (model.DeviceConnect, error) {
	query := `query GetDeviceConnectById($id: String!){
	    getDeviceConnectById(id:$id){
	        id
	        device_type
	        browser
	        ip_address
	        city
	        country
	        date
	        trust_device
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
		return model.DeviceConnect{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.DeviceConnect{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.DeviceConnect{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDeviceConnectById model.DeviceConnect `json:"getDeviceConnectById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	if len(result.Errors) > 0 {
		return model.DeviceConnect{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDeviceConnectById, nil
}

func GetDeviceConnectByIp(ip_address string) (model.DeviceConnect, error) {
	query := `query GetDeviceConnectByIp($ip_address: String!){
	    getDeviceConnectByIp(ip_address:$ip_address){
	        id
	        device_type
	        browser
	        ip_address
	        city
	        country
	        date
	        trust_device
	    }
	}`
	variables := map[string]interface{}{
		"ip_address": ip_address,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.DeviceConnect{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.DeviceConnect{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDeviceConnectByIp model.DeviceConnect `json:"getDeviceConnectByIp"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.DeviceConnect{}, err
	}

	if len(result.Errors) > 0 {
		return model.DeviceConnect{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDeviceConnectByIp, nil
}

func CreateDoubleAuth(input model.CreateDoubleAuthInput) (model.DoubleAuth, error) {
	query := `mutation CreateDoubleAuth($input: CreateDoubleAuthInput!) {
	    createDoubleAuth(input: $input) {
	        id
	        methods
	        secret
	        code
	        trust_device_id
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
		return model.DoubleAuth{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.DoubleAuth{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.DoubleAuth{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.DoubleAuth{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateDoubleAuth model.DoubleAuth `json:"createDoubleAuth"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.DoubleAuth{}, err
	}

	if len(result.Errors) > 0 {
		return model.DoubleAuth{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateDoubleAuth, nil
}

func UpdateDoubleAuth(id string, input model.UpdateDoubleAuthInput) (model.DoubleAuth, error) {
	query := `mutation UpdateDoubleAuth($id: String!, $input: UpdateDoubleAuthInput!) {
	    updateDoubleAuth(id: $id, input: $input) {
	        id
	        methods
	        secret
	        code
	        trust_device_id
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
		return model.DoubleAuth{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.DoubleAuth{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.DoubleAuth{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.DoubleAuth{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateDoubleAuth model.DoubleAuth `json:"updateDoubleAuth"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.DoubleAuth{}, err
	}

	if len(result.Errors) > 0 {
		return model.DoubleAuth{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateDoubleAuth, nil
}

func DeleteDoubleAuth(id string) (bool, error) {
	query := `mutation DeleteDoubleAuth($id: String!){
	    deleteDoubleAuth(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteDoubleAuth bool `json:"deleteDoubleAuth"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteDoubleAuth, nil
}

func GetDoubleAuths(option *model.Options) ([]model.DoubleAuth, error) {
	query := `query GetDoubleAuths($option: Options){
	    getDoubleAuths(option: $option){
	        id
	        methods
	        secret
	        code
	        trust_device_id
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDoubleAuths []model.DoubleAuth `json:"getDoubleAuths"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDoubleAuths, nil
}

func GetDoubleAuthById(id string) (model.DoubleAuth, error) {
	query := `query GetDoubleAuthById($id: String!){
	    getDoubleAuthById(id: $id){
	        id
	        methods
	        secret
	        code
	        trust_device_id
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
		return model.DoubleAuth{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.DoubleAuth{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.DoubleAuth{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.DoubleAuth{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetDoubleAuthById model.DoubleAuth `json:"getDoubleAuthById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.DoubleAuth{}, err
	}

	if len(result.Errors) > 0 {
		return model.DoubleAuth{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetDoubleAuthById, nil
}

func CreateBlackList(input model.CreateBlackListInput) (model.BlackList, error) {
	query := `mutation CreateBlackList($input: CreateBlackListInput!) {
	    createBlackList(input: $input){
	        id
	        token
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
		return model.BlackList{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.BlackList{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.BlackList{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.BlackList{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateBlackList model.BlackList `json:"createBlackList"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.BlackList{}, err
	}

	if len(result.Errors) > 0 {
		return model.BlackList{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateBlackList, nil
}

func UpdateBlackList(id string, input model.UpdateBlackListInput) (model.BlackList, error) {
	query := `mutation UpdateBlackList($id: String!, $input: UpdateBlackListInput!) {
	    updateBlackList(id: $id, input: $input){
	        id
	        token
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
		return model.BlackList{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.BlackList{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.BlackList{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.BlackList{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateBlackList model.BlackList `json:"updateBlackList"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.BlackList{}, err
	}

	if len(result.Errors) > 0 {
		return model.BlackList{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateBlackList, nil
}

func DeleteBlackList(id string) (bool, error) {
	query := `mutation DeleteBlackList($id: String!){
	    deleteDoubleAuth(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteBlackList bool `json:"deleteDoubleAuth"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteBlackList, nil
}

func GetBlackListById(id string) (model.BlackList, error) {
	query := `query GetBlackListById($id: String!){
	    getBlackListById(id: $id){
	        id
	        token
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
		return model.BlackList{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.BlackList{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.BlackList{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.BlackList{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetBlackListById model.BlackList `json:"getBlackListById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.BlackList{}, err
	}

	if len(result.Errors) > 0 {
		return model.BlackList{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetBlackListById, nil
}

func GetBlackList(option *model.Options) ([]model.BlackList, error) {
	query := `query GetBlackList($option: Options){
	    getBlackList(option: $option){
	        id
	        token
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetBlackList []model.BlackList `json:"getBlackList"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetBlackList, nil
}

func CreateSaveCode(input model.CreateSaveCodeInput) (model.SaveCode, error) {
	query := `mutation CreateSaveCode($input: CreateSaveCodeInput!){
	    createSaveCode(input: $input){
	        id
	        code
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
		return model.SaveCode{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.SaveCode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.SaveCode{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.SaveCode{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateSaveCode model.SaveCode `json:"createSaveCode"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.SaveCode{}, err
	}

	if len(result.Errors) > 0 {
		return model.SaveCode{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateSaveCode, nil
}

func UpdateSaveCode(id string, input model.UpdateSaveCodeInput) (model.SaveCode, error) {
	query := `mutation UpdateSaveCode($id: String!, $input: UpdateSaveCodeInput!){
	    updateSaveCode(id: $id, input: $input){
	        id
	        code
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
		return model.SaveCode{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.SaveCode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.SaveCode{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.SaveCode{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateSaveCode model.SaveCode `json:"updateSaveCode"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.SaveCode{}, err
	}

	if len(result.Errors) > 0 {
		return model.SaveCode{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateSaveCode, nil
}

func DeleteSaveCode(id string) (bool, error) {
	query := `mutation DeleteSaveCode($id: String!){
	    deleteSaveCode(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteSaveCode bool `json:"deleteSaveCode"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteSaveCode, nil
}

func GetSaveCodeById(id string) (model.SaveCode, error) {
	query := `query GetSaveCodeById($id: String!){
	    getSaveCodeById(id: $id){
	        id
	        code
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
		return model.SaveCode{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.SaveCode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.SaveCode{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.SaveCode{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSaveCodeById model.SaveCode `json:"getSaveCodeById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.SaveCode{}, err
	}

	if len(result.Errors) > 0 {
		return model.SaveCode{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSaveCodeById, nil
}

func GetSaveCode(option *model.Options) ([]model.SaveCode, error) {
	query := `query GetSaveCode($option: Options){
	    getSaveCode(option: $option){
	        id
	        code
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetSaveCode []model.SaveCode `json:"getSaveCode"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetSaveCode, nil
}

func CreateOrdonnance(input model.CreateOrdonnanceInput) (model.Ordonnance, error) {
	query := `mutation CreateOrdonnance($input: CreateOrdonnanceInput!){
	    createOrdonnance(input: $input){
	        id
	        created_by
	        patient_id
	        medicines {
	            medicine_id
	            qsp
	            qsp_unit
	            comment
	            periods {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
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
		return model.Ordonnance{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Ordonnance{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Ordonnance{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Ordonnance{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateOrdonnance model.Ordonnance `json:"createOrdonnance"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Ordonnance{}, err
	}

	if len(result.Errors) > 0 {
		return model.Ordonnance{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateOrdonnance, nil
}

func UpdateOrdonnance(id string, input model.UpdateOrdonnanceInput) (model.Ordonnance, error) {
	query := `mutation UpdateOrdonnance($id: String!, $input: UpdateOrdonnanceInput!){
	    updateOrdonnance(id: $id, input: $input){
	        id
	        created_by
	        patient_id
	        medicines {
	            medicine_id
	            qsp
	            qsp_unit
	            comment
	            periods {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
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
		return model.Ordonnance{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Ordonnance{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Ordonnance{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Ordonnance{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateOrdonnance model.Ordonnance `json:"updateOrdonnance"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Ordonnance{}, err
	}

	if len(result.Errors) > 0 {
		return model.Ordonnance{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateOrdonnance, nil
}

func DeleteOrdonnance(id string) (bool, error) {
	query := `mutation DeleteOrdonnance($id: String!){
	    deleteOrdonnance(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteOrdonnance bool `json:"deleteOrdonnance"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteOrdonnance, nil
}

func GetOrdonnanceById(id string) (model.Ordonnance, error) {
	query := `query GetOrdonnanceById($id: String!){
	    getOrdonnanceById(id: $id){
	        id
	        created_by
	        patient_id
	        medicines {
	            medicine_id
	            qsp
	            qsp_unit
	            comment
	            periods {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
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
		return model.Ordonnance{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.Ordonnance{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Ordonnance{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Ordonnance{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetOrdonnanceById model.Ordonnance `json:"getOrdonnanceById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.Ordonnance{}, err
	}

	if len(result.Errors) > 0 {
		return model.Ordonnance{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetOrdonnanceById, nil
}

func GetOrdonnanceByDoctorId(doctor_id string, option *model.Options) ([]model.Ordonnance, error) {
	query := `query GetOrdonnanceByDoctorId($doctor_id: String!, $option: Options){
	    getOrdonnanceByDoctorId(doctor_id: $doctor_id, option: $option){
	        id
	        created_by
	        patient_id
	        medicines {
	            medicine_id
	            qsp
	            qsp_unit
	            comment
	            periods {
	                quantity
	                frequency
	                frequency_ratio
	                frequency_unit
	                period_length
	                period_unit
	            }
	        }
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetOrdonnanceByDoctorId []model.Ordonnance `json:"getOrdonnanceByDoctorId"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetOrdonnanceByDoctorId, nil
}

func CreateAutoAnswer(input model.CreateAutoAnswerInput) (model.AutoAnswer, error) {
	query := `mutation CreateAutoAnswer($input: CreateAutoAnswerInput!){
	    createAutoAnswer(input: $input){
	        id
	        name
	        values
	        type
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
		return model.AutoAnswer{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AutoAnswer{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AutoAnswer{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateAutoAnswer model.AutoAnswer `json:"createAutoAnswer"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	if len(result.Errors) > 0 {
		return model.AutoAnswer{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateAutoAnswer, nil
}

func UpdateAutoAnswer(id string, input model.UpdateAutoAnswerInput) (model.AutoAnswer, error) {
	query := `mutation UpdateAutoAnswer($id: String!, $input: UpdateAutoAnswerInput!){
	    updateAutoAnswer(id: $id, input: $input){
	        id
	        name
	        values
	        type
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
		return model.AutoAnswer{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AutoAnswer{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AutoAnswer{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateAutoAnswer model.AutoAnswer `json:"updateAutoAnswer"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	if len(result.Errors) > 0 {
		return model.AutoAnswer{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateAutoAnswer, nil
}

func DeleteAutoAnswer(id string) (bool, error) {
	query := `mutation DeleteAutoAnswer($id: String!){
	    deleteAutoAnswer(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteAutoAnswer bool `json:"deleteAutoAnswer"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteAutoAnswer, nil
}

func GetAutoAnswerById(id string) (model.AutoAnswer, error) {
	query := `query GetAutoAnswerById($id: String!){
	    getAutoAnswerById(id: $id) {
	        id
	        name
	        values
	        type
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
		return model.AutoAnswer{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AutoAnswer{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AutoAnswer{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAutoAnswerById model.AutoAnswer `json:"getAutoAnswerById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	if len(result.Errors) > 0 {
		return model.AutoAnswer{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAutoAnswerById, nil
}

func GetAutoAnswerByName(name string) (model.AutoAnswer, error) {
	query := `query GetAutoAnswerByName($name: String!){
	    getAutoAnswerByName(name: $name){
	        id
	        name
	        values
	        type
	        createdAt
	        updatedAt
	    }
	}`
	variables := map[string]interface{}{
		"name": name,
	}
	reqBody := map[string]interface{}{
		"query": query,
		"variables": variables,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.AutoAnswer{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.AutoAnswer{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAutoAnswerByName model.AutoAnswer `json:"getAutoAnswerByName"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.AutoAnswer{}, err
	}

	if len(result.Errors) > 0 {
		return model.AutoAnswer{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAutoAnswerByName, nil
}

func GetAutoAnswers(option *model.Options) ([]model.AutoAnswer, error) {
	query := `query GetAutoAnswers($option: Options){
	    getAutoAnswers(option: $option){
	        id
	        name
	        values
	        type
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetAutoAnswers []model.AutoAnswer `json:"getAutoAnswers"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetAutoAnswers, nil
}

func CreateMedicalAntecedents(input model.CreateMedicalAntecedentsInput) (model.MedicalAntecedents, error) {
	query := `mutation CreateMedicalAntecedents($input: CreateMedicalAntecedentsInput!){
	    createMedicalAntecedents(input: $input){
	        id
	        name
	        symptoms
	        treatments {
	            id
	            created_by
	            start_date
	            end_date
	            medicines {
	                medicine_id
	                comment
	                period {
	                    quantity
	                    frequency
	                    frequency_ratio
	                    frequency_unit
	                    period_length
	                    period_unit
	                }
	            }
	        }
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
		return model.MedicalAntecedents{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.MedicalAntecedents{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.MedicalAntecedents{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			CreateMedicalAntecedents model.MedicalAntecedents `json:"createMedicalAntecedents"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalAntecedents{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.CreateMedicalAntecedents, nil
}

func UpdateMedicalAntecedents(id string, input model.UpdateMedicalAntecedentsInput) (model.MedicalAntecedents, error) {
	query := `mutation UpdateMedicalAntecedents($id: String!, $input: UpdateMedicalAntecedentsInput!){
	    updateMedicalAntecedents(id: $id, input: $input){
	        id
	        name
	        symptoms
	        treatments {
	            id
	            created_by
	            start_date
	            end_date
	            medicines {
	                medicine_id
	                comment
	                period {
	                    quantity
	                    frequency
	                    frequency_ratio
	                    frequency_unit
	                    period_length
	                    period_unit
	                }
	            }
	        }
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
		return model.MedicalAntecedents{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.MedicalAntecedents{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.MedicalAntecedents{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			UpdateMedicalAntecedents model.MedicalAntecedents `json:"updateMedicalAntecedents"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalAntecedents{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.UpdateMedicalAntecedents, nil
}

func DeleteMedicalAntecedents(id string) (bool, error) {
	query := `mutation DeleteMedicalAntecedents($id: String!){
	    deleteMedicalAntecedents(id: $id)
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
		return false, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			DeleteMedicalAntecedents bool `json:"deleteMedicalAntecedents"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	if len(result.Errors) > 0 {
		return false, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.DeleteMedicalAntecedents, nil
}

func GetMedicalAntecedentsById(id string) (model.MedicalAntecedents, error) {
	query := `query GetMedicalAntecedentsById($id: String!){
	    getMedicalAntecedentsById(id: $id){
	        id
	        name
	        symptoms
	        treatments {
	            id
	            created_by
	            start_date
	            end_date
	            medicines {
	                medicine_id
	                comment
	                period {
	                    quantity
	                    frequency
	                    frequency_ratio
	                    frequency_unit
	                    period_length
	                    period_unit
	                }
	            }
	        }
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
		return model.MedicalAntecedents{}, err
	}

	resp, err := http.Post(os.Getenv("GRAPHQL_URL"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return model.MedicalAntecedents{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.MedicalAntecedents{}, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetMedicalAntecedentsById model.MedicalAntecedents `json:"getMedicalAntecedentsById"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return model.MedicalAntecedents{}, err
	}

	if len(result.Errors) > 0 {
		return model.MedicalAntecedents{}, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetMedicalAntecedentsById, nil
}

func GetMedicalAntecedents(option *model.Options) ([]model.MedicalAntecedents, error) {
	query := `query GetMedicalAntecedents($option: Options){
	    getMedicalAntecedents(option: $option){
	        id
	        name
	        symptoms
	        treatments {
	            id
	            created_by
	            start_date
	            end_date
	            medicines {
	                medicine_id
	                comment
	                period {
	                    quantity
	                    frequency
	                    frequency_ratio
	                    frequency_unit
	                    period_length
	                    period_unit
	                }
	            }
	        }
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
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Errors []model.GraphQLError `json:"errors"`
		Data struct {
			GetMedicalAntecedents []model.MedicalAntecedents `json:"getMedicalAntecedents"`
		} `json:"data"`
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", result.Errors[0].Message)
	}

	return result.Data.GetMedicalAntecedents, nil
}

