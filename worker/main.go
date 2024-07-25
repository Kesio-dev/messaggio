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
		panic(err)
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

	for {
		message, err := kafkaReader.FetchMessage(context.Background())
		if err != nil {
			log.Printf("Failed to fetch message: %v", err)
			continue
		}

		var msg Message
		err = db.Get(&msg, "SELECT * FROM messages WHERE message = $1", string(message.Value))
		if err != nil {
			log.Printf("Failed to fetch message from DB: %v", err)
			continue
		}

		time.Sleep(2 * time.Second)

		if rand.Float32() < 0.5 {
			msg.Status = "success"
		} else {
			msg.Status = "failed"
		}
		msg.UpdatedAt = time.Now()

		_, err = db.NamedExec(`UPDATE messages SET status = :status, updated_at = :updated_at WHERE id = :id`, msg)
		if err != nil {
			log.Printf("Failed to update message status: %v", err)
			continue
		}

		err = kafkaReader.CommitMessages(context.Background(), message)
		if err != nil {
			log.Printf("Failed to commit message: %v", err)
		}
	}
}
