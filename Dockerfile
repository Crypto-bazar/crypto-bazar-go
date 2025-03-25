# Используем официальный образ Go для сборки
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Сборка исполняемого файла
RUN go build -o main ./cmd/main.go

# Используем минимальный образ для финального контейнера
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарник из предыдущего контейнера
COPY --from=builder /app/main .

# Копируем директорию загрузок и миграции
COPY --from=builder /app/uploads ./uploads
COPY --from=builder /app/migrations ./migrations

# Устанавливаем переменные окружения
ENV DB_HOST=postgres_db
ENV DB_USER=root
ENV DB_PASSWORD=aiosghjdaogyiphdioUAHSDIUhasiud1erASD2q231d
ENV DB_NAME=prod
ENV DB_PORT=7432
ENV SERVER_PORT=8080
ENV UPLOAD_DIR=./uploads
ENV ETHEREUM_NODE_URL=ws://hardhat:8545
ENV CONTRACT_ADDRESS=0x5FbDB2315678afecb367f032d93F642f64180aa3

# Создаем директорию для загрузок
RUN mkdir -p ${UPLOAD_DIR}

# Выполнение миграций перед запуском приложения
CMD ./main migrate-up && ./main
