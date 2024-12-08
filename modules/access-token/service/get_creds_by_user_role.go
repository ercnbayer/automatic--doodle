package service

import (
	"automatic-doodle/ent/user"
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
)

func (svc *Service) getCredsByUserRole(
	pUserType *user.Role,
) (*types.JWTCredentials, error) {
	switch *pUserType {
	case user.RoleADMIN:
		return &svc.adminCreds, nil
	case user.RoleCUSTOMER:
		return &svc.customerCreds, nil
	default:
		svc.log.Error("Invalid jwt token create param as user type: %s", *pUserType)
		return nil, errors.New("JWTModule", "invalid user type")
	}
}
