package controller

import (
	"fmt"
	"net/http"

	jwtGo "github.com/dgrijalva/jwt-go"
	database "github.com/tatsuya06068/moneyflow-2023/internal/adapter/gateway"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/interactor"
	"github.com/tatsuya06068/moneyflow-2023/pkg/jwt"
)

type BoPController struct {
	interactor entity.IBoPInteractor
}

func NewBoPController(sqlHandler database.ISqlHandler) *BoPController {
	return &BoPController{
		interactor: interactor.BoPInteractor{
			IBoPRepository: database.BoPDBGateway{
				ISqlHandler: sqlHandler,
			},
		},
	}
}

func (bc *BoPController) GetBoPList(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	requestToken := r.FormValue("token")

	token, _ := jwt.VerifyToken(requestToken)
	fmt.Printf("token.Claims %#v\n", token.Claims)

	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
	}

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, token.Signature)
	return
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// fmt.Println(token)

	// bc.interactor.BoPList(ctx, claims.UserId)
}
