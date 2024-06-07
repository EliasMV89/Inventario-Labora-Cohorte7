package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Producto struct {
	Nombre       string
	Categoria    string
	Cantidad     int
	ID_Proveedor int
}

// Agregar un nuevo registro a la tabla productos
func AgregarProducto(db *sql.DB, producto Producto) error {
	// Consulta para agregar un nuevo productos
	query := `INSERT INTO Productos (Nombre, Categoria, Cantidad, ID_Proveedor) VALUES(?,?,?,?)`
	// Ejecuta la consulta
	_, err := db.Exec(query, producto.Nombre, producto.Categoria, producto.Cantidad, producto.ID_Proveedor)

	if err != nil {
		log.Printf("Error al registrar producto: %v", err)
		return err
	}
	fmt.Println("Producto registrado correctamente.")
	return nil
}

// Listar todos los productos disponibles en la tienda
func ListarProductos(db *sql.DB) error {
	// Consulta para listar todos los productos
	query := `SELECT * FROM Productos`
	// Ejecuta la consulta
	rows, err := db.Query(query)

	if err != nil {
		log.Printf("Error al listar los productos: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id, cantidad int
		var nombre, categoria string
		var idProveedor int

		err := rows.Scan(&id, &nombre, &categoria, &cantidad, &idProveedor)

		if err != nil {
			log.Printf("Error al leer fila: %v", err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Categoria: %s, Cantidad: %d, ID del proveedor: %d\n", id, nombre, categoria, cantidad, idProveedor)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Productos listados correctamente.")
	return nil
}

// Buscar producto por nombre, categoria o ID del proveedor
func BuscarProducto(db *sql.DB, buscar string) error {
	// Consulta para buscar producto
	query := `SELECT * FROM Productos WHERE Productos.Nombre LIKE ? OR Productos.Categoria LIKE ? OR Productos.ID_Proveedor = ?`
	// Ejecuta la consulta
	rows, err := db.Query(query, buscar, buscar, buscar)

	if err != nil {
		log.Printf("Error al buscar producto: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var nombre, categoria string
		var id, cantidad, idProveedor int

		err := rows.Scan(&id, &nombre, &categoria, &cantidad, &idProveedor)
		if err != nil {
			log.Printf("Error al leer las filas: %v", err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Categoria: %s, Cantidad: %d, ID del proveedor: %d\n", id, nombre, categoria, cantidad, idProveedor)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Productos listados correctamente.")
	return nil
}

// Obtener la cantidad de un producto antes de registrar una venta
func ObtenerCantidad(db *sql.DB, idProducto int) (int, error) {
	// Consulta para obtener la cantidad
	query := `SELECT Cantidad FROM Productos WHERE ID = ?`
	// Ejecuta la consulta
	var cantidadStock int
	err := db.QueryRow(query, idProducto).Scan(&cantidadStock)

	if err != nil {
		log.Printf("Error al obtener cantidad: %v", err)
		return 0, err
	}
	return cantidadStock, nil
}

// Actualizar stock despues de una venta
func ActualizarStock(db *sql.DB, cantidad, idProducto int) error {
	// Consulta para actualizar stock
	query := `UPDATE Productos SET Cantidad = Cantidad - ? WHERE ID = ?`
	// Ejecuta la consulta
	_, err := db.Exec(query, cantidad, idProducto)

	if err != nil {
		log.Printf("Error al actualizar la cantidad del producto: %v", err)
		return err
	}
	fmt.Println("Cantidad actualizada correctamente.")
	return nil
}
