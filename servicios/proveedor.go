package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Proveedor struct {
	Nombre   string
	Contacto string
}

// Funcion para agregar un nuevo registro a la tabla Proveedores
func AgregarProveedor(db *sql.DB, proveedor Proveedor) error {
	// Consulta para agregar proveedor
	query := `INSERT INTO Proveedores (Nombre, Contacto) VALUES(?,?)`
	// Ejecuta la consulta
	_, err := db.Exec(query, proveedor.Nombre, proveedor.Contacto)

	if err != nil {
		log.Printf("Error al agregar proveedor: %v", err)
		return err
	}
	fmt.Println("Proveedor registrado correctamente.")
	return nil
}
