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

type BoPController struct {
	interactor entity.IBoPInteractor
}

func NewBoPController(sqlHandler database.ISqlHandler, jwtHandler jwt.IJwtHandler) *BoPController {
	return &BoPController{
		interactor: interactor.BoPInteractor{
			IBoPRepository: database.BoPDBGateway{
				ISqlHandler: sqlHandler,
			},
			IBopJwt: &jwt.BopJwt{IJwtHandler: jwtHandler},
		},
	}
}

func (bc *BoPController) GetBoPList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	requestToken := r.FormValue("token")

	// token, _ := jwt.VerifyToken(requestToken)
	// fmt.Printf("token.Claims %#v\n", token.Claims)

	// if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
	// 	fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
	// 	fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
	// }

	// w.WriteHeader(http.StatusInternalServerError)
	// fmt.Fprint(w, token.Signature)
	// return
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// fmt.Println(token)

	list, err := bc.interactor.BoPList(ctx, requestToken)
	fmt.Printf("ERROR %#v", err)
	fmt.Printf("LIST %#v", list)
}
