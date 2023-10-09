package jwt

type IBopJwt interface {
	GetUserId(token string) (int, error)
}
