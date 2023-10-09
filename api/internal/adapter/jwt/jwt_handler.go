package jwt

import "github.com/dgrijalva/jwt-go"

type IJwtHandler interface {
	GenerateToken(claimsMap jwt.MapClaims) (string, error)
	VerifyToken(tokenString string) (IClimeDriver, error)
}

type IClimeDriver interface {
	GetPayload() (jwt.MapClaims, error)
}
