package mysql

import (
	"context"

	"github.com/kumin/GolangGraphQL/entities"
	"github.com/kumin/GolangGraphQL/repos"
	"gorm.io/gorm"
)

var _ repos.ProductRepo = &ProductMysqlRepo{}

type ProductMysqlRepo struct {
	db *gorm.DB
}

func NewProductMysqlRepo(
	db *gorm.DB,
) *ProductMysqlRepo {
	return &ProductMysqlRepo{
		db: db,
	}
}

func (p *ProductMysqlRepo) CreateProduct(
	ctx context.Context,
	prod *entities.Product,
) (*entities.Product, error) {
	if err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(prod).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return prod, nil
}

func (p *ProductMysqlRepo) ListProducts(
	ctx context.Context,
	page, limit int,
) ([]*entities.Product, error) {
	offset := (page - 1) * limit
	var prods []*entities.Product
	if err := p.db.WithContext(ctx).
		Model(&entities.Product{}).
		Joins("Properties").
		Offset(offset).
		Limit(limit).
		Find(&prods).Error; err != nil {
		return nil, err
	}

	return prods, nil
}

func (p *ProductMysqlRepo) AllProducts(
	ctx context.Context,
) ([]*entities.Product, error) {
	var prods []*entities.Product
	if err := p.db.WithContext(ctx).
		Model(&entities.Product{}).
		Joins("Properties").
		Scan(&prods).Error; err != nil {
		return nil, err
	}

	return prods, nil
}
