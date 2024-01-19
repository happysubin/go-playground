package main
 
import (
    "database/sql"
    "log"
 
    _ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}

	defer tx.Rollback() //중간에 에러시 롤백

	_, err = tx.Exec("Insert INTO user(username, age) VALUES (?, ?)", "Kang", 29)

	if err != nil {
		log.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}