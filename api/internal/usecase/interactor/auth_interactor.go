package interactor

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/repository"
	"github.com/tatsuya06068/moneyflow-2023/pkg/jwt"
)

type AuthInteractor struct {
	repository repository.IAuthRepository
}

func NewAuthInteractor(authRepo repository.IAuthRepository) entity.IAuthInteractor {
	return &AuthInteractor{
		repository: authRepo,
	}
}

func (ai AuthInteractor) Signup(ctx context.Context, param entity.SignupRequest) (int64, error) {
	return ai.repository.InsertAuth(ctx, param)
}

func (ai AuthInteractor) Signin(ctx context.Context, param entity.SigninRequest) (string, bool, error) {
	user, isUser, err := ai.repository.Select(ctx, param)

	if err != nil {
		return "", false, err
	}

	if !isUser {
		return "", isUser, err
	}

	// JWTの生成
	token, err := jwt.GenerateToken(int(user.UserId))
	if err != nil {
		return "", false, err
	}

	return token, true, nil

}
