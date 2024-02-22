package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"Crud-Golang-RabbitMQ/internal/model"
	"Crud-Golang-RabbitMQ/internal/service"
	"Crud-Golang-RabbitMQ/internal/queue"

	"github.com/gin-gonic/gin"
)

type ClientesHandler struct {
	Service service.ClienteService
	Queue   queue.QueueClient
}

// função NewClientesHandler para receber o QueueClient
func NewClientesHandler(service service.ClienteService, queue queue.QueueClient) *ClientesHandler {
	return &ClientesHandler{
		Service: service,
		Queue:   queue,
	}
}

// @Summary Adiciona um novo cliente
// @Description Cria um novo cliente com os dados enviados
// @Tags clientes
// @Accept json
// @Produce json
// @Param cliente body model.Clientes true "Dados do Cliente"
// @Success 201 {object} model.Clientes
// @Failure 400 {object} HTTPError
// @Router /clientes [post]
func (h *ClientesHandler) CreateCliente(c *gin.Context) {
	var cliente model.Clientes
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados de entrada inválidos", "details": err.Error()})
		return
	}

	if err := h.Service.CreateCliente(c.Request.Context(), cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao criar cliente", "details": err.Error()})
		return
	}

	// Utiliza SHA256 para gerar um hash do nome do cliente
	hasher := sha256.New()
	hasher.Write([]byte(cliente.Nome))
	hashedNome := hex.EncodeToString(hasher.Sum(nil))

	// Utiliza os primeiros 8 caracteres do hash para o nome da fila
	nomeDaFila := "clienteQueue-" + hashedNome[:8]

	// Mensagem a ser publicada no RabbitMQ
	mensagem := fmt.Sprintf("Novo cliente criado: %s", cliente.Nome) 
	if err := h.Queue.Publish(nomeDaFila, []byte(mensagem)); err != nil {
		log.Printf("Erro ao publicar mensagem no RabbitMQ: %v", err)
	}

	c.JSON(http.StatusCreated, cliente)
}

// @Summary Lista todos os clientes
// @Description Retorna uma lista de todos os clientes disponíveis
// @Tags clientes
// @Produce json
// @Success 200 {array} model.Clientes
// @Router /clientes [get]
func (h *ClientesHandler) ListarClientes(c *gin.Context) {
	// Chame o serviço para obter a lista de clientes
	clientes, err := h.Service.ListarClientes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao listar clientes", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clientes)
}



// @Summary Busca um cliente por ID
// @Description Retorna um cliente dado seu ID
// @Tags clientes
// @Produce json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} model.Clientes
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /clientes/{id} [get]
func (h *ClientesHandler) GetClienteByID(c *gin.Context) {
	// Obtive o ID como int64 a partir dos parâmetros da rota
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Converti o ID de int64 para uint antes de passar para o serviço
	cliente, err := h.Service.GetClienteByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cliente não encontrado", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

// @Summary Atualiza um cliente
// @Description Atualiza os dados de um cliente com base no ID fornecido
// @Tags clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Param cliente body model.Clientes true "Dados atualizados do Cliente"
// @Success 200 {object} model.Clientes
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /clientes/{id} [put]
func (h *ClientesHandler) UpdateCliente(c *gin.Context) {
	var cliente model.Clientes
	// Parse o ID do cliente da URL como int64
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || c.ShouldBindJSON(&cliente) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados de entrada inválidos"})
		return
	}
	// Converti o id de int64 para uint antes de atribuir ao ID do cliente
	cliente.ID = uint(id)

	if err := h.Service.UpdateCliente(c.Request.Context(), cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao atualizar cliente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

// @Summary Deleta um cliente
// @Description Deleta um cliente com base no ID fornecido
// @Tags clientes
// @Produce json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} HTTPMessage
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /clientes/{id} [delete]
func (h *ClientesHandler) DeleteCliente(c *gin.Context) {
	// Obtive o ID como int64 a partir dos parâmetros da rota
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Converti o ID de int64 para uint
	if err := h.Service.DeleteCliente(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao deletar cliente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cliente apagado com sucesso"})
}

// Estruturas auxiliares para a doc do Swagger.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HTTPMessage é uma estrutura de mensagem de sucesso.
type HTTPMessage struct {
	Message string `json:"message"`
}

