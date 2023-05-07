package presenter

import "github.com/tatsuya06068/moneyflow-2023/internal/usecase/port"

type AuthPresenter struct {
}

func NewAuthPresenter() port.IAuthOutputPort {
	return nil
}

func (ap AuthPresenter) Render(bool) {

}

func (ap AuthPresenter) RenderError(error) {

}
