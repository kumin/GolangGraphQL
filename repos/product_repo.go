package repos

import (
	"context"

	"github.com/kumin/GolangGraphQL/entities"
)

type ProductRepo interface {
	CreateProduct(ctx context.Context, prod *entities.Product) (*entities.Product, error)
	ListProducts(ctx context.Context, page, limit int) ([]*entities.Product, error)
	AllProducts(ctx context.Context) ([]*entities.Product, error)
}
