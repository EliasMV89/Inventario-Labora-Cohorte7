package routers

import (
	"Inventario/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/clientes", controllers.AgregarCliente).Methods("POST")
	router.HandleFunc("/cliente/{id}", controllers.ModificarCliente).Methods("PUT")
	router.HandleFunc("/cliente/{id}", controllers.EliminarCliente).Methods("DELETE")


	router.HandleFunc("/productos", controllers.AgregarProducto).Methods("POST")
	router.HandleFunc("/productos", controllers.ListarProductos).Methods("GET")
	router.HandleFunc("/productos/buscar", controllers.BuscarProducto).Methods("GET")

	router.HandleFunc("/proveedores", controllers.AgregarProveedor).Methods("POST")

	router.HandleFunc("/ventas", controllers.AgregarVenta).Methods("POST")
	router.HandleFunc("/ventas/buscar", controllers.BuscarVentaPorFecha).Methods("GET")
	router.HandleFunc("/ventas/informe", controllers.GenerarInforme).Methods("GET")

	return router
}
