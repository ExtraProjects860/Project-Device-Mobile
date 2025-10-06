package service

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/gin-gonic/gin"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return ProductService{repo: repo}
}

func (p *ProductService) GetAll(ctx *gin.Context, itemsPerPage, currentPage uint) (*dto.PaginationDTO, error) {
	products, totalPages, totalItems, err := p.repo.GetProducts(ctx, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	toDTO := func(product schemas.Product) *dto.ProductDTO {
		return dto.MakeProductOutput(product)
	}

	return dto.MakePaginationDTO(
		products,
		currentPage,
		totalPages,
		totalItems,
		toDTO,
	)
}
