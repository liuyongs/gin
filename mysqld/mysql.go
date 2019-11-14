package mysqld

import (
	"database/sql"
	_"log"
)

var Db *sql.DB

func DbConnect(){
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gogin?parseTime=true")
	if err != nil{
		//log.Fatalln(err)
	}
	//defer Db.Close()


	Db.SetMaxIdleConns(20)
	Db.SetMaxOpenConns(20)

	if err := Db.Ping(); err != nil{
		//log.Fatalln(err)
	}
}
