package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/tatsuya06068/moneyflow-2023/internal/constants"
)

type AuthJwt struct {
	IJwtHandler
}

func (aj *AuthJwt) GenerateToken(userId int) (string, error) {
	var claimsMap jwt.MapClaims
	claimsMap["userId"] = userId

	token, err := aj.IJwtHandler.GenerateToken(claimsMap)

	if err != nil {
		return "", fmt.Errorf(constants.ErrorFormat, "aj.IJwtHandler.GenerateToken", err)
	}

	return token, nil
}
