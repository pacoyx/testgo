package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"github.com/mitchellh/mapstructure"
	"github.com/pacoyx/go-cron-test/controller"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Indata struct {
	Producto string  `json:"producto"`
	Precio   float64 `json:"precio"`
	Cliente  string  `json:"cliente"`
}

type SlackRequestBody struct {
	Text string `json:"text"`
}

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting the application...")
	router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
	router.HandleFunc("/test", ValidateMiddleware(TestEndpoint)).Methods("GET")
	router.HandleFunc("/obtenerRoles", ValidateMiddleware(ObtenerRoles)).Methods("GET")
	router.HandleFunc("/ObtenerDatos", ValidateMiddleware(ObtenerDatos)).Methods("GET")
	router.HandleFunc("/Pdf", GenerarPdfServer).Methods("GET")
	router.HandleFunc("/TestMysql", controller.TestCnxMySQLController).Methods("GET")
	router.HandleFunc("/TestSqlServer", controller.TestCnxSQLServerController).Methods("GET")
	router.HandleFunc("/TestMongodb", controller.TestCnxMongodbController).Methods("GET")
	router.HandleFunc("/TestSlack", ValidateMiddleware(TestSalck)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

//TestingMysql probando la conexion a mysql aws
func TestingMysql(w http.ResponseWriter, req *http.Request) {
	fmt.Println("probado conexion , mensaje ok okok ...")
}

//CreateTokenEndpoint ccc
func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

//ProtectedEndpoint xxx
func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
	}
}

//ValidateMiddleware cccc
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

//TestEndpoint xxxx
func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

//ObtenerRoles xxxx
func ObtenerRoles(w http.ResponseWriter, req *http.Request) {

	decoded := context.Get(req, "decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user.Password)

}

//ObtenerDatos cccc
func ObtenerDatos(w http.ResponseWriter, req *http.Request) {

	reqBody, _ := ioutil.ReadAll(req.Body)

	var elreq Indata
	json.Unmarshal(reqBody, &elreq)

	json.NewEncoder(w).Encode(elreq)

}

// GeneratePdf generates our pdf by adding text and images to the page
// then saving it to a file (name specified in params).
func GeneratePdf(filename string) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 7, "Welcome to golangcode.com", "0", 0, "CM", false, 0, "")
	pdf.CellFormat(190, 5, "factura para cliente: ", "0", 0, "CM", false, 0, "")

	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"golang.jpg",
		80, 20,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)

	crearDirectorioSiNoExiste("./pdfs")
	//return pdf.OutputFileAndClose(filename)
	pdf.OutputFileAndClose("pdfs/" + filename)
}

//GenerarPdfServer cccc
func GenerarPdfServer(w http.ResponseWriter, req *http.Request) {

	fmt.Println("generando pdf...")

	reqBody, _ := ioutil.ReadAll(req.Body)
	var elreq Indata
	json.Unmarshal(reqBody, &elreq)

	count := 900
	for i := 0; i < count; i++ {
		go GeneratePdf(strconv.Itoa(i) + elreq.Producto)
	}
	//strconv.FormatFloat(i, 'f', 6, 64)

	fmt.Println("Termino proceso de generacion de pdf...")
}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func SendSlackNotification(webhookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}

//TestSalck ssds
func TestSalck(w http.ResponseWriter, req *http.Request) {
	webhookUrl := "https://hooks.slack.com/services/X1234"
	err := SendSlackNotification(webhookUrl, "Test Message from golangcode.com")
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("{'mensaje':'ok'}")
}

func crearDirectorioSiNoExiste(directorio string) {
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}
