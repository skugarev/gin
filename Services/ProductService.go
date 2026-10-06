package Services

import (
	"main/Models"
	"main/Repository"
)

type ProductService struct {
	ProductRepository Repository.ProductRepository
}

func ProvideProductService(p Repository.ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() []Models.Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) Find(walletId uint) Models.Product {
	return p.ProductRepository.Find(walletId)
}

func (p *ProductService) Save(valletId uint, operationType string, amount uint) Models.Product {
	return p.ProductRepository.Save(valletId, operationType, amount)
}
