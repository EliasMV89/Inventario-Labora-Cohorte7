package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Cliente struct {
	Nombre   string `json:"nombre"`
	Contacto string `json:"contacto"`
}

func AgregarCliente(db *sql.DB, cliente Cliente) error {
	query := `INSERT INTO Clientes (Nombre, Contacto) VALUES (?, ?)`
	_, err := db.Exec(query, cliente.Nombre, cliente.Contacto)
	if err != nil {
		log.Printf("Error al registrar cliente: %v", err)
		return err
	}
	fmt.Println("Cliente registrado correctamente.")
	return nil
}
