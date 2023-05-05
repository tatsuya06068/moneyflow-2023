package interactor

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/usecase/port"
)

type AuthInteractor struct {
}

func NewAuthInteractor() port.IAuthInputport {

}

func (ai AuthInteractor) Signup(ctx context.Context, userName string, password string) {

}
