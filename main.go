package main

import (
	"Inventario/routers"  
	"Inventario/utils"    
	"log"                 
	"net/http"            
)

func main() {
	// Inicializa la conexión a la base de datos
	utils.InitDB()

	// Configura el enrutador con todas las rutas de la aplicacion
	router := routers.InitRouter()

	// Inicia el servidor HTTP en el puerto 8080 con el enrutador configurado
	log.Fatal(http.ListenAndServe(":8080", router))
}

