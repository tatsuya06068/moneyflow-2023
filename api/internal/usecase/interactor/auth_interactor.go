package interactor

import (
	"context"
	"fmt"
	"log"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/repository"
	"github.com/tatsuya06068/moneyflow-2023/internal/usecase/port"
)

type AuthInteractor struct {
	repository repository.IAuthRepository
}

func NewAuthInteractor(auth repository.IAuthRepository) port.IAuthInputport {
	return &AuthInteractor{
		repository: auth,
	}
}

func (ai AuthInteractor) Signup(ctx context.Context, param entity.SignupRequest) {

	fmt.Printf("%+v", param)
	insertId, err := ai.repository.InsertAuth(param)
	if err != nil {
		log.Println(err)
	}

	log.Println(insertId)
}
