package controllers

import (
	"Inventario/services"
	"Inventario/utils"
	"encoding/json"
	"net/http"
)

func AgregarProducto(w http.ResponseWriter, r *http.Request) {
	var producto services.Producto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	err := services.AgregarProducto(db, producto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Producto registrado correctamente",
	})
}

func ListarProductos(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	productos, err := services.ListarProductos(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(productos)
}

func BuscarProducto(w http.ResponseWriter, r *http.Request) {
	buscar := r.URL.Query().Get("buscar")
	db := utils.GetDB()
	productos, err := services.BuscarProducto(db, buscar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(productos)
}
