package controller

import "github.com/tatsuya06068/moneyflow-2023/internal/adapter/database"

type AuthController struct {
	InputFactory
	RepoFactory 
	sqlhandler database.ISqlHandler
}

func (ac AuthController) Signup() {

}
