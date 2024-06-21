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

func ListarClientes(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	clientes, err := services.ListarClientes(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clientes)
}

func ModificarCliente(w http.ResponseWriter, r *http.Request) {
    var cliente services.Cliente
    if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
        log.Printf("Error al decodificar el cuerpo de la solicitud: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    log.Printf("Datos del cliente recibidos: %+v", cliente)

    db := utils.GetDB()
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



func EliminarCliente(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    db := utils.GetDB()
    err = services.EliminarCliente(db, id) // Aquí podría estar el conflicto si ya hay otra variable 'err' definida
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Cliente eliminado correctamente",
    })
}
