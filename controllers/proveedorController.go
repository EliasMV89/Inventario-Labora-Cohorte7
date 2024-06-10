package controllers

import (
	"Inventario/services"
	"Inventario/utils"
	"encoding/json"
	"net/http"
)

func AgregarProveedor(w http.ResponseWriter, r *http.Request) {
	var proveedor services.Proveedor
	if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	err := services.AgregarProveedor(db, proveedor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Proveedor agregado correctamente",
	})
}
