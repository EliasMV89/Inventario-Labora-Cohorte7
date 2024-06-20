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

func ListarClientes(db *sql.DB) ([]Cliente, error) {
	query := `SELECT id, Nombre, Contacto FROM Clientes`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al obtener la lista de clientes: %v", err)
		return nil, err
	}
	defer rows.Close()

	var clientes []Cliente
	for rows.Next() {
		var cliente Cliente
		err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.Contacto)
		if err != nil {
			log.Printf("Error al escanear fila de cliente: %v", err)
			return nil, err
		}
		clientes = append(clientes, cliente)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas de clientes: %v", err)
		return nil, err
	}

	return clientes, nil
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
