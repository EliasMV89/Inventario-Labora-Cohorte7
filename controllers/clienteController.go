package controllers

import (
	"Inventario/services"
	"Inventario/utils"
	"encoding/json"
	"net/http"
)

func AgregarCliente(w http.ResponseWriter, r *http.Request) {
	var cliente services.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	err := services.AgregarCliente(db, cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Cliente agregado correctamente",
	})
}
