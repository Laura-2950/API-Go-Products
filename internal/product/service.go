package product

import "github.com/Laura-2950/API-Go-Products.git/API-Go-Products/internal/domain"

type IService interface {
	GetProductBy(id int) (*domain.Product, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetProductBy(id int) (*domain.Product, error) {
	p, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}