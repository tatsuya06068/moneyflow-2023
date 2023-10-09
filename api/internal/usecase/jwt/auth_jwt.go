package jwt

type IAuthJwt interface {
	GenerateToken(userId int) (string, error)
}
