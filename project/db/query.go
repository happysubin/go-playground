package main

import (
    "database/sql" 
    _"github.com/go-sql-driver/mysql" //를 사용하면 해당 패키지를 가져오긴 하지만 패키지의 심볼을 현재 스코프로 노출시키지 않습니다.
	//이 패턴은 특정 패키지의 초기화 코드를 실행하기 위해 사용되며, 패키지의 기능을 직접 사용하지 않을 때 유용합니다.
    "log"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close() //Golang의 defer 키워드는 함수 호출을 현재 함수가 종료되기 직전까지 지연시키는 데 사용됩니다. 

	var name string
	err = db.QueryRow("SELECT username FROM user WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	var id int
	var username string
	rows, err := db.Query("SELECT id, username FROM user where id >= ?", 1)

	if(err != nil) {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username)
	}

	
}