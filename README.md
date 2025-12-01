Краткая инструкция

1. Запустите NATS:
docker run -p 4222:4222 -ti nats:latest

2. Запустите сервер:
go run cmd/server/server.go

3. Подключитесь и отправьте сообщения:
Подключитесь клиентом к ws://localhost:3000/ws?token=secret&user_id=user1
Вторым клиентом — к ws://localhost:3000/ws?token=secret&user_id=user2
От user1 отправьте:
{
"to": "user2",
"content": "Hello, user2!"
}
