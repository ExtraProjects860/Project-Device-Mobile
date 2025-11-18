# Package `routes`

## Visão Geral

O pacote `routes` é responsável por centralizar e registrar todas as rotas da aplicação, conectando as URLs aos seus respectivos handlers. Ele atua como a camada de entrada da aplicação HTTP, orquestrando middlewares, validações e organização de grupos de endpoints.

Principais responsabilidades:

- Definir **grupos de rotas** organizados por domínio de negócio (ex: usuários, produtos, autenticação)
- Registrar rotas na **função principal do router**
- Associar middlewares globais e específicos
- Facilitar a manutenção e expansão das rotas

---

## Estrutura e Organização

Cada arquivo `.go` representa um grupo de rotas relacionadas a um handler específico, com funções que as registram em um `RouterGroup` do Gin. A entrada principal é feita nos arquivos `router.go` e `routes.go`.

### Arquivos atuais:

- `authRoutes.go` Rotas relacionadas à autenticação (login, refresh, etc)
- `enterpriseRoutes.go` Rotas da entidade `Enterprise`
- `healthCheckRoutes.go` Rotas para verificação de saúde da API (`/health`)
- `productRoutes.go` Rotas relacionadas à entidade `Product`
- `roleRoutes.go` Rotas de controle de acesso (roles e permissões)
- `router.go` Configurações principais do servidor e inicialização do `gin.Engine`
- `routes.go` Entrada de todas as rotas da API agrupadas por versão (`/api/v1`)
- `swaggerRoutes.go` Rotas de documentação Swagger
- `userRoutes.go` Rotas de usuários
- `wishlistRoutes.go` Rotas de wishlist

---

## Boas Práticas

- **Isolamento de responsabilidade**: Cada grupo de rotas deve estar em um arquivo separado para melhorar a legibilidade e facilitar testes e manutenção.

- **Middlewares globais e locais**:
  - Use `router.Use(...)` para middlewares globais como CORS, segurança, etc.
  - Use middlewares específicos em rotas sensíveis (como JWT e permissões).

- **Evitar lógica de negócio**: As rotas devem apenas **chamar handlers**, que por sua vez chamam os **serviços**. Evite qualquer tipo de lógica dentro das funções de rota.

- **Nomenclatura clara**: Use o prefixo `registerXYZRoutes(...)` para nomear funções de registro de rotas e manter a consistência.

---

## Observação sobre Escalabilidade

> ⚠️ À medida que o número de rotas crescer significativamente, pode ser necessário repensar a estratégia de registro. Uma alternativa é utilizar um **map (hash map)** para associar caminhos a handlers dinamicamente.

### Exemplo (em pseudocódigo):

```go
var routeMap = map[string]gin.HandlerFunc{
  "GET:/api/v1/users": handler.GetUsers,
  "POST:/api/v1/products": handler.CreateProduct,
}
```

> ⚠️ Atenção:
Cuidado com uso excessivo de memória ao popular muitos caminhos no map.
Evite slices ou arrays para registrar rotas dinamicamente, pois:
Slices: podem ter busca lenta conforme crescem.
Arrays: demandariam reestruturação constante e realocação de memória.
Avalie o uso de rotas hierárquicas e auto-registro por metadata/reflection se a aplicação continuar escalando.

## Exemplos

### `router.go`

```go
package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func configureNetwork(router *gin.Engine) {
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{})
}

func InitializeRouter(appCtx *appcontext.AppContext) *gin.Engine {
	router := gin.Default()

	configureNetwork(router)

	middleware.SecurityHeaders(router)
	middleware.SetCors(router)

	InitHealthCheckRoutes(router, appCtx)
	InitRoutesApiV1(router, appCtx)
	InitSwaggerRoute(router)

	return router
}
```

### `routes.go`

```go
package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/gin-gonic/gin"
)

func InitRoutesApiV1(router *gin.Engine, appCtx *appcontext.AppContext) {
	apiV1 := router.Group("/api/v1")

	registerUserRoutes(apiV1, appCtx)
	registerProductRoutes(apiV1, appCtx)
	registerAuthRoutes(apiV1, appCtx)
	registerWishListRoutes(apiV1, appCtx)
	registerEnterpriseRoutes(apiV1, appCtx)
	registerRoleRoutes(apiV1, appCtx)
}
```

### `userRoutes.go`

```go
package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	logger := config.NewLogger("MIDDLEWARE")

	{
		rg.GET("/users",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.GetUsersHandler(
				appCtx, config.NewLogger("GET - USERS"),
			))

		rg.GET("/user",
			middleware.JWTMiddleware(appCtx, logger),
			handler.GetInfoUserHandler(
				appCtx, config.NewLogger("GET - USER"),
			))

		rg.POST("/user",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.CreateUserHandler(
				appCtx, config.NewLogger("POST - USERS"),
			))

		rg.PATCH("/user",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.UpdateUserHandler(
				appCtx, config.NewLogger("PATCH - USERS"),
			))
	}
}
```