package service

import (
	"crud-golang/domain"
)

type ProductService struct {
	pr domain.ProductRepository
}

func NewProductService(pr domain.ProductRepository) domain.ProductService {
	return &ProductService{
		pr: pr,
	}
}

func (ps *ProductService) Create(product *domain.Product) error {
	if err := ps.pr.Create(product); err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) GetAll(page, size int, name string) ([]domain.Product, error) {
	products, err := ps.pr.FindAll(page, size)
	if err != nil {
		return nil, err
	}
	return products, nil
}
