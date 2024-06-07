package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Cliente struct {
	Nombre   string
	Contacto string
}

// Agregar un registro a la tabla Cliente
func AgregarCliente(db *sql.DB, cliente Cliente) error {
	// Consulta para agregar registro
	query := `INSERT INTO Clientes (Nombre, Contacto) VALUES(?,?)`
	// Ejecuta la consulta
	_, err := db.Exec(query, cliente.Nombre, cliente.Contacto)

	if err != nil {
		log.Printf("Error al registrar cliente: %v", err)
		return err
	}
	fmt.Println("Cliente registrado correctamente.")
	return nil
}
