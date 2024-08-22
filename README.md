# ✅ go-todo

CLI tool для работы со списком задач. Сохраняет список в файле в формате JSON.

----

## Настройка

Имя файла для хранения списка задач можно установить через переменную окружения:

```bash
export GO_TODO_FILENAME=~/.go_todo.json
```

По умолчанию, файл будет сохраняться в текущей директории под именем `.todo.json`.

## Как пользоваться

```bash
# Добавить задачу
go run cmd/todo/main.go -task "Погладить кота"

# Посмотреть список задач
go run cmd/todo/main.go -list

# Завершить задачу
go run cmd/todo/main.go -complete 1

# Получить справку
go run cmd/todo/main.go -h
```

## Запуск тестов

```bash
go test -v
go test -v ./cmd/todo
```

----

## Документация к пакетам

- [flag](https://pkg.go.dev/flag)
