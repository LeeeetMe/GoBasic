package services

import (
	"IrisProduct/datamodels"
	"IrisProduct/repository"
)

type IProductService interface {
	SearchProduct(int64) (*datamodels.Product, error)
	Insert(*datamodels.Product) (int64, error)
	Update(int64, *datamodels.Product) error
	Delete(int64) bool
	SelectAll() ([]*datamodels.Product, error)
}

type ProductService struct {
	repo repository.IProduct
}

func (p *ProductService) SearchProduct(id int64) (*datamodels.Product, error) {
	return p.repo.SearchProduct(id)
}

func (p *ProductService) Insert(product *datamodels.Product) (int64, error) {
	return p.repo.Insert(product)
}

func (p *ProductService) Update(id int64, product *datamodels.Product) error {
	return p.repo.Update(id, product)
}

func (p *ProductService) Delete(id int64) bool {
	return p.repo.Delete(id)
}

func (p *ProductService) SelectAll() ([]*datamodels.Product, error) {
	return p.repo.SelectAll()
}

func NewProductService(r repository.IProduct) IProductService {
	return &ProductService{
		repo: r,
	}
}
