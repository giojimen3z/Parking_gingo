# Dockerfile
FROM golang:1.21.4-alpine

# Instalar las dependencias necesarias para CGO
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copia los archivos de dependencias primero para aprovechar la caché de Docker
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copia el resto de los archivos
COPY . .

# Establecer el directorio de trabajo para la compilación
WORKDIR /app

# Habilitar CGO y construir el ejecutable
ENV CGO_ENABLED=1
RUN go build -o /parqueadero-api main.go

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar el binario
CMD ["/parqueadero-api"]
