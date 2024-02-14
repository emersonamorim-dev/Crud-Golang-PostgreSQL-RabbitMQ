### Crud Golang com Postgres e RabbitMQ 🚀 🔄 🌐

Codificação para um Teste Técnico com Innolevels para um CRUD (Create, Read, Update, Delete) desenvolvido em GoLang com framework Gin, utilizando PostgreSQL e RabbitMQ para persistência e comunicação assíncrona. 
Desenvolvido uma API RESTful CRUD em Golang que utilizei requisições HTTP para realizar operações CRUD em um banco de dados Postgres. Os registros a serem inseridos na base de dados serão enviados para uma fila na RabbitMQ e posteriormente processados por um worker. Uso de Swagger para gerenciar os Endpoints e facilita a consulta dos Endpoints. A aplicação está implementada com Testes com uso do Testfy uma framework para testes. Também usado o Gorm para subir as Migrations de forma automatizada para o Postgres. Os dados são relacionados a dados pessoais de clientes seguindo boas práticas de programação.

#### Tecnologias Usadas:
 
- GoLang 1.21
- Gin
- Gorm
- TestFy
- Postgres
- RabbitMQ
- Docker
- Swagger
- Ubuntu/WSL - 100% funcional
- Windows/Terminal - 100% funcional

### Comandos para rodar o projeto:

``` 
go run main.go
```

### Pacotes para rodar o projeto:

- Pacote do Gorm ORM (Object-Relational Mapping)
``` 
go get -u gorm.io/gorm
```

- Pacote do Postgres
``` 
go get -u gorm.io/driver/postgres
```

- Pacote do RabbitMQ
``` 
go get github.com/streadway/amqp

```

- Pacotes para instalar Swagger
``` 
go get -u github.com/swaggo/swag/cmd/swag

```

``` 
go get -u github.com/swaggo/echo-swagger

```

- Pacote para instalar Testfy

``` 
go get github.com/stretchr/testify
```

#### Principais Endpoints do Projeto para uso com Swagger ou Postman

Acesso ao Swagger quando subir aplicação nesse link:

``` 
http://localhost:8081/swagger/index.html
```

- Consulta com Metódo Post - Via Postman
``` 
http://localhost:8081/clientes/

```
``` 
{
  "nome": "Emerson",
  "sobrenome": "Amorim FullStack 18",
  "contato": "11999887766",
  "endereco": "Rua sn",
  "nascimento": "1981-02-18",
  "cpf": "81345678918"
}
```

- Consulta com Metódo Get Listar todos
``` 
http://localhost:8081/clientes/

```

- Consulta com Metódo Get por ID
``` 
http://localhost:8081/clientes/1

```
``` 
{
  "nome": "Emerson",
  "sobrenome": "Amorim FullStack 18",
  "contato": "11999887766",
  "endereco": "Rua sn",
  "nascimento": "1981-02-18",
  "cpf": "81345678918"
}
```


- Consulta com Metódo Update por ID
``` 
http://localhost:8081/clientes/1

```
``` 
{
  "nome": "Emerson",
  "sobrenome": "Amorim FullStack 18",
  "contato": "11999887766",
  "endereco": "Rua sn 123",
  "nascimento": "1981-02-18",
  "cpf": "81345678918"
}
```

- Consulta com Metódo Delete por ID
``` 
http://localhost:8081/clientes/1

```


#### Funcionalidades
Adicionar Cliente: Adiciona um novo cliente ao sistema.
Listar Clientes: Retorna uma lista de todos os clientes cadastrados.
Buscar Cliente por ID: Retorna um cliente com base em seu ID.
Atualizar Cliente: Atualiza os dados de um cliente existente.
Deletar Cliente: Remove um cliente do sistema.

#### Estrutura do Projeto

- O projeto está dividido em diferentes pacotes e módulos para uma melhor organização e separação de responsabilidades:

- cmd: Contém os pontos de entrada da aplicação (api e worker).
- docs: Documentação do projeto.
- internal: Código-fonte da aplicação.
- config: Configurações da aplicação.
- db: Lógica relacionada ao banco de dados.
- handler: Manipuladores HTTP para cada rota.
- model: Definição das estruturas de dados.
- queue: Integração com RabbitMQ.
- repository: Acesso aos dados do banco de dados.
- routes: Definição das rotas da API.
- service: Lógica de negócio da aplicação.
- migrations: Scripts de migração do banco de dados.
- swagger: Arquivos relacionados à documentação Swagger.
- templates: Templates HTML para renderização de páginas web.
- tests: Testes unitários e de integração.

#### Melhores Práticas, Design Patterns e SOLID
Handler (handler/clientes_handler.go)

#### Melhores Práticas:

Utiliza o framework Gin para roteamento e manipulação de requisições HTTP.
Validação de entrada utilizando o método ShouldBindJSON.
Respostas HTTP consistentes e informativas.
Design Patterns:

- Injeção de dependência no construtor NewClientesHandler, facilitando os testes e a manutenção.
Utilização do padrão de projeto Controller.

#### SOLID:

- Single Responsibility Principle (SRP): Cada método do handler tem uma única responsabilidade, como criar, listar, buscar, atualizar e deletar clientes.
- Open/Closed Principle (OCP): A estrutura permite a extensão de novos endpoints sem modificar o código existente.
- Dependency Inversion Principle (DIP): O handler depende de interfaces abstratas, não de implementações concretas, facilitando a substituição de componentes.

#### Repository (repository/clientes_repository.go)
Melhores Práticas:

- Separação clara das operações de CRUD em métodos distintos.
- Uso de contextos para operações com banco de dados.
- Manipulação de erros de forma adequada.

#### Design Patterns:

Implementação da interface ClientesRepository, permitindo a troca de implementações sem modificar o código cliente.
SOLID:

- Single Responsibility Principle (SRP): Cada método do repositório tem uma única responsabilidade, como criar, listar, buscar, atualizar e deletar clientes.
- Interface Segregation Principle (ISP): A interface ClientesRepository é coesa e não impõe métodos desnecessários aos seus clientes.
- Dependency Inversion Principle (DIP): O repositório depende apenas de abstrações, não de implementações concretas, permitindo a troca de ORM sem alterar o código do serviço.
Service (service/clientes_service.go)

#### Melhores Práticas:

Separação clara da lógica de negócio em métodos distintos.
Uso de interfaces para abstrair dependências externas.
Tratamento de erros robusto.

#### Design Patterns:
Utilização do padrão Service para encapsular a lógica de negócio.

#### SOLID:

- Single Responsibility Principle (SRP): Cada método do serviço tem uma única responsabilidade, como criar, listar, buscar, atualizar e deletar clientes.
- Open/Closed Principle (OCP): O serviço é aberto para extensão (novos métodos), mas fechado para modificação.
- Dependency Inversion Principle (DIP): O serviço depende apenas de abstrações, não de implementações concretas, facilitando a substituição de componentes.


##### Bônus realizando uma Integração d o Jogo Tetris e um Endpoint de Envio de Dados Predefinidos
O projeto também inclui uma integração com o Jogo Tetris, onde um Endpoint específico envia dados predefinidos para o Backend em Golang quando jogador realiza pontuação completando uma linha na horizontal e ganha pontos no Jogo Tretris, proporcionando uma experiência Inovadora e bastante interativa.



#### Conclusão
O projeto usa Golang com GIN  que reflete em um compromisso com as melhores práticas de desenvolvimento, proporcionando uma base sólida para a construção de aplicativos web escaláveis e de alta qualidade. A arquitetura, a organização do código e a aplicação de conceitos como orientação a objetos e princípios SOLID demonstram a busca pela excelência no desenvolvimento de software.

### Autor:

Emerson Amorim [@emerson-amorim-dev](https://www.linkedin.com/in/emerson-amorim-dev/)
