package middlewares

import (
	"net/http"

	"github.com/pajaroturco/go-social-network/db"
)

/* ChequeoDB es el middleware que me permite conocer el estado de la db */
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	}
}
