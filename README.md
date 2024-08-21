# go-todo

CLI tool для работы со списком задач. Сохраняет список в файле в формате JSON.

----

## Как пользоваться

```bash
# Добавить задачу
go run cmd/todo/main.go -task "Погладить кота"

# Посмотреть список задач
go run cmd/todo/main.go -list

# Завершить задачу
go run cmd/todo/main.go -complete 1
```

## Запуск тестов

```bash
go test -v
```