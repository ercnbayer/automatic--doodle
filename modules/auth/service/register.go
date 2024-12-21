package service

import (
	"automatic-doodle/ent/user"
	"automatic-doodle/types"
	"context"
)

func (srv *Service) Register(payload *types.RegisterRequest) (types.TokenResponse, error) {

	userRow := srv.userFactory.Create(payload.PhoneNumber, payload.Email, payload.FirstName, payload.LastName, payload.Password, user.RoleCUSTOMER, user.StateACTIVE)

	user, err := srv.userRepository.Create(userRow, context.Background())

	if err != nil {
		srv.log.Info("Err:%s", err)
		return types.TokenResponse{}, err
	}

	tokens, err := srv.CreateTokens(user.ID, user.Role)

	if err != nil {
		srv.log.Info("Err:%s", err)
		return types.TokenResponse{}, err
	}

	return types.TokenResponse{User: types.AuthenticatedUserFromUser(user), Token: tokens}, nil
}
