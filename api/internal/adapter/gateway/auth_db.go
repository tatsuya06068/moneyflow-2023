package database

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
)

type AuthDBGateway struct {
	ISqlHandler
}

type auth struct {
	userName     string `db:"user_name"`
	hashPassword string `db:"password"`
}

func (ag AuthDBGateway) InsertAuth(ctx context.Context, param entity.SignupRequest) (int64, error) {

	insertAuth := auth{
		userName:     param.UserName,
		hashPassword: fmt.Sprintf("%x", sha256.Sum256([]byte(param.Password))),
	}

	result, err := ag.Execute("INSERT INTO t_users(user_name, password) VALUES(?,?)", insertAuth.userName, insertAuth.hashPassword)

	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return insertId, nil
}

func (ag AuthDBGateway) Select(ctx context.Context, param entity.SigninRequest) (entity.User, bool, error) {

	targetUser := auth{
		userName:     param.UserName,
		hashPassword: fmt.Sprintf("%x", sha256.Sum256([]byte(param.Password))),
	}

	user := entity.User{}

	row, err := ag.Query("SELECT user_id, user_name FROM t_users WHERE user_name = ? AND password = ?", targetUser.userName, targetUser.hashPassword)

	if err != nil {
		return user, false, err
	}

	defer row.Close()

	if row.Next() {
		err := row.Scan(&user.UserId, &user.UserName)
		if err != nil {
			return user, false, err
		}
		return user, true, err
	}

	return user, false, err

}
