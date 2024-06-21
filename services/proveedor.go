package services

import (
	"database/sql" 
	"fmt"          
	"log"          
)

// Estructura Proveedor define los campos que se corresponden con las columnas de la tabla Proveedores
type Proveedor struct {
	ID       int    `json:"id"`      // Identificador único del proveedor
	Nombre   string `json:"nombre"`  // Nombre del proveedor
	Contacto string `json:"contacto"`// Información de contacto del proveedor
}

// AgregarProveedor inserta un nuevo proveedor en la base de datos
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

// ListarProveedores recupera todos los proveedores de la base de datos
func ListarProveedores(db *sql.DB) ([]Proveedor, error) {
	query := `SELECT ID, Nombre, Contacto FROM Proveedores`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al obtener la lista de proveedores: %v", err)
		return nil, err
	}
	defer rows.Close()

	var proveedores []Proveedor
	for rows.Next() {
		var proveedor Proveedor
		// Escanea cada fila y asigna los valores a la estructura Proveedor
		err := rows.Scan(&proveedor.ID, &proveedor.Nombre, &proveedor.Contacto)
		if err != nil {
			log.Printf("Error al escanear fila de proveedor: %v", err)
			return nil, err
		}
		proveedores = append(proveedores, proveedor)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas de proveedores: %v", err)
		return nil, err
	}

	return proveedores, nil
}

// ModificarProveedor actualiza los datos de un proveedor existente en la base de datos
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

// EliminarProveedor elimina un proveedor de la base de datos por su ID
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


