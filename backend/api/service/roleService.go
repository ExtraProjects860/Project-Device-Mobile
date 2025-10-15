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

type RoleService struct {
	repo   *repository.PostgresRoleRepository
	logger *config.Logger
}

func GetRoleService(appCtx *appcontext.AppContext) RoleService {
	return RoleService{
		repository.NewPostgresRoleRepository(appCtx.DB),
		config.NewLogger("SERVICE - ROLE"),
	}
}

func (ro *RoleService) ValidateAndUpdateFields(role *schemas.Role, input request.RoleRequest) {
	if input.Name != "" {
		role.Name = input.Name
	}
}

func (ro *RoleService) Create(ctx *gin.Context, input request.RoleRequest) (*dto.RoleDTO, error) {
	role := schemas.Role{
		Name: input.Name,
	}

	if err := ro.repo.CreateRole(ctx, &role); err != nil {
		return nil, err
	}

	return dto.MakeRoleOutput(role), nil
}

func (ro *RoleService) Update(ctx *gin.Context, roleID uint, input request.RoleRequest) (*dto.RoleDTO, error) {
	role, err := ro.repo.GetRole(ctx, roleID)
	if err != nil {
		return nil, err
	}

	ro.ValidateAndUpdateFields(&role, input)

	if err = ro.repo.UpdateRole(ctx, roleID, &role); err != nil {
		return nil, err
	}

	return dto.MakeRoleOutput(role), nil
}

// TODO ficar esperto no frontend, pois vai ter que usar select, porém com paginação
func (ro *RoleService) GetAll(ctx *gin.Context, itemsPerPage, currentPage uint) (*dto.PaginationDTO, error) {
	roles, totalPages, totalItems, err := ro.repo.GetRoles(ctx, itemsPerPage, currentPage)
	if err != nil {
		ro.logger.Error(err.Error())
		return nil, err
	}

	toDTO := func(role schemas.Role) *dto.RoleDTO {
		return dto.MakeRoleOutput(role)
	}

	return dto.MakePaginationDTO(
		roles,
		currentPage,
		totalPages,
		totalItems,
		toDTO,
	)
}
