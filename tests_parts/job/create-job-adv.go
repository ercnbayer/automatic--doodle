package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type JobAdvRequest struct {
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required,gtefield=StartDate"`
	Fee         int       `json:"fee" validate:"required,min=1"`
	JobType     string    `json:"job_type" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,max=1024"`
}
type FileResponse struct {
	Key    string `json:"key"`
	Bucket string `json:"bucket"`
}

type UserPublicDetails struct {
	Id           uuid.UUID    `json:"id"`
	FirstName    string       `json:"publisherfirstName"`
	LastName     string       `json:"lastName"`
	ProfilePhoto FileResponse `json:"profilePhoto"`
}
type JobAdvResponse struct {
	Publisher   UserPublicDetails `json:"publisher"`
	StartDate   time.Time         `json:"start_date" validate:"required"`
	EndDate     time.Time         `json:"end_date" validate:"required,gtefield=StartDate"`
	Fee         int               `json:"fee" validate:"required,min=1"`
	JobType     string            `json:"job_type" validate:"required,max=255"`
	Description string            `json:"description" validate:"required,max=1024"`
}

func SendJobAdvRequest(authToken string) (*JobAdvResponse, error) {
	// Convert the JobAdvRequest struct to JSON
	jobAdvRequest := JobAdvRequest{
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		Fee:         100,
		JobType:     "Full-Time",
		Description: "A great job opportunity.",
	}
	data, err := json.Marshal(jobAdvRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "http://localhost:4000/auth/register", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Add X-Auth token to the header
	req.Header.Set("X-Auth", authToken)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// Parse the response body into JobAdvResponse
	var jobAdvResponse JobAdvResponse
	if err := json.NewDecoder(resp.Body).Decode(&jobAdvResponse); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return nil, err
	}

	// Print the JobAdvResponse
	fmt.Printf("Job Advertisement Response:\n")
	fmt.Printf("Publisher: %s %s\n", jobAdvResponse.Publisher.FirstName, jobAdvResponse.Publisher.LastName)
	fmt.Printf("Start Date: %s\n", jobAdvResponse.StartDate)
	fmt.Printf("End Date: %s\n", jobAdvResponse.EndDate)
	fmt.Printf("Fee: %d\n", jobAdvResponse.Fee)
	fmt.Printf("Job Type: %s\n", jobAdvResponse.JobType)
	fmt.Printf("Description: %s\n", jobAdvResponse.Description)
	return &jobAdvResponse, nil
}
