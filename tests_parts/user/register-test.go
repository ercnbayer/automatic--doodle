package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendUserRegistrationRequest() {

	// Define and marshal the request payload in one step
	jsonData, err := json.Marshal(map[string]string{
		"firstName":   "John",
		"lastName":    "Doe",
		"email":       "john.doe@example.com",
		"phoneNumber": "+1234567890",
		"password":    "P@ssw0rd!",
	})
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Send the POST request
	resp, err := http.Post("http://localhost:4000/auth/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("Response: %s\n", body)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string            `json:"token"`
	User  AuthenticatedUser `json:"user"`
}

type AuthenticatedUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

var LoginReq = LoginRequest{
	Email:    "john.doe@example.com",
	Password: "P@ssw0rd!",
}

func LoginAndGetTokens(url string, loginReq LoginRequest) (*TokenResponse, error) {
	// Marshal the request into JSON
	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		return nil, errors.New("failed to marshal login request: " + err.Error())
	}

	// Send the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errors.New("failed to send login request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("login failed with status code: " + http.StatusText(resp.StatusCode))
	}

	// Decode the response body
	var tokenResp TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return nil, errors.New("failed to decode response: " + err.Error())
	}

	return &tokenResp, nil
}
