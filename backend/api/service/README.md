# Pacote `service`

## Visão Geral

O pacote `service` contém as definições de serviços e lógica de negócio do projeto, desacopladas de uma struct global. Ele fornece métodos organizados para manipular dados, executar regras de negócio e interagir com os repositórios.

Além disso, o pacote inclui um módulo interno de **DTO (Data Transfer Object)** para encapsular e formatar os dados enviados nas respostas das requisições, convertendo objetos do schema do banco para formatos mais leves e específicos.

---

## Estrutura e Organização

Cada arquivo `.go` representa um serviço específico, relacionado diretamente aos repositórios instanciados externamente nos handlers.

### Arquivos de Serviço Atuais

- `authService.go` gerencia lógica de serviço de autenticação da aplicação, login, token e senhas
- `enterpriseService.go` gerencia lógica de serviço de empresas da aplicação.
- `productService.go` gerencia lógica de serviço de produtos da aplicação.
- `roleService.go` gerencia lógica de serviço de permissões da aplicação.
- `userService.go` gerencia lógica de serviço de usuários da aplicação.
- `wishlistService.go` gerencia lógica de serviço de lista de desejo da aplicação.

---

## ✅ Boas Práticas

- Crie um novo arquivo de serviço sempre que necessário para garantir separação de responsabilidades.
- Nomeie funções e métodos de forma **declarativa**.
- Utilize structs para instanciar serviços com injeção de dependências.
> ⚠️ Em um cenário de escalabilidade, recomenda-se revisar a arquitetura, pois atualmente a criação de instâncias depende de uma função `Get` com atributos fixos.

---

## Testes

Os testes são **unitários** e têm como objetivo verificar o comportamento das funções e métodos individualmente, sem depender dos handlers.

- Biblioteca: `testing` (padrão do Go)
- Pacote de testes: `service_test`
- Estrutura dos arquivos de teste segue o nome dos arquivos de produção correspondentes.

### Arquivos de Teste Planejados

- `auth_unit_test.go`
- `enterprise_unit_test.go`
- `product_unit_test.go`
- `role_unit_test.go`
- `user_unit_test.go`
- `wishlist_unit_test.go`

---

## Exemplos

```go
package service

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo   *repository.PostgresUserRepository
	logger *config.Logger
}

func GetUserService(appCtx *appcontext.AppContext) UserService {
	return UserService{
		repo:   repository.NewPostgresUserRepository(appCtx.DB),
		logger: config.NewLogger("SERVICE - USER"),
	}
}

func (u *UserService) ValidateAndUpdateFields(user *schemas.User, input request.UserRequest) error {
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		hashed, err := utils.GenerateHashPassword(input.Password)
		if err != nil {
			return fmt.Errorf("password hash: %v", err)
		}
		user.Password = hashed
	}
	if input.Cpf != "" && utils.ValidateCPF(input.Cpf) {
		user.Cpf = input.Cpf
	}
	if input.RegisterNumber != "" && len(input.RegisterNumber) == 7 {
		user.RegisterNumber = input.RegisterNumber
	}
	if input.RoleID != 0 {
		user.RoleID = input.RoleID
	}
	if input.EnterpriseID != nil && *input.EnterpriseID != 0 {

		user.EnterpriseID = input.EnterpriseID
	}
	if input.PhotoUrl != nil && *input.PhotoUrl != "" {
		user.PhotoUrl = input.PhotoUrl
	}
	return nil
}

func (u *UserService) Create(ctx *gin.Context, input request.UserRequest) (*dto.UserDTO, error) {
	hashedPassword, err := utils.GenerateHashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := schemas.User{
		RoleID:         input.RoleID,
		EnterpriseID:   input.EnterpriseID,
		Name:           input.Name,
		Email:          input.Email,
		Password:       hashedPassword,
		Cpf:            input.Cpf,
		RegisterNumber: input.RegisterNumber,
		PhotoUrl:       input.PhotoUrl,
	}

	if err = u.repo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) Update(ctx *gin.Context, id uint, input request.UserRequest) (*dto.UserDTO, error) {
	user, err := u.repo.GetInfoUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = u.ValidateAndUpdateFields(&user, input); err != nil {
		return nil, err
	}

	if err = u.repo.UpdateUser(ctx, id, &user); err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) Get(ctx *gin.Context, id uint) (*dto.UserDTO, error) {
	user, err := u.repo.GetInfoUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) GetAll(ctx *gin.Context, itemsPerPage, currentPage uint) (*dto.PaginationDTO, error) {
	users, totalPages, totalItems, err := u.repo.GetUsers(ctx, itemsPerPage, currentPage)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}

	toDTO := func(user schemas.User) *dto.UserDTO {
		return dto.MakeUserOutput(user)
	}

	return dto.MakePaginationDTO(
		users,
		currentPage,
		totalPages,
		totalItems,
		toDTO,
	)
}
```