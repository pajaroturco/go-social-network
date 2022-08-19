package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pajaroturco/go-social-network/db"
	"github.com/pajaroturco/go-social-network/models"
)

/* Registro es la func para crear un usuario en mongo */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos de entrada "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El correo de usuario es requerido ", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser al menos de seis caracteres ", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := db.ChequeoUsuarioExistente(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese correo ", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertoRegistroUsuario(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al guardar un registro de usuario "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "Error al insertar el registro de un usuario "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
