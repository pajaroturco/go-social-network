package main

import (
	"log"

	"github.com/pajaroturco/go-social-network/db"
	"github.com/pajaroturco/go-social-network/handlers"
)

func main() {
	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}

	handlers.Manejadores()
}
