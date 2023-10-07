package interactor

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/repository"
)

type BoPInteractor struct {
	repository.IBoPRepository
}

func (bpi BoPInteractor) BoPList(ctx context.Context, userId int) ([]entity.BoPList, error) {
	return bpi.GetBoPList(ctx, userId)
}
