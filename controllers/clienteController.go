package controllers

import (
	"Inventario/services"
	"Inventario/utils"    
	"encoding/json"       
	"log"                 
	"net/http"            
	"strconv"             
	"github.com/gorilla/mux" 
)

// AgregarCliente maneja la solicitud para agregar un nuevo cliente
func AgregarCliente(w http.ResponseWriter, r *http.Request) {
	var cliente services.Cliente
	// Decodifica el cuerpo de la solicitud JSON en la estructura Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la funcion AgregarCliente del paquete services para agregar el cliente a la base de datos
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

// ListarClientes maneja la solicitud para listar todos los clientes
func ListarClientes(w http.ResponseWriter, r *http.Request) {
	// Obtiene la conexion de la base de datos
	db := utils.GetDB()
	// Llama a la funcion ListarClientes del paquete services para obtener la lista de clientes
	clientes, err := services.ListarClientes(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// Codifica la lista de clientes en formato JSON y la envia en la respuesta
	json.NewEncoder(w).Encode(clientes)
}

// ModificarCliente maneja la solicitud para modificar un cliente existente
func ModificarCliente(w http.ResponseWriter, r *http.Request) {
    var cliente services.Cliente
    // Decodifica el cuerpo de la solicitud JSON en la estructura Cliente
    if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
        log.Printf("Error al decodificar el cuerpo de la solicitud: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    log.Printf("Datos del cliente recibidos: %+v", cliente)

    // Obtiene la conexion de la base de datos
    db := utils.GetDB()
    // Llama a la funcion ModificarCliente del paquete services para modificar el cliente en la base de datos
    err := services.ModificarCliente(db, cliente)
    if err != nil {
        log.Printf("Error al modificar el cliente en la base de datos: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("Cliente modificado correctamente")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Cliente modificado correctamente",
    })
}

// EliminarCliente maneja la solicitud para eliminar un cliente por su ID
func EliminarCliente(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    // Convierte el ID del cliente de string a int
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    // Obtiene la conexion de la base de datos
    db := utils.GetDB()
    // Llama a la funcion EliminarCliente del paquete services para eliminar el cliente de la base de datos
    err = services.EliminarCliente(db, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Cliente eliminado correctamente",
    })
}


