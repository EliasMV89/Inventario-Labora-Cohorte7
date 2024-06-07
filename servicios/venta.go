package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Venta struct {
	ID_Producto int
	ID_Cliente  int
	Cantidad    int
	Fecha       string
}

// Agregar un registro a la tabla Ventas
func AgregarVenta(db *sql.DB, venta Venta) error {
	// Consulta para agregar registro
	query := `INSERT INTO Ventas (ID_Producto, ID_Cliente, Cantidad, Fecha) VALUES (?,?,?,?)`
	// Ejecuta la consulta
	_, err := db.Exec(query, venta.ID_Producto, venta.ID_Cliente, venta.Cantidad, venta.Fecha)

	if err != nil {
		log.Printf("Error al registrar venta: %v", err)
	}
	fmt.Println("Venta registrada correctamente.")
	return nil
}

// Buscar ventas realizadas en un dia especifico
func BuscarVentaPorFecha(db *sql.DB, fechaBusqueda string) error {
	// Consulta para buscar ventas
	query := `SELECT * FROM Ventas WHERE Fecha = ?`
	// Ejecuta la consulta
	rows, err := db.Query(query, fechaBusqueda)

	if err != nil {
		log.Printf("Error al buscar venta: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id, idProducto, idCliente, cantidad int
		var fecha string

		err := rows.Scan(&id, &idProducto, &cantidad, &fecha, &idCliente)

		if err != nil {
			log.Printf("Error al leer fila: %v", err)
		}
		fmt.Printf("ID: %d, ID del producto: %d, ID del cliente: %d, Cantidad: %d, Fecha: %v\n", id, idProducto, idCliente, cantidad, fecha)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Productos listados correctamente")
	return nil
}

// Generar informe de productos mas vendidos en un periodo de tiempo
func GenerarInforme(db *sql.DB, fechaInicio, fechaFin string) error {
	// Consulta para generar informe
	query := `SELECT ID_Producto, SUM(Cantidad) as Total_Vendido FROM Ventas WHERE Fecha BETWEEN ? AND ? GROUP BY ID_Producto ORDER BY Total_Vendido DESC`
	// Ejecuta la consulta
	rows, err := db.Query(query, fechaInicio, fechaFin)

	if err != nil {
		log.Printf("Error al generar informe: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var idProducto, totalVendido int

		err := rows.Scan(&idProducto, &totalVendido)

		if err != nil {
			log.Printf("Error al leer fila: %v", err)
		}
		fmt.Printf("ID del producto: %d, Total vendido: %d\n", idProducto, totalVendido)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Informe generado correctamente")
	return nil
}
