package store

import "github.com/Laura-2950/API-Go-Products.git/API-Go-Products/internal/domain"

type StoreInterface interface {
	Read(id int) (*domain.Product, error)
	ReadAll() ([]domain.Product, error)
	Create(product domain.Product) (*domain.Product, error)
	Exist(codeValue string) bool
}