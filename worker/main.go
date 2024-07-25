package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/segmentio/kafka-go"
)

type Message struct {
	ID        int       `db:"id"`
	Message   string    `db:"message"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func main() {
	db, err := sqlx.Connect("pgx", "postgresql://user:password@db:5432/messages_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	addr := "kafka:9092"
	topic := "messages"
	groupID := "message-workers"
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{addr},
		Topic:   topic,
		GroupID: groupID,
	})
	defer kafkaReader.Close()

	log.Println("Начало обработки сообщений...")

	for {
		message, err := kafkaReader.FetchMessage(context.Background())
		if err != nil {
			log.Printf("Не удалось получить сообщение: %v", err)
			continue
		}

		log.Printf("Получено сообщение: %s с ключом: %s", string(message.Value), string(message.Key))

		var msg Message
		err = db.Get(&msg, "SELECT * FROM messages WHERE id = $1", string(message.Key))
		if err != nil {
			log.Printf("Не удалось получить сообщение из БД: %v", err)
			continue
		}

		log.Printf("Сообщение из БД: %v", msg)

		time.Sleep(2 * time.Second)

		if rand.Float32() < 0.5 {
			msg.Status = "success"
			log.Println("Статус сообщения обновлен на 'success'")
		} else {
			msg.Status = "failed"
			log.Println("Статус сообщения обновлен на 'failed'")
		}
		msg.UpdatedAt = time.Now()

		_, err = db.NamedExec(`UPDATE messages SET status = :status, updated_at = :updated_at WHERE id = :id`, msg)
		if err != nil {
			log.Printf("Не удалось обновить статус сообщения: %v", err)
			continue
		}

		err = kafkaReader.CommitMessages(context.Background(), message)
		if err != nil {
			log.Printf("Не удалось зафиксировать сообщение: %v", err)
		} else {
			log.Printf("Сообщение зафиксировано: %s", string(message.Value))
		}
	}
}
