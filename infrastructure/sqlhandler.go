package infrastructure

import "database/sql"

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error)
	}
	SqlHandler := new(SqlHandler)
	SqlHandler.Conn = conn
	return SqlHandler
}