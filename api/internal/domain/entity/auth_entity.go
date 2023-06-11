package entity

import "context"

type SignupRequest struct {
	UserName string
	Password string
}

type SigninRequest struct {
	UserName string
	Password string
}

type User struct {
	UserId   int64
	UserName string
}

type IAuthInteractor interface {
	Signup(ctx context.Context, param SignupRequest) (int64, error)
	Signin(ctx context.Context, param SigninRequest) (string, bool, error)
}
