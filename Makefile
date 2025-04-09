include .env
export $(shell sed 's/=.*//' .env)

APP_NAME=main.go
MIGRATIONS_DIR=./migrations
DSN=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

dev:
	air -c .air.toml

# Установка зависимостей
deps:
	go mod download

# Сборка проекта
build:
	go build -o $(APP_NAME) ./cmd/$(APP_NAME)

# Запуск проекта
run:
	go run ./cmd/$(APP_NAME)

# Тестирование
test:
	go test -v ./...

# Очистка
clean:
	go clean
	rm -f $(APP_NAME)

# Создание новой миграции
migrate-create:
	@read -p "Введите название миграции: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $${name}

# Применение миграций
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DSN)" up

# Откат миграций
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DSN)" down

# Откат на 1 миграцию
migrate-down-1:
	migrate -path $(MIGRATIONS_DIR) -database "$(DSN)" down 1

# Применение на 1 миграцию
migrate-up-1:
	migrate -path $(MIGRATIONS_DIR) -database "$(DSN)" up 1

# Проверка версии миграций
migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DSN)" version

# Установка golang-migrate (требуется brew)
install-migrate:
	brew install golang-migrate

# Форматирование кода
fmt:
	go fmt ./...

# Линтинг
lint:
	golangci-lint run ./...

# Запуск всех проверок (тесты, линтинг, форматирование)
check: test lint fmt

.PHONY: deps build run test clean migrate-create migrate-up migrate-down migrate-down-1 migrate-up-1 migrate-version install-migrate fmt lint check
