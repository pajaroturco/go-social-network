package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pajaroturco/go-social-network/db"
	"github.com/pajaroturco/go-social-network/models"
)

/*SubirAvatar router para subit avatar*/
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, _ := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = db.ModificoRegistro(usuario, IDUsuario)

	if err != nil || !status {
		http.Error(w, "Error al grabar la imagen en bd!"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
