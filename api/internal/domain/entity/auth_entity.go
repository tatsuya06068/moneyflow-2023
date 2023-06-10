package entity

import "context"

type SignupRequest struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
}

type IAuthInteractor interface {
	Signup(ctx context.Context, param SignupRequest) (int64, error)
}
