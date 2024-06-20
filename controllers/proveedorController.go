package controllers

import (
	"Inventario/services"
	"Inventario/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ListarProveedores(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	proveedores, err := services.ListarProveedores(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(proveedores)
}

func ModificarProveedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var proveedor services.Proveedor
	if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	err = services.ModificarProveedor(db, id, proveedor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Proveedor modificado correctamente",
	})
}

func EliminarProveedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	err = services.EliminarProveedor(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Proveedor eliminado correctamente",
	})
}
