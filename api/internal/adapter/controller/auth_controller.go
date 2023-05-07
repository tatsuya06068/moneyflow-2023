package controller

import (
	"context"
	"net/http"

	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/database"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/repository"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/port"
)

type AuthController struct {
	// database
	RepoFactory func(sqlhandler database.ISqlHandler) repository.IAuthRepository

	// iteractor
	InputFactory func(auth repository.IAuthRepository) port.IAuthInputport

	// database.NewAuthDatabase
	SqlHandler database.ISqlHandler
}

func (ac AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	param := entity.SignupRequest{
		UserName: r.PostFormValue("name"),
		Password: r.PostFormValue("password"),
	}

	repository := ac.RepoFactory(ac.SqlHandler)
	inputPort := ac.InputFactory(repository)
	inputPort.Signup(ctx, param)
}
