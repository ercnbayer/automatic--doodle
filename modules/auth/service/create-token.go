package service

import (
	"automatic-doodle/ent/user"
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"

	"github.com/google/uuid"
)

func (svc *Service) CreateTokens(
	userId uuid.UUID,
	userRole user.Role,
) (string, error) {
	jwtPayload := types.JWTTokenPayload{
		UserRole: user.RoleCUSTOMER,
		UserId:   userId,
	}
	accessToken, err := svc.accessTokenService.Create(&jwtPayload)
	if err != nil {
		// TODO: add something clear
		svc.log.Error(
			`Something went wrong during CreateTokens->CreateJWTToken: %w`,
			err,
		)
		return "", errors.NewUnauthorizedError("Something went wrong!")
	}

	refreshTokenCreate := svc.refreshTokenFactory.Create(userId)
	refreshTokenRow, err := svc.refreshTokenRepository.Create(
		refreshTokenCreate,
		context.Background(),
	)
	if err != nil {
		// TODO: add something clear
		svc.log.Error(
			`Something went wrong during CreateTokens->CreateRefreshToken: %w`,
			err,
		)
		return "", errors.NewUnauthorizedError("Something went wrong!")
	}

	tokens, err := svc.encryptionService.EncryptTokens(types.EncryptedTokenPayload{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenRow.Token,
		UserRole:     userRole,
	})
	if err != nil {
		// TODO: add something clear
		svc.log.Error(
			`Something went wrong during CreateTokens->EncryptTokens: %w`,
			err,
		)
		return "", errors.NewUnauthorizedError("Something went wrong!")
	}

	return tokens, nil
}
