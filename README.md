# Tasks API

REST API для управления списком задач (to-do). Данные хранятся в памяти.

## Запуск

```bash
go run ./cmd/server/main.go
```


## Эндпоинты

| Метод | URL | Описание | Код успеха |
|-------|-----|----------|------------|
| GET | /tasks | Получить все задачи | 200 |
| POST | /tasks | Создать задачу | 201 |
| GET | /tasks/{id} | Получить задачу по ID | 200 |
| PUT | /tasks/{id} | Обновить задачу | 200 |
| DELETE | /tasks/{id} | Удалить задачу | 204 |

## Примеры

Создать задачу
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить молоко","done":false}'
```

Ответ:
```json
{
  "id": 1,
  "title": "Купить молоко",
  "done": false,
  "created_at": "2024-01-01T12:00:00Z"
}
```

Получить все задачи
```bash
curl http://localhost:8080/tasks
```

Ответ:
```json
[
  {
    "id": 1,
    "title": "Купить молоко",
    "done": false,
    "created_at": "2024-01-01T12:00:00Z"
  }
]
```

Получить задачу по ID
```bash
curl http://localhost:8080/tasks/1
```

Ответ:
```json
{
  "id": 1,
  "title": "Купить молоко",
  "done": false,
  "created_at": "2024-01-01T12:00:00Z"
}
```

Обновить задачу
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить хлеб","done":true}'
```

Ответ:
```json
{
  "id": 1,
  "title": "Купить хлеб",
  "done": true,
  "created_at": "2024-01-01T12:00:00Z"
}
```

Обновить несуществующую задачу
```bash
curl -X PUT http://localhost:8080/tasks/999 \
  -H "Content-Type: application/json" \
  -d '{"title":"Тест","done":false}'
```

Ответ (404):
```json
{
  "error": "task with id 999 not found"
}
```

Удалить задачу
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

Ответ: (пусто, статус 204)

Удалить несуществующую задачу
```bash
curl -X DELETE http://localhost:8080/tasks/999
```
Ответ (404):

```json
{
  "error": "task with id 999 not found"
}
```

## Ошибки

Неверный JSON
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":}'
```
Ответ (400):
```json
{
  "error": "Invalid JSON format"
}
```

Нет поля title
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"done":false}'
```

Ответ (400):
```json
{
  "error": "Title is required"
}
```

Задача не найдена
```bash
curl http://localhost:8080/tasks/999
```
Ответ (404):
```json
{
  "error": "Task not found"
}
```

Неподдерживаемый метод
```bash
curl -X PATCH http://localhost:8080/tasks/1
```
Ответ (405):

```json
{
  "error": "Method not allowed"
}
```

## Коды статусов

- 200 - успех (GET, PUT)
- 201 - создано (POST)
- 204 - удалено (DELETE)
- 400 - неверные данные
- 404 - задача не найдена
- 405 - метод не поддерживается
- 500 - внутренняя ошибка