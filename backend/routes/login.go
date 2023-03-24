package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pajaroturco/go-social-network/db"
	"github.com/pajaroturco/go-social-network/jwt"
	"github.com/pajaroturco/go-social-network/models"
)

/* Login es la func para crean un login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Credenciales inválidas "+err.Error(), http.StatusUnauthorized)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El correo de usuario es requerido ", http.StatusBadRequest)
		return
	}

	usu, existe := db.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Credenciales inválidas ", http.StatusUnauthorized)
		return
	}

	jwtKey, err := jwt.GeneroJWT(usu)
	if err != nil {
		http.Error(w, "Ocurrio un error "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
