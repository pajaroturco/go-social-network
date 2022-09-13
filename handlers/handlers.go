package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/pajaroturco/go-social-network/middlewares"
	"github.com/pajaroturco/go-social-network/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores set de puerto handler y pongo a escuchar el servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoDB(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoDB(routes.Login)).Methods("POST")
	router.HandleFunc("/perfil", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/perfil", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.GraboTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9090"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
