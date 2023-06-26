package product

import (
	"fmt"

	"github.com/Laura-2950/API-Go-Products/API-Go-Products/internal/domain"
	"github.com/Laura-2950/API-Go-Products/API-Go-Products/pkg/store"
	"github.com/Laura-2950/API-Go-Products/API-Go-Products/pkg/web"
)

type IRepository interface {
	GetByID(id int) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	CreateNewProducts(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
	/*Store(name, productType string, count int, price float64) (*domain.Product, error)---CreateNewProducts
	Update(product domain.Product) (*domain.Product, error)*/
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) GetByID(id int) (*domain.Product, error) {
	prod, err := r.Storage.Read(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("product_id %d not found", id))
	}
	return prod, nil
}

func (r *Repository) GetAllProducts() ([]domain.Product, error) {
	prod, err := r.Storage.ReadAll()
	if err != nil {
		return nil, web.NewInternalServerApiError("unexpected error")
	}
	return prod, nil
}

func (r *Repository) CreateNewProducts(product *domain.Product) (*domain.Product, error) {
	if r.Storage.Exist(product.CodeValue) {
		return nil, web.NewBadRequestApiError("existing product")
	}
	prod, err := r.Storage.Create(*product)
	if err != nil {
		return nil, web.NewInternalServerApiError("unexpected error")
	}
	return prod, nil
}

func (r *Repository) DeleteProduct(id int) error {
	if r.Storage.ExistId(id) {
		return web.NewBadRequestApiError(fmt.Sprintf("nonexistent product with id %d.", id))
	}
	err := r.Storage.Delete(id)
	if err != nil {
		return web.NewNotFoundApiError(fmt.Sprintf("product_id %d not found", id))
	}
	return nil
}
