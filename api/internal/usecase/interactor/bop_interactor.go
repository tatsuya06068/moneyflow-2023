package interactor

import (
	"context"
	"fmt"

	"github.com/tatsuya06068/moneyflow-2023/internal/constants"
	"github.com/tatsuya06068/moneyflow-2023/internal/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/jwt"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/repository"
)

type BoPInteractor struct {
	repository.IBoPRepository
	jwt.IBopJwt
}

func (bpi BoPInteractor) BoPList(ctx context.Context, token string) ([]entity.BoPList, error) {
	userId, getUserIderr := bpi.IBopJwt.GetUserId(token)
	if getUserIderr != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, "bpi.IBopJwt.GetUserId", getUserIderr)
	}

	list, listErr := bpi.GetBoPList(ctx, userId)
	if listErr != nil {
		return list, fmt.Errorf(constants.ErrorFormat, " bpi.GetBoPList", listErr)
	}
	return list, nil
}
