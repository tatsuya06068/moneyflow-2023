package controller

import (
	"context"
	"fmt"
	"net/http"

	database "github.com/tatsuya06068/moneyflow-2023/internal/adapter/gateway"
	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/jwt"
	"github.com/tatsuya06068/moneyflow-2023/internal/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/interactor"
)

type AuthController struct {
	interactor entity.IAuthInteractor
}

func NewAuthController(sqlHandler database.ISqlHandler, jwtHandler jwt.IJwtHandler) *AuthController {
	return &AuthController{
		interactor: &interactor.AuthInteractor{
			IAuthRepository: database.AuthDBGateway{ISqlHandler: sqlHandler},
			IAuthJwt:        &jwt.AuthJwt{IJwtHandler: jwtHandler},
		},
	}
}

func (ac *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := entity.SignupRequest{
		UserName: r.PostFormValue("name"),
		Password: r.PostFormValue("password"),
	}
	// バリデーションチェック
	if !validate(param.UserName, param.Password) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "未入力の項目があります。")
		return
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

func (ac *AuthController) Signin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := entity.SigninRequest{
		UserName: r.FormValue("name"),
		Password: r.FormValue("password"),
	}

	// バリデーションチェック
	if !validate(param.UserName, param.Password) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "未入力の項目があります。")
		return
	}

	token, isUser, err := ac.interactor.Signin(ctx, param)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	if !isUser {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "false")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, token)
}

// / バリデーションチェック
func validate(userName string, password string) bool {

	if userName == "" || password == "" {
		return false
	}

	return true
}
