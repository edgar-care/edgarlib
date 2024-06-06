package medical_folder

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateMedicalInfo(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

}
