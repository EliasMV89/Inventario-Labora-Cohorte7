package routers

import (
	"Inventario/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Clientes
	router.HandleFunc("/clientes", controllers.AgregarCliente).Methods("POST")
	router.HandleFunc("/cliente/{id}", controllers.ModificarCliente).Methods("PUT")
	router.HandleFunc("/cliente/{id}", controllers.EliminarCliente).Methods("DELETE")

	// Productos 
	router.HandleFunc("/productos", controllers.AgregarProducto).Methods("POST")
	router.HandleFunc("/producto/{id}", controllers.ModificarProducto).Methods("PUT")
	router.HandleFunc("/producto/{id}", controllers.EliminarProducto).Methods("DELETE")
	router.HandleFunc("/productos", controllers.ListarProductos).Methods("GET")
	router.HandleFunc("/productos/buscar", controllers.BuscarProducto).Methods("GET")
	
	// Proveedores
	router.HandleFunc("/proveedores", controllers.AgregarProveedor).Methods("POST")
	router.HandleFunc("/proveedor/{id}", controllers.ModificarProveedor).Methods("PUT")
	router.HandleFunc("/proveedor/{id}", controllers.EliminarProveedor).Methods("DELETE")	

	// Ventas
	router.HandleFunc("/ventas", controllers.AgregarVenta).Methods("POST")
	router.HandleFunc("/ventas/buscar", controllers.BuscarVentaPorFecha).Methods("GET")
	router.HandleFunc("/ventas/informe", controllers.GenerarInforme).Methods("GET")

	return router
}
