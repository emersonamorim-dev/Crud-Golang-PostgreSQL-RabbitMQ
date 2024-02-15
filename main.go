package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Crud-Golang-RabbitMQ/docs"
	"Crud-Golang-RabbitMQ/internal/handler"
	"Crud-Golang-RabbitMQ/internal/queue"
	"Crud-Golang-RabbitMQ/internal/repository"
	"Crud-Golang-RabbitMQ/internal/routes"
	"Crud-Golang-RabbitMQ/internal/service"
	"Crud-Golang-RabbitMQ/migrations"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Crud-Golang-RabbitMQ API RESTful
// @version 1.0
// @description API RESTful CRUD em Golang usando o Swagger.
// @termsOfService https://github.com/emersonamorim-dev
// @contact.name Emerson Amorim DEV
// @contact.url https://github.com/emersonamorim-dev
// @host localhost:8081
// @BasePath /api/v1
func main() {

	// Config do Gin
	r := gin.Default()

	// Config do RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		log.Fatal("RABBITMQ_URL é requerido")
	}
	rabbitMQConn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Falha ao conectar-se ao RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()

	// Conexão com o RabbitMQ
	rabbitMQ, err := queue.NewRabbitMQ(rabbitMQConn)
	if err != nil {
		log.Fatalf("Falha ao inicializar RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Conexão com o PostgreSQL usando GORM
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL é requerido")
	}
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Não é possível se conectar ao banco de dados: %v", err)
	}

	// Executa as migrações
	migrator := &migrations.CreateClientesTable{}
	if err := migrator.Migrate(db); err != nil {
		log.Fatalf("Falha ao executar migrações: %v", err)
	}

	// Repository
	clientesRepo := repository.NewClientesRepository(db)

	// Service
	clienteService := service.NewClienteService(clientesRepo, rabbitMQ)

	// Handler
	clientesHandler := handler.NewClientesHandler(clienteService)

	// Config das rotas usando o clientesHandler
	routes.SetupRouter(r, clientesHandler)

	// Config do diretório de templates HTML
	r.LoadHTMLGlob("templates/*")

	// Rota de boas-vindas
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Teste Crud Golang com Postgres e RabbitMQ",
		})
	})

	// Carrega a doc do Swagger
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Inicia o servidor
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}
}
