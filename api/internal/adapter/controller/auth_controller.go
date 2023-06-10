package controller

import (
	"context"
	"fmt"
	"net/http"

	database "github.com/tatsuya06068/moneyflow-2023/internal/adapter/gateway"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/interactor"
)

type AuthController struct {
	interactor entity.IAuthInteractor
}

func NewAuthController(sqlHandler database.ISqlHandler) *AuthController {
	return &AuthController{
		interactor: interactor.NewAuthInteractor(database.NewAuthDB(sqlHandler)),
	}
}

func (ac AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := entity.SignupRequest{
		UserName: r.PostFormValue("name"),
		Password: r.PostFormValue("password"),
	}
	// バリデーションチェック
	if param.UserName == "" || param.Password == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "未入力の項目があります。")
	}

	id, err := ac.interactor.Signup(ctx, param)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, id)
}
