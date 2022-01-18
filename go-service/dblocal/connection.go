package dblocal

import (
	"github.com/jackc/pgx/v4"
	"context"
	"os"
	"fmt"
)

type Connection struct {
	// conn *pgx.Conn
}

type IConnection interface {}

func Connect() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@db:5432/db?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// defer conn.Close(context.Background())
	return conn
}

// func (c Connection)Transaction() {
// 	tx, err := c.conn.Begin()
// }

const NO_ROWS_IN_RESULT_SET = "no rows in result set"

// func (connection *Connection) Query(query string, field string) {
// 	// query := "select name, weight from widgets where id=$1", 42
// 	err := connection.connec.QueryRow(context.Background(), query).Scan(field)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}
// }