#!/bin/bash

# Загружаем переменные из .env
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo "❌ Файл .env не найден!"
    exit 1
fi

# Формируем строку подключения к БД
DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

echo $DB_URL

# Проверяем, установлен ли golang-migrate
if ! command -v migrate &> /dev/null
then
    echo "❌ golang-migrate не установлен! Установи его: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
fi

# Выполняем миграции
echo "Migration started..."
migrate -database "$DB_URL" -path ./migrations up

if [ $? -eq 0 ]; then
    echo "Migration completed successfully!"
else
    echo "Migration failed!"
    exit 1
fi
