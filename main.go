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


