package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pajaroturco/go-social-network/db"
)

/*LeoTweetsSeguidores lee los tweets de todos nuestros seguidores*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina mayor a cero", http.StatusBadRequest)
		return
	}

	respuesta, correcto := db.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)

}
