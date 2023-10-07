package controller

import (
	"net/http"

	database "github.com/tatsuya06068/moneyflow-2023/internal/adapter/gateway"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/interactor"
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

	// requestToken := r.FormValue("token")

	// token, _ := jwt.VerifyToken(requestToken)
	// test := token.Claims

	// for _, val := range test.Valid() {
	// 	fmt.Println(val)
	// }
	// w.WriteHeader(http.StatusInternalServerError)
	// fmt.Fprint(w, token.Signature)
	return
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// fmt.Println(token)

	// bc.interactor.BoPList(ctx, claims.UserId)
}
