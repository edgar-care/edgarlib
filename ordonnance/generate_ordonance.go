package ordonnance

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/edgar-care/edgarlib/v2/document"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/go-pdf/fpdf"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CreateOrdonnaceResponse struct {
	Ordonnance model.Ordonnance
	Url        string
	Code       int
	Err        error
}

type CreateOrdonnaceInput struct {
	PatientID string          `json:"patient_id"`
	Medicines []MedicineInput `json:"medicines"`
}

type MedicineInput struct {
	MedicineID string        `json:"medicine_id"`
	Qsp        int           `json:"qsp"`
	QspUnit    string        `json:"qsp_unit"`
	Comment    string        `json:"comment"`
	Periods    []PeriodInput `json:"periods"`
}

type PeriodInput struct {
	Quantity       int    `json:"quantity"`
	Frequency      int    `json:"frequency"`
	FrequencyRatio int    `json:"frequency_ratio"`
	FrequencyUnit  string `json:"frequency_unit"`
	PeriodLength   int    `json:"period_length"`
	PeriodUnit     string `json:"period_unit"`
}

func CreateOrdonnance(input CreateOrdonnaceInput, ownerID string) CreateOrdonnaceResponse {
	doctor, err := graphql.GetDoctorById(ownerID)
	if err != nil {
		return CreateOrdonnaceResponse{Code: 400, Err: errors.New("unable to create ordonnance: " + err.Error())}
	}

	medicines := make([]*model.MedicineOrdonnanceInput, 0)
	for _, med := range input.Medicines {
		periods := make([]*model.PeriodOrdonnanceInput, 0)
		for _, period := range med.Periods {
			var periodLength *int
			var periodUnit *model.TimeUnitEnum
			if period.PeriodLength != 0 {
				periodLength = &period.PeriodLength
			}
			if period.PeriodUnit != "" {
				unit := model.TimeUnitEnum(period.PeriodUnit)
				periodUnit = &unit
			}
			periods = append(periods, &model.PeriodOrdonnanceInput{
				Quantity:       period.Quantity,
				Frequency:      period.Frequency,
				FrequencyRatio: period.FrequencyRatio,
				FrequencyUnit:  model.TimeUnitEnum(period.FrequencyUnit),
				PeriodLength:   periodLength,
				PeriodUnit:     periodUnit,
			})
		}
		qspUnit := model.TimeUnitEnum(med.QspUnit)
		medicines = append(medicines, &model.MedicineOrdonnanceInput{
			MedicineID: med.MedicineID,
			Qsp:        med.Qsp,
			QspUnit:    qspUnit,
			Comment:    &med.Comment,
			Periods:    periods,
		})
	}

	newOrdonnance, err := graphql.CreateOrdonnance(model.CreateOrdonnanceInput{
		CreatedBy: ownerID,
		PatientID: input.PatientID,
		Medicines: medicines,
	})
	if err != nil {
		return CreateOrdonnaceResponse{Code: 400, Err: errors.New("unable to create ordonnance: " + err.Error())}
	}

	url, err := GeneratePrescriptionPDF(newOrdonnance, doctor)
	if err != nil {
		return CreateOrdonnaceResponse{Code: 400, Err: errors.New("unable to generate PDF: " + err.Error())}
	}

	_, err = graphql.UpdateDoctor(ownerID, model.UpdateDoctorInput{
		OrdonnanceIds: append(doctor.OrdonnanceIds, &newOrdonnance.ID),
	})
	if err != nil {
		return CreateOrdonnaceResponse{Code: 400, Err: errors.New("unable to update doctor: " + err.Error())}
	}

	return CreateOrdonnaceResponse{Ordonnance: newOrdonnance, Url: url, Code: 200, Err: nil}
}

func formatString(input string) string {
	return strings.ReplaceAll(strings.ToLower(input), "_", " ")
}

func formatTitleCase(input string) string {
	words := strings.Split(strings.ReplaceAll(input, "_", " "), " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func replaceDosageForm(dosageForm string, quantity int) string {
	var replacement string
	switch dosageForm {
	case "CREME", "POMMAGE", "GELE":
		replacement = "application"
	case "SOLUTION_BUVABLE", "POUDRE", "SOLUTION_INJECTABLE":
		replacement = "dose"
	case "SUSPENSION_NASALE", "SPRAY", "COLUTOIRE", "SHAMPOIN":
		replacement = "utilisation"
	case "GRANULER_EN_SACHET":
		replacement = "granulé"
	default:
		replacement = formatString(dosageForm)
	}

	if quantity > 1 {
		replacement += "s"
	}

	return replacement
}

func checkPeriod(quantity int, periods string) string {

	if quantity > 1 {
		if periods == "JOUR" {
			periods += "S"
		} else if periods == "SEMAINE" {
			periods += "S"
		} else if periods == "ANNEE" {
			periods += "S"
		} else {
			periods = periods
		}
		return periods
	}
	return periods
}

func GeneratePrescriptionPDF(prescription model.Ordonnance, doctor model.Doctor) (string, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	imageURL := "https://edgar-sante.fr/assets/logo/colored-edgar-logo.png"
	response, err := http.Get(imageURL)
	if err != nil {
		log.Fatalf("Failed to download image: %v", err)
	}
	defer response.Body.Close()

	// Step 2: Save the image to a temporary file
	tempFile, err := os.CreateTemp("", "logo-*.png")
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, response.Body)
	if err != nil {
		log.Fatalf("Failed to save image to temporary file: %v", err)
	}

	// Step 3: Use the temporary file path in the pdf.ImageOptions function
	pdf.ImageOptions(tempFile.Name(), 160, 10, 40, 0, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	// Utilisez une police de base compatible ISO-8859-1
	pdf.SetFont("Arial", "", 12)

	// Fonction pour convertir UTF-8 en ISO-8859-1
	utf8ToISO8859 := func(input string) string {
		encoder := charmap.ISO8859_1.NewEncoder()
		output, _ := encoder.String(input)
		return output
	}

	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 10, utf8ToISO8859("ORDONNANCE"))
	pdf.Ln(8)
	pdf.SetFont("Arial", "I", 12)
	pdf.Cell(0, 10, utf8ToISO8859("Établissement de traitement"))
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(100, 10, utf8ToISO8859(fmt.Sprintf("Dr.%s %s", doctor.Name, doctor.Firstname)))
	pdf.Ln(6)
	pdf.Cell(100, 10, utf8ToISO8859(doctor.Address.Street+", "+doctor.Address.ZipCode+" "+doctor.Address.City+", "+doctor.Address.Country))
	pdf.Cell(0, 10, utf8ToISO8859(fmt.Sprintf("%s", time.Unix(int64(prescription.CreatedAt), 0).Format("02-Jan-2006"))))
	pdf.Ln(12)

	// Section Détails du Patient
	patient, err := graphql.GetPatientById(prescription.PatientID)
	if err != nil {
		return "", errors.New("unable to retrieve patient: " + err.Error())
	}
	info_patient, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return "", errors.New("unable to retrieve patient: " + err.Error())
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "", 12)
	name_patient := fmt.Sprintf("%s %s", info_patient.Name, info_patient.Firstname)
	pdf.Cell(0, 10, utf8ToISO8859(name_patient))
	pdf.Ln(5)
	birthdate := time.Unix(int64(info_patient.Birthdate), 0)
	formattedBirthdate := birthdate.Format("02-Jan-2006")
	pdf.Cell(100, 10, utf8ToISO8859(formattedBirthdate))

	//pdf.Cell(100, 10, utf8ToISO8859(fmt.Sprintf("%s", strconv.Itoa(info_patient.Birthdate))))
	pdf.Ln(12)
	pdf.Line(10, 55, 200, 55) // Ligne horizontale

	// Section Médicaments
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 20, utf8ToISO8859("MÉDICAMENTS"))
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	for _, med := range prescription.Medicines {
		medication, err := graphql.GetMedicineByID(med.MedicineID)
		if err != nil {
			return "", errors.New("unable to retrieve medication: " + err.Error())
		}

		pdf.Ln(6)
		pdf.SetFont("Arial", "", 12)
		qspInfo := fmt.Sprintf("QSP %d %s", med.Qsp, checkPeriod(med.Qsp, string(med.QspUnit)))
		pdf.Cell(0, 10, utf8ToISO8859(qspInfo))
		pdf.Ln(6)
		pdf.SetFont("Arial", "B", 12)
		medicationInfo := fmt.Sprintf("%s (%s %s %s/%d%s)", medication.Dci, medication.Name, formatTitleCase(string(medication.DosageForm)), formatTitleCase(string(medication.Container)), medication.Dosage, medication.DosageUnit)
		pdf.Cell(0, 10, utf8ToISO8859(medicationInfo))
		pdf.Ln(6)
		if med.Comment != nil {
			pdf.SetFont("Arial", "I", 11)
			comment := fmt.Sprintf("Commentaire: %s", *med.Comment)
			pdf.Cell(0, 10, utf8ToISO8859(comment))
		}
		pdf.SetFont("Arial", "I", 11)
		pdf.Ln(6)

		var periodDescriptions []string
		for _, period := range med.Periods {
			periodInfo := fmt.Sprintf("%d %s à prendre %d fois", period.Quantity, replaceDosageForm(string(medication.DosageForm), period.Quantity), period.Frequency)
			if period.FrequencyRatio == 1 {
				if period.PeriodLength == nil || period.PeriodUnit == nil {
					periodInfo += fmt.Sprintf(" par %s ", formatString(checkPeriod(period.FrequencyRatio, string(period.FrequencyUnit))))
				} else {
					periodInfo += fmt.Sprintf(" par %s pendant %d %s", formatString(checkPeriod(period.FrequencyRatio, string(period.FrequencyUnit))), *period.PeriodLength, formatString(checkPeriod(*period.PeriodLength, string(*period.PeriodUnit))))
				}
			} else if period.FrequencyRatio > 1 {
				if period.PeriodLength == nil || period.PeriodUnit == nil {
					periodInfo += fmt.Sprintf(" tous les %d %s", period.FrequencyRatio, formatString(checkPeriod(period.FrequencyRatio, string(period.FrequencyUnit))))

				} else {
					periodInfo += fmt.Sprintf(" tous les %d %s pendant %d %s", period.FrequencyRatio, formatString(checkPeriod(period.FrequencyRatio, string(period.FrequencyUnit))), *period.PeriodLength, formatString(checkPeriod(*period.PeriodLength, string(*period.PeriodUnit))))
				}
			}
			periodDescriptions = append(periodDescriptions, periodInfo)
			//}
		}
		if len(periodDescriptions) > 1 {
			for i, desc := range periodDescriptions {
				if i == 0 {
					pdf.Cell(0, 10, utf8ToISO8859(desc))
					pdf.Ln(6)
					continue
				}

				pdf.Cell(0, 10, "puis "+utf8ToISO8859(desc))
				pdf.Ln(6)
			}
		} else if len(periodDescriptions) == 1 {
			pdf.Cell(0, 10, utf8ToISO8859(periodDescriptions[0]))
			pdf.Ln(6)
		}
		pdf.Ln(8)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return "", err
	}
	nameFileS3 := prescription.ID + ".pdf"
	_, err = UploadOrdonnance(&buf, nameFileS3)
	if err != nil {
		return "", err
	}

	custom_name := prescription.ID + "-" + info_patient.Name + ".pdf"

	url, err := document.GenerateURL("doctor-ordonnance", nameFileS3, custom_name)
	if err != nil {
		return "", errors.New("error generating signed URL")
	}

	return url, nil
}

func UploadOrdonnance(file io.Reader, filename string) (string, error) {
	awsAccessKeyID := os.Getenv("EDGAR__ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("EDGAR_SECRET_ACCESS_KEY")
	region := os.Getenv("EDGAR_REGION")
	bucketName := "doctor-ordonnance"

	// Validate the filename
	if filename == "" {
		return "", fmt.Errorf("empty filename")
	}

	// Check for an empty file
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	if buf.Len() == 0 {
		return "", fmt.Errorf("empty file content")
	}

	// Set the content type based on the file extension
	contentType := "application/octet-stream"
	if ext := filepath.Ext(filename); ext != "" {
		contentType = mime.TypeByExtension(ext)
	}

	cfg := aws.NewConfig().WithRegion(region)
	if awsAccessKeyID != "" && awsSecretAccessKey != "" {
		cfg = cfg.WithCredentials(credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""))
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		return "", fmt.Errorf("error creating session: %w", err)
	}

	s3Client := s3.New(sess)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		log.Printf("Error uploading file to S3: %v", err)
		return "", fmt.Errorf("error uploading file to S3: %w", err)
	}

	downloadURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)

	return downloadURL, nil
}
