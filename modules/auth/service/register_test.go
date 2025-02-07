package service

import (
	"automatic-doodle/types"
	"testing"
)

func TestRegisterSuccess(t *testing.T) {
	srv := Service{
		log:                    &LoggerMock{},
		encryptionService:      &EncryptionMock{},
		userFactory:            &UserFacMock{},
		refreshTokenFactory:    &RefreshTokenFactoryMock{},
		accessTokenService:     &AccessTokenMock{},
		userRepository:         &UserRepoMock{},
		refreshTokenRepository: &RefreshTokenRepositoryMock{},
	}

	token, err := srv.Register(&types.RegisterRequest{})

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(token)
}
