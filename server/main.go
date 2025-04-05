package main

import (
	"net/http"
	"server/utils"
)

func main() {
	// Multiplexor de servidor HTTP
	// Examina la URL de cada solicitud HTTP y la dirige al controlador correspondiente
	mux := http.NewServeMux()

	// Maneja funciones segun URL de la solicitud
	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	// Inicia un servidor que escucha en el puerto 8080
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
