package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	database "github.com/tatsuya06068/moneyflow-2023/internal/adapter/gateway"
	"github.com/tatsuya06068/moneyflow-2023/internal/adapter/jwt"
	"github.com/tatsuya06068/moneyflow-2023/internal/constants"
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

	list, err := bc.interactor.BoPList(ctx, requestToken)
	if err != nil {
		log.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, constants.ErrorText500)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, list)

}
