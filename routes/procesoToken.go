package routes

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/pajaroturco/go-social-network/db"
	"github.com/pajaroturco/go-social-network/models"
)

var Email string
var IDUsuario string

/* ProcesoToken */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrolloEnUdemy")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := db.ChequeoUsuarioExistente(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}
