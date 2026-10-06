package Repository

import (
	"main/Models"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() []Models.Product {
	var products []Models.Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) Find(walletId uint) Models.Product {
	var product Models.Product
	result := p.DB.Where("vallet_id = ?", walletId).Find(&product)

	if result.RowsAffected == 0 {
		product = Models.Product{}
	}

	return product
}

func (p *ProductRepository) Save(valletId uint, operationType string, amount uint) Models.Product {

	var product Models.Product
	result := p.DB.Where("vallet_id = ?", valletId).Find(&product)

	product = Models.Product{
		ValletId:      valletId,
		OperationType: operationType,
		Amount:        amount,
	}

	if result.RowsAffected == 0 {
		p.DB.Create(&product)
	} else {
		result.Update(&product)
	}

	return product
}
