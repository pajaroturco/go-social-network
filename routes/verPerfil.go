package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pajaroturco/go-social-network/db"
)

/*ver Perfil*/
func verPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("Id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intenter buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
