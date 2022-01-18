package schemas

import (
	// "database/sql"
	"fmt"
	"main/dblocal"
	"os"
	"context"
)

type User struct {
	first_name string `json:"first_name"`
	last_name string `json:"last_name"`
	created_at string `json:"created_at"`
}

func CreateUser(first_name string, last_name string) bool {
	sqlStatement := fmt.Sprintf(`
	INSERT INTO users (first_name, last_name)
	VALUES ('%s', '%s')`, first_name, last_name)
	db := dblocal.Connect()
	ctx := context.Background()
	fmt.Println("HIT METHOD")
		tx, _ := db.Begin(ctx)
		defer tx.Rollback(ctx)
		
		_, err := db.Exec(ctx, sqlStatement)
			if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			// os.Exit(1)
			return false
	
		}
		tx.Commit(ctx)
	

	// fmt.Println(err)
	return true
}

func GetUser(id int) *User{
	connection := dblocal.Connect()
	var first_name string
	var last_name string
	var created_at string

	row := connection.QueryRow(context.Background(), "select * from users where id=7") //, id)
	fmt.Println(row)
	err := row.Scan(&first_name, &last_name, &created_at)
	// fmt.Println(err.Error())
	// fmt.Println(sql.ErrNoRows.Error())
	// fmt.Println(sql.ErrNoRows)
	if err.Error() == dblocal.NO_ROWS_IN_RESULT_SET {
		return nil
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return &User{ 
		first_name: first_name,
		last_name: last_name,
		created_at: created_at,
	}
}