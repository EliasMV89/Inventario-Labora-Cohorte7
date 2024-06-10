package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Proveedor struct {
	Nombre   string `json:"nombre"`
	Contacto string `json:"contacto"`
}

func AgregarProveedor(db *sql.DB, proveedor Proveedor) error {
	query := `INSERT INTO Proveedores (Nombre, Contacto) VALUES (?, ?)`
	_, err := db.Exec(query, proveedor.Nombre, proveedor.Contacto)
	if err != nil {
		log.Printf("Error al agregar proveedor: %v", err)
		return err
	}
	fmt.Println("Proveedor registrado correctamente.")
	return nil
}
