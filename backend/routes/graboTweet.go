package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pajaroturco/go-social-network/db"
	"github.com/pajaroturco/go-social-network/models"
)

/*GraboTweet funcion para endpoitn de guardar tweet*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Datos incorrectos, guarda tweet"+err.Error(), 400)
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	var status bool
	_, status, err = db.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar guardar el tweet"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se la logrado guardar el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
