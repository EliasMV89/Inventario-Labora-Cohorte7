package services

import (
	"database/sql" 
	"fmt"          
	"log"          
)

// Estructura Venta define los campos que se corresponden con las columnas de la tabla Ventas
type Venta struct {
	Id          int     `json:"id"`           // Identificador único de la venta
	ID_Producto int     `json:"id_producto"`  // Identificador del producto vendido
	ID_Cliente  int     `json:"id_cliente"`   // Identificador del cliente
	Cantidad    int     `json:"cantidad"`     // Cantidad de producto vendido
	Total       float64 `json:"total"`        // Total de la venta
	Fecha       string  `json:"fecha"`        // Fecha de la venta
}

// VerificarStock verifica si hay suficiente stock de un producto antes de realizar una venta
func VerificarStock(db *sql.DB, idProducto, cantidad int) (bool, error) {
	var stock int
	// Consulta la cantidad de stock del producto
	err := db.QueryRow("SELECT Cantidad FROM Productos WHERE ID = ?", idProducto).Scan(&stock)
	if err != nil {
		log.Printf("Error al obtener stock del producto: %v", err)
		return false, err
	}

	return stock >= cantidad, nil
}

// AgregarVenta inserta una nueva venta en la base de datos y actualiza el stock del producto
func AgregarVenta(db *sql.DB, venta Venta) error {
	enStock, err := VerificarStock(db, venta.ID_Producto, venta.Cantidad)
	if err != nil {
		return err
	}

	if !enStock {
		fmt.Println("No hay suficiente stock para realizar la venta.")
		return nil
	}

	var precioProducto float64
	// Consulta el precio del producto
	err = db.QueryRow("SELECT Precio FROM Productos WHERE ID = ?", venta.ID_Producto).Scan(&precioProducto)
	if err != nil {
		log.Printf("Error al obtener precio del producto: %v", err)
		return err
	}

	// Calcula el total de la venta
	venta.Total = precioProducto * float64(venta.Cantidad)

	// Inicia una transacción
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error al iniciar transacción: %v", err)
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Actualiza la cantidad de productos en stock
	_, err = tx.Exec("UPDATE Productos SET Cantidad = Cantidad - ? WHERE ID = ?", venta.Cantidad, venta.ID_Producto)
	if err != nil {
		log.Printf("Error al actualizar cantidad de productos: %v", err)
		return err
	}

	// Inserta la nueva venta en la base de datos
	query := `INSERT INTO Ventas (ID_Producto, ID_Cliente, Cantidad, Total, Fecha) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, venta.ID_Producto, venta.ID_Cliente, venta.Cantidad, venta.Total, venta.Fecha)
	if err != nil {
		log.Printf("Error al registrar venta: %v", err)
		return err
	}
	fmt.Println("Venta registrada correctamente.")
	return nil
}

// ListarVentas recupera todas las ventas de la base de datos
func ListarVentas(db *sql.DB) ([]Venta, error) {
	query := `SELECT ID, ID_Producto, ID_Cliente, Cantidad, Total, Fecha FROM Ventas`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al obtener la lista de ventas: %v", err)
		return nil, err
	}
	defer rows.Close()

	var ventas []Venta
	for rows.Next() {
		var venta Venta
		// Escanea cada fila y asigna los valores a la estructura Venta
		err := rows.Scan(&venta.Id, &venta.ID_Producto, &venta.ID_Cliente, &venta.Cantidad, &venta.Total, &venta.Fecha)
		if err != nil {
			log.Printf("Error al escanear fila de venta: %v", err)
			return nil, err
		}
		ventas = append(ventas, venta)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas de ventas: %v", err)
		return nil, err
	}

	return ventas, nil
}

// BuscarVentaPorFecha busca ventas por una fecha específica en la base de datos
func BuscarVentaPorFecha(db *sql.DB, fechaBusqueda string) ([]Venta, error) {
	query := `SELECT ID_Producto, ID_Cliente, Cantidad, Total, Fecha FROM Ventas WHERE Fecha = ?`
	rows, err := db.Query(query, fechaBusqueda)
	if err != nil {
		log.Printf("Error al buscar venta: %v", err)
		return nil, err
	}
	defer rows.Close()

	var ventas []Venta
	for rows.Next() {
		var venta Venta
		// Escanea cada fila y asigna los valores a la estructura Venta
		if err := rows.Scan(&venta.ID_Producto, &venta.ID_Cliente, &venta.Cantidad, &venta.Total, &venta.Fecha); err != nil {
			log.Printf("Error al leer fila: %v", err)
			continue
		}
		ventas = append(ventas, venta)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return ventas, nil
}

// GenerarInforme genera un informe de ventas para un rango de fechas
func GenerarInforme(db *sql.DB, fechaInicio, fechaFin string) ([]map[string]interface{}, error) {
	query := `SELECT ID_Producto, SUM(Cantidad) as Total_Vendido FROM Ventas WHERE Fecha BETWEEN ? AND ? GROUP BY ID_Producto ORDER BY Total_Vendido DESC`
	rows, err := db.Query(query, fechaInicio, fechaFin)
	if err != nil {
		log.Printf("Error al generar informe: %v", err)
		return nil, err
	}
	defer rows.Close()

	var informe []map[string]interface{}
	for rows.Next() {
		var idProducto, totalVendido int
		// Escanea cada fila y asigna los valores a un mapa
		if err := rows.Scan(&idProducto, &totalVendido); err != nil {
			log.Printf("Error al leer fila: %v", err)
			return nil, err
		}
		informe = append(informe, map[string]interface{}{
			"id_producto":   idProducto,
			"total_vendido": totalVendido,
		})
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return informe, nil
}



