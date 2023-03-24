package db

import (
	"github.com/pajaroturco/go-social-network/models"
	"golang.org/x/crypto/bcrypt"
)

/* IntentoLogin */
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ChequeoUsuarioExistente(email)
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true

}
