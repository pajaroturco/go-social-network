package db

import (
	"context"
	"time"

	"github.com/pajaroturco/go-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoTweet funcion para guardar un tweet*/
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweets")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	resultado, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := resultado.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
