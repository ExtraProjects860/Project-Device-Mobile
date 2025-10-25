package service

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/gin-gonic/gin"
)

type ProductService struct {
	repo   *repository.PostgresProductRepository
	logger *config.Logger
}

func GetProductService(appCtx *appcontext.AppContext) ProductService {
	return ProductService{
		repo:   repository.NewPostgresProductRepository(appCtx.DB),
		logger: config.NewLogger("SERVICE - PRODUCT"),
	}
}

func (p *ProductService) ValidateAndUpdateFields(product *schemas.Product, input request.ProductRequest) error {
	return nil
}

func (p *ProductService) Create(ctx *gin.Context, input request.ProductRequest) (*dto.ProductDTO, error) {
	product := schemas.Product{
		Name:               input.Name,
		Description:        input.Description,
		Value:              input.Value,
		Quantity:           input.Quantity,
		IsPromotionAvaible: input.IsPromotionAvaible,
		Discount:           input.Discount,
		PhotoUrl:           input.PhotoUrl,
		IsAvaible:          input.IsAvaible,
	}

	if err := p.repo.CreateProduct(ctx, &product); err != nil {
		return nil, err
	}

	return dto.MakeProductOutput(product), nil
}

func (p *ProductService) Update(ctx *gin.Context, id uint, input request.ProductRequest) (*dto.ProductDTO, error) {
	product, err := p.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = p.ValidateAndUpdateFields(&product, input); err != nil {
		return nil, err
	}

	if err = p.repo.UpdateProducts(ctx, id, &product); err != nil {
		return nil, err
	}

	return dto.MakeProductOutput(product), nil
}

func (p *ProductService) GetAll(ctx *gin.Context, paginationSearch request.PaginationSearch) (*dto.PaginationDTO, error) {
	products, totalPages, totalItems, err := p.repo.GetProducts(
		ctx,
		paginationSearch,
	)
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}

	toDTO := func(product schemas.Product) *dto.ProductDTO {
		return dto.MakeProductOutput(product)
	}

	return dto.MakePaginationDTO(
		products,
		paginationSearch.CurrentPage,
		totalPages,
		totalItems,
		toDTO,
	)
}
