package product

import (
	"fmt"

	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/internal/domain"
	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/pkg/store"
	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/pkg/web"
)

type IRepository interface {
	GetByID(id int) (*domain.Product, error)
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