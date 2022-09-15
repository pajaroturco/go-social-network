package db

import (
	"context"
	"fmt"
	"time"

	"github.com/pajaroturco/go-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeoUsuariosTodos */
func LeoUsuariosTodos(ID string, pagina int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((pagina - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, _ = ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.Banner = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Email = ""

			results = append(results, &s)
		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)

	return results, true
}
