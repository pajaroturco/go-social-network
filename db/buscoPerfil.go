package db

import (
	"context"
	"log"
	"time"

	"github.com/pajaroturco/go-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	perfil.Password = ""
	if err != nil {
		log.Fatal("Registro no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil
}
