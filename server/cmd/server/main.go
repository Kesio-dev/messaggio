package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"messaggio/internal/kafka"
	"messaggio/internal/repositories"
	"messaggio/internal/routers/message"
	"messaggio/internal/routers/stats"
)

func main() {
	app := fiber.New()

	db, err := sqlx.Connect("pgx", "postgresql://user:password@db:5432/messages_db?sslmode=disable")
	if err != nil {
		panic(err)
	}

	addr := "kafka:9092"
	topic := "messages"
	kafkaProducer := kafka.NewProducer(addr, topic)
	if kafkaProducer == nil {
		panic("Failed to create Kafka producer")
	}
	defer kafkaProducer.Close()

	messageRepo := repositories.NewMessageRepository(db)

	message.SetupRouter(app, messageRepo, kafkaProducer)
	stats.SetupRouter(app, messageRepo)

	err = app.Listen(":5001")
	if err != nil {
		panic(err)
	}
}
