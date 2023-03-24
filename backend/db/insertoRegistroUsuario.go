package db

import (
	"context"
	"time"

	"github.com/pajaroturco/go-social-network/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistroUsuario es la parada final con la BD */
func InsertoRegistroUsuario(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
