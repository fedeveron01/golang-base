# Usa la imagen oficial de Golang como base
FROM golang:1.20

# Establece el directorio de trabajo en /app
WORKDIR /

# Copia el código fuente de tu aplicación al contenedor
COPY . .

# Compila la aplicación
RUN go build -o main

# Expone el puerto en el que la aplicación escucha
EXPOSE 8080

# Comando para ejecutar la aplicación cuando se inicie el contenedor
CMD ["./main"]