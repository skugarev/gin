package Mapper

import (
	"main/DTO"
	"main/Models"
)

func ToProductDTO(product Models.Product) DTO.ProductDTO {
	return DTO.ProductDTO{ValletId: product.ValletId, OperationType: product.OperationType, Amount: product.Amount}
}

func ToProductDTOs(products []Models.Product) []DTO.ProductDTO {
	productdtos := make([]DTO.ProductDTO, len(products))

	for i, item := range products {
		productdtos[i] = ToProductDTO(item)
	}

	return productdtos
}
