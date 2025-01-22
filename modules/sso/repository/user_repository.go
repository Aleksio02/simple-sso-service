package repository

import (
	"context"
	"database/sql"
	"errors"
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
	_, err := ur.db.Exec("INSERT INTO USER(username, password, role) VALUES(?, ?, ?)", username, password, role)
	return err
}

func (ur SQLiteUserRepository) GetUserByUsername(username string) (model.User, error) {
	var result model.User
	row, _ := ur.db.Query("SELECT id, username, password, role FROM USER WHERE username = ?", username)
	defer row.Close()

	if row.Next() {
		row.Scan(&result.Id, &result.Username, &result.Password, &result.Role)
		return result, nil
	}

	return model.User{}, errors.New("user not found")
}

func CreateSQLiteUserRepository() SQLiteUserRepository {
	db, err := sql.Open("sqlite3", "db/sso/sso.db")

	if err != nil {
		panic(err)
	}
	return SQLiteUserRepository{db: db, context: context.Background()}
}
