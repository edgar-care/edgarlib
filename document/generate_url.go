package document

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/url"
	"time"
)

func generateURL(bucket string, key string, customFilename string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-3")},
	)
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	// Add custom filename as a query parameter
	req.HTTPRequest.URL.RawQuery = url.Values{
		"response-content-disposition": []string{"attachment; filename=\"" + customFilename + "\""},
	}.Encode()

	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}
