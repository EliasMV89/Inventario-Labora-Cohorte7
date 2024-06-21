package controllers

import (
	"Inventario/services" 
	"Inventario/utils"    
	"encoding/json"       
	"net/http"            
)

// AgregarVenta maneja la solicitud para agregar una nueva venta
func AgregarVenta(w http.ResponseWriter, r *http.Request) {
	var venta services.Venta
	// Decodifica el cuerpo de la solicitud JSON en la estructura Venta
	if err := json.NewDecoder(r.Body).Decode(&venta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtiene la conexión de la base de datos
	db := utils.GetDB()
	// Llama a la función AgregarVenta del paquete services para agregar la venta a la base de datos
	err := services.AgregarVenta(db, venta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Venta registrada correctamente",
	})
}

// ListarVentas maneja la solicitud para listar todas las ventas
func ListarVentas(w http.ResponseWriter, r *http.Request) {
	// Obtiene la conexión de la base de datos
	db := utils.GetDB()
	// Llama a la función ListarVentas del paquete services para obtener la lista de ventas
	ventas, err := services.ListarVentas(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// Codifica la lista de ventas en formato JSON y la envía en la respuesta
	json.NewEncoder(w).Encode(ventas)
}

// BuscarVentaPorFecha maneja la solicitud para buscar ventas por fecha específica
func BuscarVentaPorFecha(w http.ResponseWriter, r *http.Request) {
	// Obtiene el parámetro de fecha de la URL
	fecha := r.URL.Query().Get("fecha")
	// Obtiene la conexión de la base de datos
	db := utils.GetDB()
	// Llama a la función BuscarVentaPorFecha del paquete services para buscar ventas en la base de datos
	ventas, err := services.BuscarVentaPorFecha(db, fecha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Codifica la lista de ventas encontradas en formato JSON y la envía en la respuesta
	json.NewEncoder(w).Encode(ventas)
}

// GenerarInforme maneja la solicitud para generar un informe de ventas entre dos fechas
func GenerarInforme(w http.ResponseWriter, r *http.Request) {
	// Obtiene los parámetros de fechaInicio y fechaFin de la URL
	fechaInicio := r.URL.Query().Get("fechaInicio")
	fechaFin := r.URL.Query().Get("fechaFin")
	// Obtiene la conexión de la base de datos
	db := utils.GetDB()
	// Llama a la función GenerarInforme del paquete services para generar el informe de ventas
	informe, err := services.GenerarInforme(db, fechaInicio, fechaFin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Codifica el informe generado en formato JSON y lo envía en la respuesta
	json.NewEncoder(w).Encode(informe)
}
