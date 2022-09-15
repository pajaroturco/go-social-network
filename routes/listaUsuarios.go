package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pajaroturco/go-social-network/db"
)

/*ListaUsuarios funcion para endpoint de traer lista de usuarios*/
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTmp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pageTmp)

	result, status := db.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
