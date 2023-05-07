package repository

import "github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"

type IAuthRepository interface {
	InsertAuth(param entity.SignupRequest) (int64, error)
}
