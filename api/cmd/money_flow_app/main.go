package main

import (
	"log"
	"net/http"

	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/controller"
	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/database"
	dbDriver "github.com/tatsuya06068/moneyflow-2023/internal/driver"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/interactor"
)

func main() {
	mux := http.NewServeMux()

	auth := controller.AuthController{
		RepoFactory:  database.NewAuthDB,
		InputFactory: interactor.NewAuthInteractor,
		SqlHandler:   dbDriver.NewSqlHandler(),
	}

	// signup
	mux.HandleFunc("/auth/signup/", auth.Signup)

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", mux))
}
