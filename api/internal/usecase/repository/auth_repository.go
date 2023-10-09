package repository

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/entity"
)

type IAuthRepository interface {
	InsertAuth(ctx context.Context, param entity.SignupRequest) (int64, error)
	Select(ctx context.Context, param entity.SigninRequest) (entity.User, bool, error)
}
