package controllers

import (
	"Inventario/services" 
	"Inventario/utils"    
	"encoding/json"       
	"net/http"            
	"strconv"             

	"github.com/gorilla/mux" 
)

// AgregarProveedor maneja la solicitud para agregar un nuevo proveedor
func AgregarProveedor(w http.ResponseWriter, r *http.Request) {
	var proveedor services.Proveedor
	// Decodifica el cuerpo de la solicitud JSON en la estructura Proveedor
	if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la función AgregarProveedor del paquete services para agregar el proveedor a la base de datos
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

// ListarProveedores maneja la solicitud para listar todos los proveedores
func ListarProveedores(w http.ResponseWriter, r *http.Request) {
	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la función ListarProveedores del paquete services para obtener la lista de proveedores
	proveedores, err := services.ListarProveedores(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// Codifica la lista de proveedores en formato JSON y la envía en la respuesta
	json.NewEncoder(w).Encode(proveedores)
}

// ModificarProveedor maneja la solicitud para modificar un proveedor existente
func ModificarProveedor(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID del proveedor de la URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var proveedor services.Proveedor
	// Decodifica el cuerpo de la solicitud JSON en la estructura Proveedor
	if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la función ModificarProveedor del paquete services para modificar el proveedor en la base de datos
	err = services.ModificarProveedor(db, id, proveedor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Proveedor modificado correctamente",
	})
}

// EliminarProveedor maneja la solicitud para eliminar un proveedor por su ID
func EliminarProveedor(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID del proveedor de la URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la función EliminarProveedor del paquete services para eliminar el proveedor de la base de datos
	err = services.EliminarProveedor(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Proveedor eliminado correctamente",
	})
}
