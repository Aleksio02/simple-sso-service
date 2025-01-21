package main

import "database/sql"

func main() {
	db, err := sql.Open("sqlite3", "db/sso/sso.db")
	// TODO: создать таблицы для бд (находятся в sql/sso/*.up.sql)
	// TODO: создать downgrade версию скрипта
}
