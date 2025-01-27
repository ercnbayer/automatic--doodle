package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

//  represents a mock implementation for testing.

// Mock methods for .
func GetByIdentifier() (User, error) {
	return User{
		UserID:   uuid.NewString(),
		UserRole: "USER",
	}, nil
}

func GetByIdentifierF() (User, error) {
	return User{
		UserID:   uuid.NewString(),
		UserRole: "USER",
	}, errors.New("some err")
}

func CheckPasswordHash(password, hashedPassword string) bool {
	// Simulate password hash checking.
	return hashedPassword == "return True"
}

func CheckPasswordHashF(password, hashedPassword string) bool {
	return false
}

func CreateTokens(userID, userRole string) (string, error) {
	// Simulate token creation.
	return "mockToken", nil
}

func CreateTokenFalse(userID, userRole string) (string, error) {
	// Simulate token creation.
	return "mockToken", errors.New("simulate error")
}

// User structure to simulate a user.
type User struct {
	UserID   string
	UserRole string
}

// TokenResponse structure for the test.
type TokenResponse struct {
	Token string
	User  AuthenticatedUser
}

// AuthenticatedUser represents authenticated user information.
type AuthenticatedUser struct {
	Id           string
	FirstName    string
	LastName     string
	Email        string
	ProfilePhoto FileResponse
	Role         string
	State        string
}

// FileResponse represents a file's details.
type FileResponse struct {
	Key    string
	Bucket string
}

func TestLogin(t *testing.T) {

	// Simulate getting a user by identifier.
	user, err := GetByIdentifier()

	if err != nil {
		t.Error("Get By Identifier Fail")
		return
	}

	// Validate password.
	password := "password123"
	if !CheckPasswordHash(password, "return True") {
		t.Error("Password hash check failed")
		return
	}

	// Create tokens for the user.
	token, err := CreateTokens(user.UserID, user.UserRole)
	if err != nil {
		t.Errorf("Failed to create tokens: %v", err)
		return
	}

	// Construct expected response.
	_, err = TokenResponse{
		Token: token,
		User: AuthenticatedUser{
			Id:        user.UserID,
			FirstName: "Test",
			LastName:  "Tester",
			Email:     "Testmail@mail.com",
			ProfilePhoto: FileResponse{
				Key:    "somefile",
				Bucket: "somebucket",
			},
			Role:  user.UserRole,
			State: "ACTIVE",
		},
	}, nil

}

func TestLoginGetByIdFail(t *testing.T) {

	// Simulate getting a user by identifier.
	user, err := GetByIdentifierF()

	if err != nil {
		t.Log("Get By Identifier Fail So ")
		return
	}

	// Validate password.
	password := "password123"
	if !CheckPasswordHash(password, "return True") {
		t.Log("Password hash check failed, so it passed")
		return
	}

	t.Fail()

	// Create tokens for the user.
	token, err := CreateTokens(user.UserID, user.UserRole)
	if err != nil {
		t.Errorf("Failed to create tokens: %v", err)
		return
	}

	// Construct expected response.
	_, err = TokenResponse{
		Token: token,
		User: AuthenticatedUser{
			Id:        user.UserID,
			FirstName: "Test",
			LastName:  "Tester",
			Email:     "Testmail@mail.com",
			ProfilePhoto: FileResponse{
				Key:    "somefile",
				Bucket: "somebucket",
			},
			Role:  user.UserRole,
			State: "ACTIVE",
		},
	}, nil

}

func TestLoginByCheckTokenFail(t *testing.T) {

	// Simulate getting a user by identifier.
	user, err := GetByIdentifier()

	if err != nil {
		t.Error("Get By Identifier Fail")
		return
	}

	// Validate password.
	password := "password123"
	if !CheckPasswordHash(password, "return True") {
		t.Error("Password hash check failed")
		return
	}

	// Create tokens for the user.
	token, err := CreateTokenFalse(user.UserID, user.UserRole)
	if err != nil {
		t.Logf("Failed to create tokens: %v", err)
		return
	}
	t.Fail()

	// Construct expected response.
	_, err = TokenResponse{
		Token: token,
		User: AuthenticatedUser{
			Id:        user.UserID,
			FirstName: "Test",
			LastName:  "Tester",
			Email:     "Testmail@mail.com",
			ProfilePhoto: FileResponse{
				Key:    "somefile",
				Bucket: "somebucket",
			},
			Role:  user.UserRole,
			State: "ACTIVE",
		},
	}, nil

}

func TestLoginHashFalse(t *testing.T) {

	// Simulate getting a user by identifier.
	user, err := GetByIdentifier()

	if err != nil {
		t.Error("Get By Identifier Fail")
		return
	}

	// Validate password.
	password := "password123"
	if !CheckPasswordHash(password, "") {
		t.Log("Password hash check failed, so it passed")
		return
	}
	t.Fail()

	// Create tokens for the user.
	token, err := CreateTokens(user.UserID, user.UserRole)
	if err != nil {
		t.Errorf("Failed to create tokens: %v", err)
		return
	}

	// Construct expected response.
	_, err = TokenResponse{
		Token: token,
		User: AuthenticatedUser{
			Id:        user.UserID,
			FirstName: "Test",
			LastName:  "Tester",
			Email:     "Testmail@mail.com",
			ProfilePhoto: FileResponse{
				Key:    "somefile",
				Bucket: "somebucket",
			},
			Role:  user.UserRole,
			State: "ACTIVE",
		},
	}, nil

}
