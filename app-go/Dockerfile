# ==========================================
# STAGE 1: Builder (Compilação)
# ==========================================
FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# CGO_ENABLED=0 gera um binário estático (independe de C libraries no Alpine).
# -ldflags="-s -w" remove símbolos de debug, diminuindo o tamanho final do binário.
# -o korp-app define o nome exato do arquivo de saída.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o korp-app .

# ==========================================
# STAGE 2: Runtime (Execução)
# ==========================================
FROM alpine:3.24 AS runtime

# Instalação de certificados CA  e criação do usuário não-root 
RUN apk --no-cache add ca-certificates && \
    adduser -D -u 1000 appuser

WORKDIR /app

# Copiar APENAS o binário compilado do stage 'builder' para o stage 'runtime'.
# A imagem final não terá o código fonte, nem o Go, nem as dependências.
COPY --from=builder /build/korp-app .

# Mudamos para o usuário não-root. A partir daqui, nada roda como root.
USER appuser

EXPOSE 8080

ENTRYPOINT ["./korp-app"]