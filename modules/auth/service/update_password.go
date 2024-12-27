package service

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"

	"github.com/google/uuid"
)

func (srv *Service) UpdatePassword(req types.UpdateUserPassword, id uuid.UUID) (types.AuthenticatedUser, error) {
	ctx := context.Background()

	user, err := srv.userRepository.GetById(id, ctx)

	if err != nil {
		return types.AuthenticatedUser{}, err
	}

	check := srv.encryptionService.CheckPasswordHash(req.CurrentPassword, user.PasswordSalt, user.PasswordHash)

	if !check {
		return types.AuthenticatedUser{}, errors.New("Update Password", "Wrong Password")
	}

	user, err = srv.userFactory.UpdatePassword(id, req.NewPassword, ctx)

	if err != nil {
		return types.AuthenticatedUser{}, err
	}
	return types.AuthenticatedUserFromUser(user), nil

}
