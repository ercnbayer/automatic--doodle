package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"automatic-doodle/types"
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

type RefreshTokenRepositoryMock struct {
}

func (*RefreshTokenRepositoryMock) Create(
	*ent.RefreshTokenCreate,
	context.Context,
) (*ent.RefreshToken, error) {
	return &ent.RefreshToken{}, nil
}

func (*RefreshTokenRepositoryMock) GetByTokenWithUser(string, context.Context) (*ent.RefreshToken, error) {
	return &ent.RefreshToken{}, nil
}

type AuthTestStruct struct {
	Srv *Service
}

type EncryptionMock struct {
}

func (*EncryptionMock) EncryptTokens(
	tokens types.EncryptedTokenPayload,
) (string, error) {
	return "TestToken", nil
}

func (*EncryptionMock) DecryptTokens(string) (types.EncryptedTokenPayload, error) {
	return types.EncryptedTokenPayload{}, nil
}
func (*EncryptionMock) CheckPasswordHash(string, string, string) bool {
	return true
}

type AccessTokenMock struct {
}

func (*AccessTokenMock) Create(*types.JWTTokenPayload) (string, error) {

	return "tokencreated", nil
}

func (*AccessTokenMock) Validate(user.Role, *string) (string, error) {
	return "validated", nil
}

type UserFacMock struct {
}

func (*UserFacMock) Create(
	phoneNumber,
	email,
	firstName,
	lastName,

	password string,
	role user.Role,
	state user.State,
) *ent.UserCreate {
	return &ent.UserCreate{}
}

func (*UserFacMock) UpdatePassword(id uuid.UUID, password string, ctx context.Context) (*ent.User, error) {
	return &ent.User{}, nil
}

type UserRepoMock struct {
}

func (*UserRepoMock) Create(*ent.UserCreate, context.Context) (*ent.User, error) {
	return &ent.User{}, nil
}

func (*UserRepoMock) GetById(uuid.UUID, context.Context) (*ent.User, error) {
	return &ent.User{}, nil
}
func (*UserRepoMock) DeleteUser(uuid.UUID, context.Context) error {
	return nil

}
func (*UserRepoMock) GetByIdentifier(string, context.Context) (*ent.User, error) {
	return &ent.User{}, nil

}

type AccessTokenRepositoryMock struct {
}

func (*AccessTokenRepositoryMock) Validate(user.Role, *string) (string, error) {
	return "validated", nil
}

type RefreshTokenFactoryMock struct {
}

func (*RefreshTokenFactoryMock) Create(uuid.UUID) *ent.RefreshTokenCreate {
	return &ent.RefreshTokenCreate{}
}

type LoggerMock struct {
}

func (*LoggerMock) Trace(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Fatal(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Warning(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Info(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Error(format string, args ...any) {
	fmt.Sprintf(format, args...)
}

func TestLoginSuccess(t *testing.T) {
	srv := Service{
		log:                    &LoggerMock{},
		encryptionService:      &EncryptionMock{},
		userFactory:            &UserFacMock{},
		refreshTokenFactory:    &RefreshTokenFactoryMock{},
		accessTokenService:     &AccessTokenMock{},
		userRepository:         &UserRepoMock{},
		refreshTokenRepository: &RefreshTokenRepositoryMock{},
	}

	token, err := srv.Login(&types.LoginRequest{})

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(token)
}
