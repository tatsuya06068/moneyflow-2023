package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/controller"
	"github.com/tatsuya06068/moneyflow-2023/internal/constants"
	"github.com/tatsuya06068/moneyflow-2023/internal/driver/database"
	"github.com/tatsuya06068/moneyflow-2023/internal/driver/jwt"
)

func main() {
	mux := http.NewServeMux()
	// create database driver
	auth := controller.NewAuthController(database.NewSqlHandler())

	// create jwt driver
	jwtDriver := jwt.NewJwtHandler(os.Getenv(constants.JwtSecretKey))

	// signup
	mux.HandleFunc("/auth/signup/", auth.Signup)
	// signin
	mux.HandleFunc("/auth/signin/", auth.Signin)

	bop := controller.NewBoPController(database.NewSqlHandler())
	// bopList
	mux.HandleFunc("/bop/list/", bop.GetBoPList)
	//taskDetail

	//taskDelete

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", mux))
}
