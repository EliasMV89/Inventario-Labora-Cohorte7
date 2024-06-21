package controllers

import (
	"Inventario/services"
	"Inventario/utils"    
	"encoding/json"       
	"net/http"            

	"github.com/gorilla/mux" 
)

// AgregarProducto maneja la solicitud para agregar un nuevo producto
func AgregarProducto(w http.ResponseWriter, r *http.Request) {
	var producto services.Producto
	// Decodifica el cuerpo de la solicitud JSON en la estructura Producto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la funcion AgregarProducto del paquete services para agregar el producto a la base de datos
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

// ModificarProducto maneja la solicitud para modificar un producto existente
func ModificarProducto(w http.ResponseWriter, r *http.Request) {
	var producto services.Producto
	// Decodifica el cuerpo de la solicitud JSON en la estructura Producto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la funcion ModificarProducto del paquete services para modificar el producto en la base de datos
	err := services.ModificarProducto(db, producto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Producto modificado correctamente",
	})
}

// EliminarProducto maneja la solicitud para eliminar un producto por su ID
func EliminarProducto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la funcion EliminarProducto del paquete services para eliminar el producto de la base de datos
	err := services.EliminarProducto(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Producto eliminado correctamente",
	})
}

// ListarProductos maneja la solicitud para listar todos los productos
func ListarProductos(w http.ResponseWriter, r *http.Request) {
	// Obtiene la conexión de la base de datos
	db := utils.GetDB()
	// Llama a la funcion ListarProductos del paquete services para obtener la lista de productos
	productos, err := services.ListarProductos(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Codifica la lista de productos en formato JSON y la envía en la respuesta
	json.NewEncoder(w).Encode(productos)
}

// BuscarProducto maneja la solicitud para buscar productos por nombre o categoría
func BuscarProducto(w http.ResponseWriter, r *http.Request) {
	// Obtiene el parámetro de búsqueda de la URL
	buscar := r.URL.Query().Get("buscar")
	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la funcion BuscarProducto del paquete services para buscar productos en la base de datos
	productos, err := services.BuscarProducto(db, buscar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Codifica la lista de productos encontrados en formato JSON y la envia en la respuesta
	json.NewEncoder(w).Encode(productos)
}
