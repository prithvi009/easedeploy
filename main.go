package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/prithvi009/easedeploy/config"
)

func main() {

	client := config.S3Instance()
	// Specify the bucket name
	bucketName := "deployease"

	// List local files to upload
	files := []string{"file3.txt"}

	for _, file := range files {
		// Open the local file
		f, err := os.Open(file)
		if err != nil {
			log.Printf("failed to open file %s: %v\n", file, err)
			continue
		}
		defer f.Close()

		// Upload the file to S3
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(file), // Use the same key as the filename
			Body:   f,
		})
		if err != nil {
			log.Printf("failed to upload file %s: %v\n", file, err)
			continue
		}

		log.Printf("file %s uploaded successfully\n", file)
	}
}
