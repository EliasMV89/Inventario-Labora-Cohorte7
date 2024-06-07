package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarBaseDeDatos() (*sql.DB, error) {
	// Establece la conexión con la base de datos
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Inventario")
	if err != nil {
		log.Printf("Error al abrir la conexión: %v", err)
		return nil, err
	}

	// Verifica la conexión
	err = db.Ping()
	if err != nil {
		log.Printf("Error al establecer la conexión: %v", err)
		return nil, err
	}

	fmt.Println("Conexión establecida con éxito.")
	return db, nil
}
func main() {
	// Intenta conectar con la base de datos
	db, err := ConectarBaseDeDatos()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
