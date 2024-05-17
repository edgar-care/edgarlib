package document

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func generateSignedURL(bucketName, objectKey string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("erreur lors du chargement du fichier .env: %w", err)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("EDGAR_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("EDGAR_ACCESS_KEY_ID"),
			os.Getenv("EDGAR_SECRET_ACCESS_KEY"),
			"",
		),
	}))

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	url, err := req.Presign(20 * time.Minute)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la génération de l'URL signée: %w", err)
	}

	return url, nil
}
