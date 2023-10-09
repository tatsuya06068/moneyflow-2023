package repository

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/entity"
)

type IBoPRepository interface {
	GetBoPList(ctx context.Context, userId int) ([]entity.BoPList, error)
}
