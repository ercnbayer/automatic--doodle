package service

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"encoding/json"
	"strconv"
	"time"

	_jwt "github.com/golang-jwt/jwt/v5"
)

func (svc *Service) Create(
	pTokenPayload *types.JWTTokenPayload,
) (string, error) {
	pCreds, err := svc.getCredsByUserRole(&pTokenPayload.UserRole)
	if err != nil {
		return "", err
	}
	pCreds.UserId = pTokenPayload.UserId

	claims := _jwt.MapClaims{
		"exp": json.Number(strconv.FormatInt(time.Now().Add(pCreds.ExpireDuration).Unix(), 10)),
		"iss": "arrndme-jwt-issuer",
		"aud": pCreds.Aud,
		"sub": pCreds.UserId,
	}
	token := _jwt.NewWithClaims(_jwt.SigningMethodHS512, claims)

	var signedToken string
	if signedToken, err = token.SignedString([]byte(pCreds.Key)); err != nil {
		svc.log.Error("Something went wrong during sign token: %s", err)
		return "", errors.New("JWTModule", "sign token failed")
	}

	return signedToken, nil
}
