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

func ModificarProveedor(db *sql.DB, id int, proveedor Proveedor) error {
	query := `UPDATE Proveedores SET Nombre=?, Contacto=? WHERE ID=?`
	_, err := db.Exec(query, proveedor.Nombre, proveedor.Contacto, id)
	if err != nil {
		log.Printf("Error al modificar proveedor: %v", err)
		return err
	}
	fmt.Println("Proveedor modificado correctamente.")
	return nil
}

func EliminarProveedor(db *sql.DB, id int) error {
	query := `DELETE FROM Proveedores WHERE ID=?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error al eliminar proveedor: %v", err)
		return err
	}
	fmt.Println("Proveedor eliminado correctamente.")
	return nil
}

