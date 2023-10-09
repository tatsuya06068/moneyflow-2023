package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/controller"
	"github.com/tatsuya06068/moneyflow-2023/internal/constants"
	"github.com/tatsuya06068/moneyflow-2023/internal/driver/database"
	"github.com/tatsuya06068/moneyflow-2023/internal/driver/jwt"
	"github.com/tatsuya06068/moneyflow-2023/internal/entity"
)

func main() {
	mux := http.NewServeMux()

	baseDB := entity.BaseDbInfo{
		HostName: os.Getenv(constants.DbHost),
		User:     os.Getenv(constants.DbUser),
		Password: os.Getenv(constants.DbPassword),
	}

	// create database driver
	dbDriver := database.NewSqlHandler(baseDB)
	// create jwt driver
	jwtDriver := jwt.NewJwtHandler(os.Getenv(constants.JwtSecretKey))

	auth := controller.NewAuthController(dbDriver, jwtDriver)

	// signup
	mux.HandleFunc("/auth/signup/", auth.Signup)
	// signin
	mux.HandleFunc("/auth/signin/", auth.Signin)

	bop := controller.NewBoPController(dbDriver, jwtDriver)
	// bopList
	mux.HandleFunc("/bop/list/", bop.GetBoPList)
	//taskDetail

	//taskDelete

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", mux))
}
