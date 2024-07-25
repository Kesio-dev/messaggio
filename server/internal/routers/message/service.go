package message

import (
	"log"
	"messaggio/internal/kafka"
	"messaggio/internal/repositories"
	"strconv"
)

type Service struct {
	messageRepo   repositories.MessageRepositoryInterface
	kafkaProducer *kafka.Producer
}

func NewService(messageRepo repositories.MessageRepositoryInterface, kafkaProducer *kafka.Producer) *Service {
	if kafkaProducer == nil {
		log.Fatalf("Kafka producer is nil")
	}
	return &Service{
		messageRepo:   messageRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (s *Service) SaveMessage(message string) error {
	id, err := s.messageRepo.Create(message)
	if err != nil {
		return err
	}

	err = s.kafkaProducer.SendMessage(strconv.Itoa(id), message)
	return err
}
