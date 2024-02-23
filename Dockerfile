FROM golang:1.21.5-alpine AS builder

# Configurar o diretório de trabalho
WORKDIR /app

# Copiar o módulo go e baixar dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar os arquivos de código-fonte do projeto
COPY . .

# Compilar a aplicação para um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Etapa 2: Configurar a imagem de produção
FROM alpine:latest

# Adicionar suporte a CA certificates para chamadas HTTPS dentro do contêiner
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário pré-compilado da etapa anterior
COPY --from=builder /app/main .

# Porta em que a aplicação irá rodar
EXPOSE 8081

# Comando para executar a aplicação
CMD ["./main"]
