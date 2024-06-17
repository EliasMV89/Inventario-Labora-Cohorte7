package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Producto struct {
	Id           string `json:"id"`
	Nombre       string `json:"nombre"`
	Categoria    string `json:"categoria"`
	Cantidad     int    `json:"cantidad"`
	ID_Proveedor int    `json:"id_proveedor"`
}

func AgregarProducto(db *sql.DB, producto Producto) error {
	query := `INSERT INTO Productos (Nombre, Categoria, Cantidad, ID_Proveedor) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, producto.Nombre, producto.Categoria, producto.Cantidad, producto.ID_Proveedor)
	if err != nil {
		log.Printf("Error al registrar producto: %v", err)
		return err
	}
	fmt.Println("Producto registrado correctamente.")
	return nil
}

func ModificarProducto(db *sql.DB, producto Producto) error {
	query := `UPDATE Productos SET Nombre=?, Categoria=?, Cantidad=?, ID_Proveedor=? WHERE Id=?`
	_, err := db.Exec(query, producto.Nombre, producto.Categoria, producto.Cantidad, producto.ID_Proveedor, producto.Id)
	if err != nil {
		log.Printf("Error al modificar producto: %v", err)
		return err
	}
	fmt.Println("Producto modificado correctamente.")
	return nil
}

func EliminarProducto(db *sql.DB, id string) error {
	query := `DELETE FROM Productos WHERE Id=?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error al eliminar producto: %v", err)
		return err
	}
	fmt.Println("Producto eliminado correctamente.")
	return nil
}


func ListarProductos(db *sql.DB) ([]Producto, error) {
	query := `SELECT * FROM Productos`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al listar los productos: %v", err)
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var producto Producto
		if err := rows.Scan(&producto.Id, &producto.Nombre, &producto.Categoria, &producto.Cantidad, &producto.ID_Proveedor); err != nil {
			log.Printf("Error al leer fila: %v", err)
			continue
		}
		productos = append(productos, producto)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return productos, nil
}

func BuscarProducto(db *sql.DB, buscar string) ([]Producto, error) {
	buscar = "%" + buscar + "%"
	query := `SELECT ID, Nombre, Categoria, Cantidad, ID_Proveedor FROM Productos WHERE Nombre LIKE ? OR Categoria LIKE ? OR ID_Proveedor LIKE ?`

	log.Printf("Executing query: %s with value: %s", query, buscar)

	rows, err := db.Query(query, buscar, buscar, buscar)
	if err != nil {
		log.Printf("Error al buscar producto: %v", err)
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var producto Producto
		if err := rows.Scan(&producto.Id, &producto.Nombre, &producto.Categoria, &producto.Cantidad, &producto.ID_Proveedor); err != nil {
			log.Printf("Error al leer fila: %v", err)
			continue
		}
		productos = append(productos, producto)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return productos, nil
}

func ActualizarStock(db *sql.DB, cantidad, idProducto int) error {
	query := `UPDATE Productos SET Cantidad = Cantidad - ? WHERE ID = ?`
	_, err := db.Exec(query, cantidad, idProducto)
	if err != nil {
		log.Printf("Error al actualizar la cantidad del producto: %v", err)
		return err
	}
	fmt.Println("Cantidad actualizada correctamente.")
	return nil
}
