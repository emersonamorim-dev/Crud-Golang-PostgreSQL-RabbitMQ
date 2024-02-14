package repository_test

import (
	"Crud-Golang-RabbitMQ/internal/model"
	"Crud-Golang-RabbitMQ/internal/repository"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestClienteRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cliente Repository Suite")
}

var _ = Describe("ClienteRepository", func() {
	var (
		db           *gorm.DB
		clienteRepo  repository.ClientesRepository
		testClientes []model.Clientes
	)

	BeforeSuite(func() {
		// Obtenha o diretório do teste
		testDir := filepath.Dir(os.Args[0])

		// Carrega variáveis de ambiente do arquivo .env no diretório de teste
		err := godotenv.Load(filepath.Join(testDir, "..", "..", ".env"))
		Expect(err).To(BeNil())

		// Conecta ao BD PostgreSQL usando as variáveis de ambiente
		dsn := "host=" + os.Getenv("PGHOST") + " user=" + os.Getenv("PGUSER") + " password=" + os.Getenv("PGPASSWORD") + " dbname=" + os.Getenv("PGDATABASE") + " port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		Expect(err).To(BeNil())

		db.AutoMigrate(&model.Clientes{})
		clienteRepo = repository.NewClientesRepository(db)

		// Insira alguns clientes de teste no BD
		testClientes = []model.Clientes{
			{Nome: "Emerson Amorim"},
			{Nome: "Emerson DEV"},
			{Nome: "Emerson Luiz"},
		}
		for _, cliente := range testClientes {
			clienteRepo.Create(context.Background(), cliente)
		}
	})

	Describe("Create", func() {
		It("Deve criar um novo cliente", func() {
			newCliente := model.Clientes{Nome: "David"}
			err := clienteRepo.Create(context.Background(), newCliente)
			Expect(err).To(BeNil())

			// Verifica se o cliente foi inserido corretamente
			var createdCliente model.Clientes
			db.First(&createdCliente, newCliente.ID)
			Expect(createdCliente.Nome).To(Equal(newCliente.Nome))
		})
	})

	Describe("ListarTodos", func() {
		It("Deve listar todos os clientes disponíveis", func() {
			clientes, err := clienteRepo.ListarTodos(context.Background())
			Expect(err).To(BeNil())
			Expect(clientes).To(HaveLen(len(testClientes)))
		})
	})

	Describe("GetByID", func() {
		It("Deve retornar um cliente pelo ID", func() {
			expectedCliente := testClientes[0]
			cliente, err := clienteRepo.GetByID(context.Background(), expectedCliente.ID)
			Expect(err).To(BeNil())
			Expect(cliente.Nome).To(Equal(expectedCliente.Nome))
		})

		It("Deve retornar um erro se o cliente não existir", func() {
			cliente, err := clienteRepo.GetByID(context.Background(), 999)
			Expect(err).ToNot(BeNil())
			Expect(errors.Is(err, gorm.ErrRecordNotFound)).To(BeTrue())
			Expect(cliente).To(Equal(model.Clientes{}))
		})
	})

	Describe("Update", func() {
		It("Deve atualizar um cliente existente", func() {
			expectedCliente := testClientes[0]
			expectedCliente.Nome = "Eva"
			err := clienteRepo.Update(context.Background(), expectedCliente)
			Expect(err).To(BeNil())

			// Verifica se o cliente foi atualizado corretamente
			var updatedCliente model.Clientes
			db.First(&updatedCliente, expectedCliente.ID)
			Expect(updatedCliente.Nome).To(Equal(expectedCliente.Nome))
		})
	})

	Describe("Delete", func() {
		It("Deve deletar um cliente existente", func() {
			cliente := testClientes[0]
			err := clienteRepo.Delete(context.Background(), cliente.ID)
			Expect(err).To(BeNil())

			// Verifica se o cliente foi deletado corretamente
			var deletedCliente model.Clientes
			db.First(&deletedCliente, cliente.ID)
			Expect(deletedCliente).To(Equal(model.Clientes{}))
		})
	})
})
