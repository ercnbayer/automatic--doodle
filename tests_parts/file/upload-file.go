package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// CreateUploadUrlRequest represents the request payload
type CreateUploadUrlRequest struct {
	Filename string `json:"fileName" validate:"required,filename,fileext"`
	Type     string `json:"type" validate:"required,oneof=PROFILE_IMAGE COVER_IMAGE POST_FILE"`
}

// CreateUploadUrlResponse represents the response structure
type CreateUploadUrlResponse struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

func UploadToLocalStackS3(endpoint, bucketName, filePath, objectKey string) error {
	// Load AWS configuration with a custom endpoint

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-2"), // Required for S3 even with LocalStack
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == s3.ServiceID {
					return aws.Endpoint{
						URL: endpoint,
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			},
		)),
	)

	if err != nil {
		panic(fmt.Errorf("unable to load SDK config: %w", err))
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)

	// Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file %q: %w", filePath, err)
	}
	defer file.Close()

	// Read the file content into a buffer
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read file content: %w", err)
	}

	// Upload the file to LocalStack S3
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String("jpg"), // Update content type as needed
	})

	if err != nil {

		fmt.Errorf("\n\nunable to upload file: %w", err)
		panic(err)
	}

	return nil
}

func CreateUploadUrl(filename, fileType, authToken string) (*CreateUploadUrlResponse, error) {
	// Prepare the request data
	requestData := CreateUploadUrlRequest{
		Filename: filename,
		Type:     fileType,
	}

	// Marshal the request data into JSON
	reqBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	// Create a new HTTP request
	apiURL := "http://localhost:4000/files" // Replace with your API URL
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth", authToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Check for non-2xx HTTP responses
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("non-2xx response: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response JSON into the response struct
	var response CreateUploadUrlResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}
