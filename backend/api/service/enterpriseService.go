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

type EnterpriseService struct {
	repo   *repository.PostgresEnterpriseRepository
	logger *config.Logger
}

func GetEnterpriseService(appCtx *appcontext.AppContext) EnterpriseService {
	return EnterpriseService{
		repo:   repository.NewPostgresEnterpriseRepository(appCtx.DB),
		logger: config.NewLogger("SERVICE - ENTERPRISE"),
	}
}

func (e *EnterpriseService) ValidateAndUpdateFields(enterprise *schemas.Enterprise, input request.EnterpriseRequest) {
	if input.Name != "" {
		enterprise.Name = input.Name
	}
}

func (e *EnterpriseService) Create(ctx *gin.Context, input request.EnterpriseRequest) (*dto.EnterpriseDTO, error) {
	enterprise := schemas.Enterprise{
		Name: input.Name,
	}

	if err := e.repo.CreateEnterprise(ctx, &enterprise); err != nil {
		return nil, err
	}

	return dto.MakeEnterpriseOutput(enterprise), nil
}

func (e *EnterpriseService) Update(ctx *gin.Context, enterpriseID uint, input request.EnterpriseRequest) (*dto.EnterpriseDTO, error) {
	enterprise, err := e.repo.GetEnterprise(ctx, enterpriseID)
	if err != nil {
		return nil, err
	}

	e.ValidateAndUpdateFields(&enterprise, input)

	if err = e.repo.UpdateEnterprise(ctx, enterpriseID, &enterprise); err != nil {
		return nil, err
	}

	return dto.MakeEnterpriseOutput(enterprise), nil
}

// TODO ficar esperto no frontend, pois vai ter que usar select, porém com paginação
func (e *EnterpriseService) GetAll(ctx *gin.Context, paginationSearch request.PaginationSearch) (*dto.PaginationDTO, error) {
	enterprises, totalPages, totalItems, err := e.repo.GetEnterprises(
		ctx, 
		paginationSearch,
	)
	if err != nil {
		e.logger.Error(err.Error())
		return nil, err
	}

	toDTO := func(enterprise schemas.Enterprise) *dto.EnterpriseDTO {
		return dto.MakeEnterpriseOutput(enterprise)
	}

	return dto.MakePaginationDTO(
		enterprises,
		paginationSearch.CurrentPage,
		totalPages,
		totalItems,
		toDTO,
	)
}
