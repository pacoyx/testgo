package models

import (
	"context"
	"fmt"
	"log"

	"github.com/pacoyx/go-cron-test/data"
	"github.com/pacoyx/go-cron-test/utils"
	"go.mongodb.org/mongo-driver/bson"
)

//Proyecto tabla de sql server para TEST
type Proyecto struct {
	Idproyecto     string `json:"idproyecto"`
	Nombreproyecto string `json:"nombreproyecto"`
	Duracionmeses  int    `json:"duracionmeses"`
	Inicio         string `json:"inicio"`
}

//Testgo estructura para testing
type Testgo struct {
	Nombre string `json:"nombre"`
	Estado string `json:"estado"`
	Fecha  string `json:"fecha"`
}

// Shopping - Model
type Shopping struct {
	Idper    string   `bson:"idper" json:"idper"`
	User     int      `bson:"user" json:"user"`
	Products []string `bson:"products" json:"products"`
	Payment  string   `bson:"payment" json:"payment"`
	Total2   int      `bson:"total2" json:"total2"`
	Estado   string   `bson:"estado" json:"estado"`
}

var (
	ctx context.Context
)

//TestCnxMysql Insert tabla en mysql
func TestCnxMysql() string {

	fecha := utils.GetFechaL()

	clienteaws := data.GetClientDBMySQLaws()
	insert, err := clienteaws.Query("INSERT INTO test VALUES ( 'parametros', 'inactivo','" + fecha + "' )")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	clienteaws.Close()

	return "se grabo correctamente en aurora"
}

//TestCnxMysqlSelect lista tabla
func TestCnxMysqlSelect() (string, *[]Testgo) {
	db := data.GetClientDBMySQLaws()

	results, err := db.Query("SELECT Nombre, Estado, Fecha FROM test")
	if err != nil {
		panic(err.Error())
	}

	var arrTest []Testgo

	for results.Next() {
		var tag Testgo

		err = results.Scan(&tag.Nombre, &tag.Estado, &tag.Fecha)
		if err != nil {
			panic(err.Error())
		}

		arrTest = append(arrTest, tag)

	}

	log.Print(arrTest)

	defer results.Close()
	db.Close()

	return "listado ok", &arrTest
}

//TestCnxSQLServer prueba conexion
func TestCnxSQLServer() string {
	msgRes := data.GetPruebaCnx()
	return "todo ok en sql Server | from base:" + msgRes
}

//TestCnxSQLServerInsert test de insert en sql server
func TestCnxSQLServerInsert(eProyecto Proyecto) string {
	db := data.GetClienteMMSS()
	fecha := utils.GetFechaL()
	tsql := fmt.Sprintf("INSERT INTO proyecto(idproyecto, nombreproyecto, duracionmeses, inicio) VALUES( '%s', '%s', '%d', '%s' );",
		eProyecto.Idproyecto, eProyecto.Nombreproyecto, eProyecto.Duracionmeses, fecha)

	_, err := db.Exec(tsql)

	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return "err"
	}

	return "todo ok en sql Server | Insert from base:"
}

//TestCnxSQLServerSelect test select sql server
func TestCnxSQLServerSelect() (string, *[]Proyecto) {
	db := data.GetClienteMMSS()
	tsql := fmt.Sprintf("SELECT idproyecto, nombreproyecto, duracionmeses, inicio FROM proyecto;")
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return "error select", nil
	}
	defer rows.Close()

	var arrTest []Proyecto
	for rows.Next() {
		var eproy Proyecto
		err := rows.Scan(&eproy.Idproyecto, &eproy.Nombreproyecto, &eproy.Duracionmeses, &eproy.Inicio)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return "error select", nil
		}
		arrTest = append(arrTest, eproy)
	}

	return "select sql server OK", &arrTest

}

//TestCnxMongodb prueba conexion
func TestCnxMongodb() string {
	_ = data.GetClientmongodb()
	return "todo ok en mongodb desde Models"
}

//TestCnxMongodbInsert prueba Insert en mongo conexion
func TestCnxMongodbInsert(dato Shopping) string {

	client := data.GetClientmongodb()
	collection := client.Database("test").Collection("shoppings")

	insertResult, err := collection.InsertOne(context.TODO(), dato)
	if err != nil {
		log.Fatal(err)
		return "Error"
	}

	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

	return "Comando Insert en mongoDB ok"
}

//TestCnxMongodbSelect test select para mongodb
func TestCnxMongodbSelect() (string, []*Shopping) {

	client := data.GetClientmongodb()
	collection := client.Database("test").Collection("shoppings")
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var arrShopping []*Shopping

	for cur.Next(context.TODO()) {
		var docu Shopping
		err = cur.Decode(&docu)
		if err != nil {
			log.Fatal("Error on Decoding the document en mongoDB", err)
		}
		arrShopping = append(arrShopping, &docu)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return "select [OK]", arrShopping

}
