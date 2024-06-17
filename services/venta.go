package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Venta struct {
	Id          int    `json:"id"`
	ID_Producto int    `json:"id_producto"`
	ID_Cliente  int    `json:"id_cliente"`
	Cantidad    int    `json:"cantidad"`
	Total       float64 `json: "total`
	Fecha       string `json:"fecha"`
}

// Función para verificar si hay suficiente stock antes de realizar la venta
func VerificarStock(db *sql.DB, idProducto, cantidad int) (bool, error) {
	var stock int
	err := db.QueryRow("SELECT Cantidad FROM Productos WHERE ID = ?", idProducto).Scan(&stock)
	if err != nil {
		log.Printf("Error al obtener stock del producto: %v", err)
		return false, err
	}

	return stock >= cantidad, nil
}

// Función para agregar una venta y actualizar la cantidad de productos en la tabla Producto
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
	err = db.QueryRow("SELECT Precio FROM Productos WHERE ID = ?", venta.ID_Producto).Scan(&precioProducto)
	if err != nil {
		log.Printf("Error al obtener precio del producto: %v", err)
		return err
	}

	venta.Total = precioProducto * float64(venta.Cantidad) // Calcula el total

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

	// Actualizar cantidad de productos
	_, err = tx.Exec("UPDATE Productos SET Cantidad = Cantidad - ? WHERE ID = ?", venta.Cantidad, venta.ID_Producto)
	if err != nil {
		log.Printf("Error al actualizar cantidad de productos: %v", err)
		return err
	}

	// Registrar la venta
	query := `INSERT INTO Ventas (ID_Producto, ID_Cliente, Cantidad, Total, Fecha) VALUES (?, ?, ?, ?, ?)` 
	_, err = tx.Exec(query, venta.ID_Producto, venta.ID_Cliente, venta.Cantidad, venta.Total, venta.Fecha)
	if err != nil {
		log.Printf("Error al registrar venta: %v", err)
		return err
	}
	fmt.Println("Venta registrada correctamente.")
	return nil
}

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
		if err := rows.Scan(&idProducto, &totalVendido); err != nil {
			log.Printf("Error al leer fila: %v", err)
			return nil, err // Devuelve el error en lugar de continuar
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


