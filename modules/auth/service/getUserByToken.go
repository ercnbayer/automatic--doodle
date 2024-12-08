package service

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"

	"github.com/google/uuid"
)

func (svc *Service) GetUserByToken(
	token *string,
) (types.AuthenticatedUser, error) {
	tokens, err := svc.encryptionService.DecryptTokens(*token)
	if err != nil {
		return types.AuthenticatedUser{},
			errors.NewUnauthorizedError("AuthenticationService")
	}

	userId, err := svc.accessTokenService.Validate(tokens.UserRole, &tokens.AccessToken)
	if err != nil {
		return types.AuthenticatedUser{},
			errors.NewUnauthorizedError("AuthenticationService")
	}

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return types.AuthenticatedUser{},
			errors.NewUnauthorizedError("AuthenticationService")
	}

	userItem, err := svc.userRepository.GetById(userUuid, context.Background())

	return types.AuthenticatedUser{
		Id:        userItem.ID,
		FirstName: userItem.FirstName,
		LastName:  userItem.LastName,
		Email:     userItem.Email,
		Role:      userItem.Role,
		State:     userItem.State,
	}, nil
}
