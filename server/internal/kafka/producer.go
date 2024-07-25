package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(addr, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{writer: writer}
}

func (p *Producer) SendMessage(key, message string) error {
	err := p.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}
	return nil
}

func (p *Producer) Close() {
	if p != nil && p.writer != nil {
		p.writer.Close()
	}
}
