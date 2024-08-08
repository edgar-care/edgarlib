package double_auth

import (
	"github.com/edgar-care/edgarlib/graphql"
	"testing"
)

func TestRemoveDoubleAuthMethod_Success(t *testing.T) {
	// Configurer les valeurs de test
	methodToRemove := "sms"
	patientId := "valid_patient_id"

	// Appel de la fonction
	err := RemoveDoubleAuthMethod(methodToRemove, patientId)

	// Vérifier les résultats
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Vérifier que la méthode a bien été supprimée
	patient, _ := graphql.GetPatientById(patientId)
	doubleAuth, _ := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
	for _, method := range doubleAuth.Methods {
		if method == methodToRemove {
			t.Errorf("Expected method %s to be removed, but it still exists", methodToRemove)
		}
	}
}

func TestRemoveDoubleAuthMethod_PatientNotFound(t *testing.T) {
	// Configurer les valeurs de test
	methodToRemove := "sms"
	invalidPatientId := "invalid_patient_id"

	// Appel de la fonction
	err := RemoveDoubleAuthMethod(methodToRemove, invalidPatientId)

	// Vérifier les résultats
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	if err.Error() != "unable to fetch patient" {
		t.Errorf("Expected error message 'unable to fetch patient', got: %v", err)
	}
}

func TestRemoveDoubleAuthMethod_NoDoubleAuthMethodsID(t *testing.T) {
	// Configurer les valeurs de test
	methodToRemove := "sms"
	patientId := "patient_without_double_auth_methods_id"

	// Appel de la fonction
	err := RemoveDoubleAuthMethod(methodToRemove, patientId)

	// Vérifier les résultats
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	if err.Error() != "patient does not have double_auth_methods_id" {
		t.Errorf("Expected error message 'patient does not have double_auth_methods_id', got: %v", err)
	}
}

func TestRemoveDoubleAuthMethod_MethodNotFound(t *testing.T) {
	// Configurer les valeurs de test
	methodToRemove := "non_existing_method"
	patientId := "valid_patient_id"

	// Appel de la fonction
	err := RemoveDoubleAuthMethod(methodToRemove, patientId)

	// Vérifier les résultats
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	if err.Error() != "method to remove not found in current methods" {
		t.Errorf("Expected error message 'method to remove not found in current methods', got: %v", err)
	}
}

func TestRemoveDoubleAuthMethod_RemoveLastMethod(t *testing.T) {
	// Configurer les valeurs de test
	methodToRemove := "sms"
	patientId := "patient_with_single_double_auth_method"

	// Appel de la fonction
	err := RemoveDoubleAuthMethod(methodToRemove, patientId)

	// Vérifier les résultats
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Vérifier que le champ DoubleAuthMethodsID est bien supprimé
	patient, _ := graphql.GetPatientById(patientId)
	if patient.DoubleAuthMethodsID != nil && *patient.DoubleAuthMethodsID != "" {
		t.Errorf("Expected DoubleAuthMethodsID to be removed, but it still exists: %v", *patient.DoubleAuthMethodsID)
	}
}
