package interactor

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/repository"
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
	return ai.repository.InsertAuth(param)

}
