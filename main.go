/*package main

import (

	"Inventario/routers"
	"Inventario/utils"
	"log"
	"net/http"
	
)

func main() {
	utils.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}*/

// main.go
package main

import (
	"Inventario/routers"
	"Inventario/utils"
	"log"
	"net/http"
)

func main() {
	// Inicializa la base de datos
	utils.InitDB()

	// Inicializa el router con el middleware de CORS
	router := routers.InitRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

