package product

import "github.com/Laura-2950/API-Go-Products/API-Go-Products/internal/domain"

type IService interface {
	GetProductBy(id int) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	CreateNewProducts(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetProductBy(id int) (*domain.Product, error) {
	product, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) GetAllProducts() ([]domain.Product, error) {
	products, err := s.Repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Service) CreateNewProducts(product *domain.Product) (*domain.Product, error) {
	product, err := s.Repository.CreateNewProducts(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) DeleteProduct(id int) error {
	err := s.Repository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
