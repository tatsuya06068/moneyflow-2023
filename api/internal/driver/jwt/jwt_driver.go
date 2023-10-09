package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	iJwt "github.com/tatsuya06068/moneyflow-2023/internal/adapter/jwt"
)

type JwtHandler struct {
	secretKey []byte
}

func NewJwtHandler(secretKey string) *JwtHandler {
	return &JwtHandler{
		secretKey: []byte(secretKey),
	}
}

// token作成
func (jh *JwtHandler) GenerateToken(claimsMap jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	for k, v := range claimsMap {
		claims[k] = v
	}

	// JWTの著名
	signedToken, err := token.SignedString(jh.secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// tokenデコード
func (jh *JwtHandler) VerifyToken(tokenString string) (iJwt.IClimeDriver, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jh.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claim := new(ClimeDriver)
	claim.token = token
	return claim, nil
}

type ClimeDriver struct {
	token *jwt.Token
}

// payload取り出し
func (cd *ClimeDriver) GetPayload() (jwt.MapClaims, error) {
	fmt.Printf("user_id: %#v\n", cd.token)
	if claims, ok := cd.token.Claims.(jwt.MapClaims); ok && cd.token.Valid {
		return claims, nil
	}
	return nil, errors.New("Fail get payload")

}
