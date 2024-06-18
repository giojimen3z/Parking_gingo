# Etapa de construcción
FROM golang:1.21.4-alpine AS builder

# Instalar las dependencias necesarias para CGO y cross-compilation
RUN apk add --no-cache gcc musl-dev linux-headers

WORKDIR /app

# Copia los archivos de dependencias primero para aprovechar la caché de Docker
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copia el resto de los archivos
COPY . .

# Habilitar CGO y establecer GOARCH
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

# Compilar el ejecutable
RUN go build -o /parqueadero-api main.go

# Etapa final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /parqueadero-api .
EXPOSE 8080
CMD ["./parqueadero-api"]
