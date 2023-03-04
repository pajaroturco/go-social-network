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
	router.HandleFunc("/tweets", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.LeoTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/avatar", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.SubirAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlewares.ChequeoDB(routes.ObtenerAvatar)).Methods("GET")

	router.HandleFunc("/banner", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.SubirBanner))).Methods("POST")
	router.HandleFunc("/banner", middlewares.ChequeoDB(routes.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/relacion", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.AltaRelacion))).Methods("POST")
	router.HandleFunc("/relacion", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/relacion", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/usuarios", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/tweetsSeguidores", middlewares.ChequeoDB(middlewares.ValidaJWT(routes.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
