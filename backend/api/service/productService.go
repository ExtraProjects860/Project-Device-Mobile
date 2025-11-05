package service

import (
	"context"
	"mime/multipart"

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

func (p *ProductService) ValidateAndUpdateFields(product *schemas.Product, input request.ProductRequest) {
	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Value > 0 {
		product.Value = input.Value
	}
	if input.Quantity >= 0 {
		product.Quantity = input.Quantity
	}
	if input.IsPromotionAvaible != nil {
		product.IsPromotionAvaible = input.IsPromotionAvaible
	}
	if input.Discount != nil && *input.Discount >= 0 {
		product.Discount = input.Discount
	}
	if input.IsAvaible != nil {
		product.IsAvaible = input.IsAvaible
	}
}

func (p *ProductService) Create(ctx *gin.Context, imageService ImageService, input request.ProductRequest) (*dto.ProductDTO, error) {
	product := schemas.Product{
		Name:               input.Name,
		Description:        input.Description,
		Value:              input.Value,
		Quantity:           input.Quantity,
		IsPromotionAvaible: input.IsPromotionAvaible,
		Discount:           input.Discount,
		IsAvaible:          input.IsAvaible,
	}

	if err := p.repo.CreateProduct(ctx, &product); err != nil {
		p.logger.Warningf("Failed to create product in database: %v", err)
		return nil, err
	}

	file, err := request.GetFileHeader(ctx, "image")
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}

	go func(product *schemas.Product, file *multipart.FileHeader) {
		secureURL, _, err := imageService.UploadImage(file, FolderUser)
		if err != nil {
			p.logger.Errorf("Failed during image upload process: %v", err)
			return
		}

		if secureURL != nil {
			product.PhotoUrl = secureURL
			if err := p.repo.UpdateProduct(
				context.Background(),
				product.ID,
				product,
			); err != nil {
				p.logger.Errorf("Failed to create photo for product %d: %v", product.ID, err)
				return
			}
		}
	}(&product, file)

	return dto.MakeProductOutput(product), nil
}

func (p *ProductService) Update(ctx *gin.Context, imageService ImageService, id uint, input request.ProductRequest) (*dto.ProductDTO, error) {
	product, err := p.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	p.ValidateAndUpdateFields(&product, input)

	if err = p.repo.UpdateProduct(ctx, id, &product); err != nil {
		p.logger.Errorf("Failed to update product in database: %v", err)
		return nil, err
	}

	file, err := request.GetFileHeader(ctx, "image")
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}

	go func(product *schemas.Product, file *multipart.FileHeader) {
		secureURL, _, err := imageService.UploadImage(file, FolderUser)
		if err != nil {
			p.logger.Errorf("Failed during image upload process: %v", err)
			return
		}

		if secureURL != nil {
			product.PhotoUrl = secureURL
			if err := p.repo.UpdateProduct(
				context.Background(),
				product.ID,
				product,
			); err != nil {
				p.logger.Errorf("Failed to update photo for product %d: %v", product.ID, err)
				return
			}
		}
	}(&product, file)

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
