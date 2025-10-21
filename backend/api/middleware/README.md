# Package `middleware`

## Visão Geral

O pacote `middleware` contém funções auxiliares que atuam entre as requisições HTTP e os handlers da aplicação, fornecendo uma camada de controle, segurança e configuração:

- Proteção e configuração de **headers**
- Verificação de **tokens JWT**
- onfiguração de **CORS**
- Verificação de **roles/permissions**
- (Em desenvolvimento) Consumo e verificação de arquivos de imagem (`.png`/`.jpg`)

---

## Estrutura e Organização

Cada arquivo `.go` neste pacote representa uma responsabilidade clara e isolada de middleware. A separação por arquivo facilita:

- Legibilidade
- Manutenção
- Testabilidade
- Reutilização em múltiplas rotas

### Arquivos atuais:

- `cors.go` Define políticas de CORS         
- `headers.go` Adiciona/configura headers comuns
- `jwt.go` Valida e decodifica tokens JWT
- `permission.go` Controla acesso por **papéis** ou **permissões** 
- `picture.go` (WIP) Middleware para upload/validação de imagem

---

## Boas Práticas

- **Responsabilidade única por arquivo**: sempre que criar um novo middleware, crie um novo arquivo com o mesmo nome para manter a clareza da função.
  
- **Testes isolados**: middlewares devem ser testados separadamente, especialmente os que lidam com autenticação e autorização.

- **Composição de middlewares**: utilize composição (chaining) ao invés de lógica aninhada para manter o código limpo e reutilizável.

  ```go
  r.Use(middleware.CORS)
  r.Use(middleware.JWT)
  r.Use(middleware.Permission("admin"))

  ou

  rg.GET("/users",
	middleware.JWTMiddleware(appCtx, logger),
	middleware.AdminPermission(appCtx, logger),
	handler.GetUsersHandler(
	appCtx, config.NewLogger("GET - USERS"),
  ))
