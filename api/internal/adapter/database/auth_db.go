package database

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
	"github.com/tatsuya06068/moneyflow-2023/internal/domain/repository"
)

type AuthDB struct {
	SqlHandler ISqlHandler
}

func NewAuthDB(sh ISqlHandler) repository.IAuthRepository {
	return &AuthDB{
		SqlHandler: sh,
	}
}

type auth struct {
	userName     string `db:"user_name"`
	hashPassword string `db:"password"`
}

func (a AuthDB) InsertAuth(param entity.SignupRequest) (int64, error) {

	insertAuth := auth{
		userName:     param.UserName,
		hashPassword: fmt.Sprintf("%x", sha256.Sum256([]byte(param.Password))),
	}

	log.Printf("%+v", insertAuth)

	result, err := a.SqlHandler.Execute("INSERT INTO t_users(user_name, password) VALUES(?,?)", insertAuth.userName, insertAuth.hashPassword)

	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return insertId, nil
}
