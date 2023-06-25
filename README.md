
Adding go.mod:
go mod init github.com/Laura-2950/API-Go-Products.git

Adding go.sum:
go get -u github.com/gin-gonic/gin
go get "github.com/go-sql-driver/mysql" 



docker-compose up
go run .\cmd\main.go


port: 8080

Endpoints:
- Get Product by ID: http://localhost:8080/products/:id
- Get All Products: http://localhost:8080/products
- Post New Product: http://localhost:8080/products
- Delete Product by ID: http://localhost:8080/products/:id
