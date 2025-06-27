package service

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"
	"fmt"
)

func (srv *Service) Login(req *types.LoginRequest) (types.TokenResponse, error) {

	user, err := srv.userRepository.GetByIdentifier(req.Email, context.Background())

	if err != nil {
		fmt.Println("Error getting user:", err)
		return types.TokenResponse{}, err
	}

	if ok := srv.encryptionService.CheckPasswordHash(
		req.Password,
		user.PasswordSalt,
		user.PasswordHash,
	); !ok {
		fmt.Println("Password is incorrect")
		return types.TokenResponse{},
			errors.NewUnauthorizedError("AuthenticationService")
	}

	token, err := srv.CreateTokens(user.ID, user.Role)

	if err != nil {
		fmt.Println("Error creating tokens:", err)
		return types.TokenResponse{}, err
	}
	return types.TokenResponse{Token: token, User: types.AuthenticatedUserFromUser(user)}, nil
}
