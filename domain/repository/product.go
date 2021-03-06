package repository

import (
	"github.com/koutsoumposval/ddd-playground-golang/domain/entity"
)

// IProductRepository contract for product repository
type IProductRepository interface {
	Get(id int64) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Save(p *entity.Product) error
}
