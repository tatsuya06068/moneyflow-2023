package port

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
)

type IAuthInputport interface {
	Signup(ctx context.Context, param entity.SignupRequest)
}

type IAuthOutputPort interface {
	Render(bool)
	RenderError(error)
}
