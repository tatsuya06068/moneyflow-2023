package database

import (
	"github.com/tatsuya06068/moneyflow-2023/repository"
)

type AuthDB struct {
	SqlHandler ISqlHandler
}

func NewAuthDB(sh ISqlHandler) repository.IAuthRepository {
	return &AuthDB{
		SqlHandler: sh,
	}
}

func (a AuthDB) InsertAuth() bool {
	return true
}
