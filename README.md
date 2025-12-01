Краткая инструкция

1. Запустите NATS:
docker run -p 4222:4222 -ti nats:latest

2. Запустите сервер:
go run cmd/server/server.go

3. Подключитесь и отправьте сообщения:
Подключитесь клиентом к ws://localhost:3000/ws?token=secret&user_id=user1
Вторым клиентом — к ws://localhost:3000/ws?token=secret&user_id=user2
От user1 отправьте:
```json
{
"to": "user2",
"content": "Hello, user2!"
}
```

Изначальное задание
Реализовать обработчик WebSocket соединения. (gofiber/websocket)
Обработчик должен уметь:
1. Авторизовывать соединение по токену
2. При получении сообщения от клиента WebSocket-сервер публикует его в очередь (Kafka, NATS, Redis Streams)
3. Consumer получает сообщение и доставляет его конкретным получателям через WebSocket в виде json (структуру сообщения организовать самостоятельно)
4. Добавить описание в README инструкцию по запуску приложения