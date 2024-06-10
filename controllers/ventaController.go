package controllers

import (
	"Inventario/services"
	"Inventario/utils"
	"encoding/json"
	"net/http"
)

func AgregarVenta(w http.ResponseWriter, r *http.Request) {
	var venta services.Venta
	if err := json.NewDecoder(r.Body).Decode(&venta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
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

func BuscarVentaPorFecha(w http.ResponseWriter, r *http.Request) {
	fecha := r.URL.Query().Get("fecha")
	db := utils.GetDB()
	ventas, err := services.BuscarVentaPorFecha(db, fecha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ventas)
}

func GenerarInforme(w http.ResponseWriter, r *http.Request) {
	fechaInicio := r.URL.Query().Get("fechaInicio")
	fechaFin := r.URL.Query().Get("fechaFin")
	db := utils.GetDB()
	informe, err := services.GenerarInforme(db, fechaInicio, fechaFin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(informe)
}
