package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Cliente struct {
    ID       int    `json:"id"`
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

func ModificarCliente(db *sql.DB, cliente Cliente) error {
    query := `UPDATE Clientes SET Nombre=?, Contacto=? WHERE id=?`
    _, err := db.Exec(query, cliente.Nombre, cliente.Contacto, cliente.ID)
    if err != nil {
        log.Printf("Error al modificar cliente: %v", err)
        return err
    }
    fmt.Println("Cliente modificado correctamente.")
    return nil
}

func EliminarCliente(db *sql.DB, id int) error {
    query := `DELETE FROM Clientes WHERE id=?`
    _, err := db.Exec(query, id)
    if err != nil {
        log.Printf("Error al eliminar cliente: %v", err)
        return err
    }
    fmt.Println("Cliente eliminado correctamente.")
    return nil
}
