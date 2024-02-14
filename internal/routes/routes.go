package routes

import (
	"Crud-Golang-RabbitMQ/internal/handler"

	"github.com/gin-gonic/gin"
)

// Config Rotas
func SetupRouter(r *gin.Engine, clienteHandler *handler.ClientesHandler) {
	r.POST("/clientes", clienteHandler.CreateCliente)
	r.GET("/clientes", clienteHandler.ListarClientes)
	r.GET("/clientes/:id", clienteHandler.GetClienteByID)
	r.PUT("/clientes/:id", clienteHandler.UpdateCliente)
	r.DELETE("/clientes/:id", clienteHandler.DeleteCliente)
}
