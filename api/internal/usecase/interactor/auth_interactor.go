package interactor

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/jwt"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/repository"
)

type AuthInteractor struct {
	repository.IAuthRepository
	jwt.IAuthJwt
}

func (ai *AuthInteractor) Signup(ctx context.Context, param entity.SignupRequest) (string, error) {
	userID, err := ai.IAuthRepository.InsertAuth(ctx, param)

	if err != nil {
		return "", err
	}

	// JWTの生成
	token, err := ai.IAuthJwt.GenerateToken(int(userID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (ai *AuthInteractor) Signin(ctx context.Context, param entity.SigninRequest) (string, bool, error) {
	user, isUser, err := ai.IAuthRepository.Select(ctx, param)

	if err != nil {
		return "", false, err
	}

	if !isUser {
		return "", isUser, err
	}

	// JWTの生成
	token, err := ai.IAuthJwt.GenerateToken(int(user.UserId))
	if err != nil {
		return "", false, err
	}

	return token, true, nil

}
