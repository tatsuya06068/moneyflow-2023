package interactor

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/repository"
	"github.com/tatsuya06068/moneyflow-2023/pkg/jwt"
)

type AuthInteractor struct {
	repository.IAuthRepository
}

func (ai AuthInteractor) Signup(ctx context.Context, param entity.SignupRequest) (string, error) {
	userID, err := ai.InsertAuth(ctx, param)

	if err != nil {
		return "", err
	}

	// JWTの生成
	token, err := jwt.GenerateToken(int(userID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (ai AuthInteractor) Signin(ctx context.Context, param entity.SigninRequest) (string, bool, error) {
	user, isUser, err := ai.Select(ctx, param)

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
