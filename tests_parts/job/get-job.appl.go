package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type AuthenticatedUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetJobApplResponse struct {
	User        AuthenticatedUser `json:"user"`
	FileKey     string            `json:"fileKey"`
	BucketName  string            `json:"bucketName"`
	Description string            `json:"description"`
}

type GetJobApplQuery struct {
	JobID      uuid.UUID `query:"id" validate:"required,uuid"`
	PageNumber int       `query:"pageNumber" validate:"gte=1"`
}

// FetchAndParseJobApplications sends an HTTP GET request with query parameters and parses the response
func FetchAndParseJobApplications(baseURL string, query GetJobApplQuery) ([]GetJobApplResponse, error) {
	// Build query parameters
	params := url.Values{}
	params.Set("id", query.JobID.String())
	params.Set("pageNumber", fmt.Sprintf("%d", query.PageNumber))

	// Construct the full URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the HTTP GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP GET request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var responses []GetJobApplResponse
	if err := json.Unmarshal(body, &responses); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return responses, nil
}

func main() {
	// Example query parameters
	query := GetJobApplQuery{
		JobID:      uuid.New(),
		PageNumber: 2,
	}

	// Base URL for the API
	apiURL := "http://localhost:8080/get-job-applications"

	// Fetch and parse job applications
	responses, err := FetchAndParseJobApplications(apiURL, query)
	if err != nil {
		panic(fmt.Sprintf("Error fetching or parsing job applications: %v", err))
	}

	// Print the parsed data
	for _, item := range responses {
		fmt.Printf("User: %s (%s), FileKey: %s, BucketName: %s, Description: %s\n",
			item.User.Name, item.User.ID, item.FileKey, item.BucketName, item.Description)
	}
}
