package routes

import (
	"net/http"

	"github.com/pajaroturco/go-social-network/db"
	"github.com/pajaroturco/go-social-network/models"
)

/*BajaRelacion realiza el registro de la relacion entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar eliminar relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se la logrado eliminar la relacion", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
