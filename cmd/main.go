package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Laura-2950/API-Go-Products/API-Go-Products/cmd/server/handler"
	_ "github.com/Laura-2950/API-Go-Products/API-Go-Products/doc"
	"github.com/Laura-2950/API-Go-Products/API-Go-Products/internal/product"
	"github.com/Laura-2950/API-Go-Products/API-Go-Products/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API-Go-Products
// @version         1.0
// @description     Products API example.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /products

// @externalDocs.description  OpenAPI
// @externalDocs.url          http://localhost:8080/swagger/index.html
func main() {

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

	storage := store.SqlStore{DB: db}
	repo := product.Repository{Storage: &storage}
	serv := product.Service{Repository: &repo}
	prodHandler := handler.ProductHandler{ProductService: &serv}
	r := gin.Default()

	r.GET("ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
	productGroup := r.Group("/products")
	{
		productGroup.GET(":id", prodHandler.GetById)
		productGroup.GET("", prodHandler.GetAll)
		productGroup.POST("", prodHandler.NewProduct)
		productGroup.DELETE(":id", prodHandler.Delete)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
