package db

import (
	"context"
	"time"

	"github.com/pajaroturco/go-social-network/models"
)

/*BorroRelacion elimina la relacion en la base de datos*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
