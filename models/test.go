package models

import (
	"github.com/pacoyx/go-cron-test/data"
)

//Testgo estructura para testing
type Testgo struct {
	Nombre string `json:"nombre"`
	Estado string `json:"estado"`
}

//TestCnxMysql prueba conexion
func TestCnxMysql() string {

	clienteaws := data.GetClientDBMySQLaws()
	insert, err := clienteaws.Query("INSERT INTO test VALUES ( 'corona', 'activo' )")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	clienteaws.Close()

	return "se grabo correctamente en aurora"
}

//TestCnxSQLServer prueba conexion
func TestCnxSQLServer() string {
	//db := GetClienteMySql()
	return "todo ok en sql Server"
}

//TestCnxMongodb prueba conexion
func TestCnxMongodb() string {
	//db := GetClienteMySql()
	return "todo ok en mongodb"
}
