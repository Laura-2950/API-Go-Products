package store

import "github.com/Laura-2950/API-Go-Products/API-Go-Products/internal/domain"

type StoreInterface interface {
	Read(id int) (*domain.Product, error)
	ReadAll() ([]domain.Product, error)
	Create(product domain.Product) (*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(id int) error
	Exist(codeValue string) bool
	ExistId(id int) bool
}
