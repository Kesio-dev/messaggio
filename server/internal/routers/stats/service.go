package stats

import (
	"log"
	"messaggio/internal/repositories"
)

type Service struct {
	messageRepo repositories.MessageRepositoryInterface
}

func NewService(messageRepo repositories.MessageRepositoryInterface) *Service {
	return &Service{
		messageRepo: messageRepo,
	}
}

func (s Service) GetAll() (map[string]int, error) {
	stats, err := s.messageRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return stats, nil
}
