package data

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client //cliente mongodb
var dbname string

// GetClientmongodb obtiene el cliente de mongodb para conexion
func GetClientmongodb() *mongo.Client {
	return db
}

// GetDBName obtiene el nombre de la Colleccion mongodb
func GetDBName() string {
	return os.Getenv("db_name")
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")

	clientOptions := options.Client().ApplyURI("mongodb+srv://" + username + ":" + password + "@cluster0-0cggd.mongodb.net/test?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db = client

}
