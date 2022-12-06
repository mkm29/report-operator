package aws

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSClient struct {
	// AWS session
	S3       *s3.S3
	Uploader *s3manager.Uploader
}

func (c *AWSClient) Init() {
	// Create a session with the default configuration
	// get reguin from environment variable
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		log.Fatal("AWS_DEFAULT_REGION environment variable not set")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatal(err)
	}
	uploader := s3manager.NewUploader(sess)
	// Create S3 service client
	s3client := s3.New(sess)
	c.S3 = s3client
	c.Uploader = uploader
}

func (c *AWSClient) Upload(key string, bs []byte) (error, *string) {
	// check if S3 has been initialized
	if c.S3 == nil {
		// initialize S3
		c.Init()
	}

	// Get bucket name from environment variable
	bucket := os.Getenv("S3_BUCKET")
	// check if bucket is empty
	if bucket == "" {
		log.Fatal("S3_BUCKET environment variable is not set")
		return errors.New("S3_BUCKET environment variable is not set"), nil
	}

	res, err := c.S3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String("application/json"),
		Body:        bytes.NewReader(bs),
	})
	if err != nil {
		log.Fatalf("Unable to upload %q to %q, %v", key, bucket, err)
		return err, nil
	}

	return nil, res.VersionId
}
