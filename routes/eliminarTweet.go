package routes

import (
	"net/http"

	"github.com/pajaroturco/go-social-network/db"
)

/*EliminarTweet router para eliminar tweet*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := db.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al eliminar un tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
