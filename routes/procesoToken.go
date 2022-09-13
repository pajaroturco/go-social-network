package routes

import (
	"github.com/pajaroturco/go-social-network/models"
)

/* ProcesoToken */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrolloEnUdemy")
	claims := &models.Claim{}

	splitToken := 
}
