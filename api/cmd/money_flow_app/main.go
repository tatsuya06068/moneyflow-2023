package main

import (
	"log"
	"net/http"

	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/controller"
	dbDriver "github.com/tatsuya06068/moneyflow-2023/internal/driver"
)

func main() {
	mux := http.NewServeMux()
	auth := controller.NewAuthController(dbDriver.NewSqlHandler())

	// signup
	mux.HandleFunc("/auth/signup/", auth.Signup)
	// signin
	mux.HandleFunc("/auth/signin/", auth.Signin)

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", mux))
}
