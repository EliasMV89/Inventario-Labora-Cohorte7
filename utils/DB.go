/*
package utils

import (

	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

	func InitDB() {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Inventario")
		if err != nil {
			log.Fatal(err)
		}
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
	}

	func GetDB() *sql.DB {
		return db
	}
*/
package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
