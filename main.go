/*
package main

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
	}
*/
package main

import (
	"Inventario/routers"
	"Inventario/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	utils.InitDB()
	router := routers.InitRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
