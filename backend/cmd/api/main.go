package main

import (
	"log"

	"reservations-backend/internal/adapters/http"
)

func main() {
	e := http.GetServerInstance()

	err := e.Start(":8080")
	if err != nil {
		log.Println("Error al inciar el servidor")
	}
}
