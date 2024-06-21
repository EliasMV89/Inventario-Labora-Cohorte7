package utils

import (
    "net/http"
)

// CorsMiddleware es un middleware que configura los encabezados CORS para permitir solicitudes desde cualquier origen
func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Permitir cualquier origen (dominio) a acceder a los recursos
        w.Header().Set("Access-Control-Allow-Origin", "*")
        // Permitir métodos HTTP especificos
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        // Permitir ciertos encabezados en las solicitudes
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Si el método HTTP es OPTIONS, se responde con un estado 200 OK y se termina aquí
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        // Para otros métodos, se pasa la solicitud al siguiente manejador en la cadena.
        next.ServeHTTP(w, r)
    })
}
