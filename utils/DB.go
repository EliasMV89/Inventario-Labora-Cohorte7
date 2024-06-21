package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Variable global para mantener la conexion a la base de datos
var db *sql.DB

// InitDB inicializa la conexion a la base de datos MySQL
func InitDB() {
	var err error
	// Abre la conexión a la base de datos con el usuario root y la base de datos Inventario
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Inventario")
	if err != nil {
		// Si hay un error al abrir la conexión, registra el error y detiene el programa
		log.Fatal(err)
	}
	// Verifica que la conexión a la base de datos sea exitosa
	if err = db.Ping(); err != nil {
		// Si hay un error al hacer ping a la base de datos, registra el error y detiene el programa
		log.Fatal(err)
	}
}

// GetDB devuelve la conexion de base de datos global
func GetDB() *sql.DB {
	return db
}
