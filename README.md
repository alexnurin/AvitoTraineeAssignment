# Avito Trainee Assignment

Это тестовое задание для отбора на стажировку Авито backend 2024

## Предварительные требования

Перед началом работы убедитесь, что у вас установлены:

- Go (версия 1.15 или выше)
- PostgreSQL (версия 12 или выше)
- Пакетный менеджер для Go (`go mod`)

## Запуск приложения

### Клонирование репозитория:

```
git clone https://github.com/alexnurin/AvitoTraineeAssignment.git
cd AvitoTraineeAssignment
```
_Далее все команды выполняются из корневой директории проекта_

### Запуск контейнера PostgreSQL

Убедитесь что стоит докер

```
docker-compose -f deployments/docker-compose.yml up -d
```

### Установка зависимостей Go:

```
go mod tidy
```

### Запуск приложения:

```
go run cmd/AvitoTraineeAssignment/main.go
```

При успешном запуске в консоль будет выведено несколько сообщений вида:

```
[GIN-debug] GET    /user_banner --> github.com/alexnurin/.../api.InitializeRoutes.func5 (5 handlers)```
```
### Завершение работы

Для завершения приложения достаточно нажать `CTRL + C`

Для остановки и очистки контейнера PostgreSQL можно воспользоваться командой:

```
docker-compose -f deployments/docker-compose.yml down -v
```