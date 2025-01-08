package test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type JobQuery struct {
	Identifier string `query:"params"`
	PageNumber int    `query:"pageNumber"`
}

func buildQuery(query JobQuery) string {
	// Create a new url.Values object to store query parameters
	params := url.Values{}

	// Set the 'params' query parameter (Identifier)
	params.Set("params", query.Identifier)

	// Set the 'pageSize' query parameter
	params.Set("pageNumber", fmt.Sprintf("%d", query.PageNumber))

	// Return the encoded query string
	return params.Encode()
}

func BrowseJob(authToken, baseURL string, jobQuery JobQuery) error {
	// Parse the base URL and append query parameters
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	// Add query parameters from JobQuery struct
	queryString := buildQuery(jobQuery)
	parsedURL.RawQuery = queryString

	// Create the GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth", authToken)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read and handle the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s, response: %s", resp.Status, string(respBody))
	}

	fmt.Printf("Request successful. Response: %s\n", string(respBody))
	return nil
}
