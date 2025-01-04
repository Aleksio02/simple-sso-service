package repository

import (
	"context"
	"database/sql"
	"simple-sso-service/modules/sso/model"
)

type UserRepository interface {
	SaveUser(username string, password string, role string) error
	GetUserByUsername(username string) (model.User, error)
}

type SQLiteUserRepository struct {
	db      *sql.DB
	context context.Context
}

func (ur *SQLiteUserRepository) SaveUser(username string, password string, role string) error {
	_, err := ur.db.ExecContext(ur.context, "INSERT INTO USER(username, password, role) VALUES($1, $2, $3)", username, password, role)
	return err
}

func (ur *SQLiteUserRepository) GetUserByUsername(username string) (model.User, error) {
	// TODO: implement me
	//var result model.User
	//_, err := ur.db.QueryRow("SELECT id, username, password, role FROM USER WHERE username = $1", username)
	//return err
	panic("Not implemented")
}

func CreateSQLiteUserRepository() *SQLiteUserRepository {
	// TODO: aleksioi: проверить название драйвера
	// TODO: aleksioi: вставить строку подключения
	// TODO: aleksioi: придумать способ доставки данных для подключения
	db, err := sql.Open("SQLite", "")
	if err != nil {
		panic(err)
	}
	return &SQLiteUserRepository{db: db, context: context.Background()}
}
