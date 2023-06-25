package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/cmd/server/handler"
	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/internal/product"
	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/pkg/middleware"
	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title API Go-Products
// @version 1.0
// @description This API Handle Products.
// @host      localhost:8080
// @BasePath  /products
// @termsOfService https://developers.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	fmt.Println("buscando el archivo")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error al intentar cargar archivo .env")
	}
	username := os.Getenv("USER_MYSQL")
	password := os.Getenv("PASS_MYSQL")
	dbName := os.Getenv("DB_MYSQL")

	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:33060)/%s", username, password, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing.Error())
	}

	storage := store.SqlStore{db}
	repo := product.Repository{&storage}
	serv := product.Service{&repo}
	prodHandler := handler.ProductHandler{&serv}
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
			productGroup := r.Group("/products")
	{ 
		productGroup.GET(":id", prodHandler.GetById)
		productGroup.GET("", prodHandler.GetAll)
		productGroup.POST("", middleware.Authentication(), prodHandler.NewProduct)
		productGroup.DELETE(":id", middleware.Authentication(), prodHandler.Delete)
	}

	r.Run(":8080")
}
