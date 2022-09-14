package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pajaroturco/go-social-network/db"
)

/*LeoTweets funcion para endpoint de leer tweets*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := db.LeoTweets(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer tweets", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)

}
