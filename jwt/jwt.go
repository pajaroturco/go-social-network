package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pajaroturco/go-social-network/models"
)

/* GeneroJWT es la func para crear un token */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("MasterdelDesarrolloEnUdemy")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err

	}

	return tokenStr, nil
}
