package service

import (
	"automatic-doodle/ent/user"
	"automatic-doodle/pkg/errors"

	_jwt "github.com/golang-jwt/jwt/v5"
)

func (svc *Service) Validate(
	userRole user.Role,
	pToken *string,
) (string, error) {
	pCreds, err := svc.getCredsByUserRole(&userRole)
	if err != nil {
		svc.log.Error("Creds not found for user type: %s", userRole)
		return "", err
	}

	token, err := _jwt.Parse(*pToken, func(jwtToken *_jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*_jwt.SigningMethodHMAC); !ok {
			svc.log.Error("Signing method is not correct!")
			return nil, errors.New("JWTModule", "signing method is incorrect")
		}

		return []byte(pCreds.Key), nil
	})
	if err != nil {
		return "", err
	}

	claims, _ := token.Claims.(_jwt.MapClaims)

	if claims["aud"].(string) != pCreds.Aud {
		svc.log.Error(
			"Tokens audiance is not correct! User Type: %s, Expected Audiance: %, Audiance: %s",
			userRole,
			pCreds.Aud,
			claims["aud"].(string),
		)
		return "", errors.NewUnauthorizedError("JWTModule")
	}

	return claims["sub"].(string), nil
}
