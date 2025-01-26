# Imagen base de Go
FROM golang:1.20-alpine

# Instalar dependencias del sistema
RUN apk add --no-cache git

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto definido en el .env (por defecto 8080)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
