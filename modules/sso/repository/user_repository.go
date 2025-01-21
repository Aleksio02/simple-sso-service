package repository

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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

func (ur SQLiteUserRepository) SaveUser(username string, password string, role string) error {
	_, err := ur.db.ExecContext(ur.context, "INSERT INTO USER(username, password, role) VALUES($1, $2, $3)", username, password, role)
	return err
}

func (ur SQLiteUserRepository) GetUserByUsername(username string) (model.User, error) {
	// TODO: implement me
	//var result model.User
	//_, err := ur.db.QueryRow("SELECT id, username, password, role FROM USER WHERE username = $1", username)
	//return err
	//panic("Not implemented")

	return model.User{AuthRequest: model.AuthRequest{Username: "supervisor", Password: "supervisor"}, Role: "USER", Id: 1}, nil

}

func CreateSQLiteUserRepository() SQLiteUserRepository {
	db, err := sql.Open("sqlite3", "db/sso/sso.db")

	if err != nil {
		panic(err)
	}
	return SQLiteUserRepository{db: db, context: context.Background()}
}
