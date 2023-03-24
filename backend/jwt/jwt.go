package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pajaroturco/go-social-network/models"
	"github.com/pajaroturco/go-social-network/sistema"
)

/* GeneroJWT es la func para crear un token */
func GeneroJWT(t models.Usuario) (string, error) {

	miLlave := sistema.GetEnv("JWT_SECRET")
	if miLlave == "" {
		return "", errors.New("no se pudo obtener la llave")
	}

	miClave := []byte(miLlave)

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
