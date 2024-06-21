package services

import (
	"database/sql" 
	"fmt"          
	"log"          
)

// Estructura Cliente define los campos que se corresponden con las columnas de la tabla Clientes
type Cliente struct {
    ID       int    `json:"id"`      // Identificador único del cliente
	Nombre   string `json:"nombre"`  // Nombre del cliente
	Contacto string `json:"contacto"`// Información de contacto del cliente
}

// AgregarCliente inserta un nuevo cliente en la base de datos
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

// ListarClientes recupera todos los clientes de la base de datos
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

// ModificarCliente actualiza los datos de un cliente existente en la base de datos
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

// EliminarCliente elimina un cliente de la base de datos
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


