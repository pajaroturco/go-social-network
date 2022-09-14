package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN objeto de conexion a base de datos*/
var MongoCN = ConectarDB()

//var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/twitter")
var clientOptions = options.Client().ApplyURI("mongodb+srv://pajaro:p4j4r0@cluster0.9vgnjdm.mongodb.net/twitter")

/*ConectarDB Funcion para conectar la base de datos*/
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa a DB")
	return client
}

/*ChequeoConnection Funcion para checar la conexion a la base de datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return 1

}
