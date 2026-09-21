# Estágio de compilação (Build)
FROM golang:1.26-trixie AS builder

WORKDIR /app

# Baixa as dependências do módulo Go
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte e compila binário estático e otimizado
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Estágio final (Runtime enxuto para Cloud Run Free Tier: ~20MB)
FROM alpine:3.20

# Certificados CA para chamadas HTTPS externas e timezone
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copia apenas o binário compilado e os templates HTML necessários
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates

# O Cloud Run usa a porta 8080 por padrão via variável PORT
EXPOSE 8080

CMD ["./main"]
