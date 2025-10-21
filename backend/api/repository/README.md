# Package `repository`

## Visão Geral

O pacote `repository` contém funções e structs que gerenciam diretamente os dados com o banco de dados. Seu foco é fornecer uma camada de acesso aos dados sem conter regras de negócio.

Principais responsabilidades:

- Gerenciar **operações CRUD** diretamente com o banco
- Fornecer uma **camada de abstração** para os dados persistidos
- Utilizar o padrão **DAO (Data Access Object)** para isolar o acesso ao banco
- Manter a **organização e reutilização** de funções comuns (ex: paginação, consultas genéricas)

---

## Estrutura e Organização

Cada arquivo `.go` representa a manipulação de uma entidade específica do banco de dados, normalmente relacionada a um schema. A estrutura também contém arquivos utilitários para comportamentos genéricos, como paginação e funções de acesso genérico.

### Arquivos atuais:

- `auth_postgres.go` Gerencia dados de autenticação (ex: login, tokens)
- `dao.go` Funções genéricas de acesso a dados (ex: `getByID`, `create`, `update`)
- `enterprise_postgres.go` Operações relacionadas à entidade `Enterprise`
- `pagination.go` Estruturas e funções de suporte à paginação de resultados
- `postgres.go` Inicialização e configuração do banco Postgres
- `product_postgres.go` Gerencia dados da entidade `Product`
- `role_postgres.go` Manipulação de permissões e papéis (`Role`)
- `user_postgres.go` Acesso e manipulação de dados da entidade `User`
- `wishlist_postgres.go` Operações sobre listas de desejo (`Wishlist`)

---

## Boas Práticas

- **Isolamento de responsabilidade**: Cada arquivo deve conter apenas operações de acesso a uma entidade específica do banco.
  
- **Reutilização de funções genéricas**: Funções como `getByID`, `create`, `update`, `delete` devem ser mantidas em arquivos utilitários como `dao.go`.

- **Evitar lógica de negócio**: Regras de negócio não devem ser implementadas nesta camada — isso é responsabilidade da camada de serviço.

- **Uso adequado de contexto (`context.Context`)**: Sempre propague o `ctx` nas operações de banco para permitir cancelamento, timeout e trace.

---

## Testes

Os testes são unitários e utilizam mocks para verificar o comportamento das funções de acesso a dados.

- Os testes são feitos com a biblioteca padrão do Go (`testing`)
- Devem ser escritos no pacote `repository_test`
- Cada arquivo de teste está alinhado com o nome do seu respectivo arquivo de produção

### Arquivos de teste planejados:

- `auth_unit_test.go`
- `dao_unit_test.go`
- `enterprise_unit_test.go`
- `pagination_unit_test.go`
- `product_unit_test.go`
- `role_unit_test.go`
- `user_unit_test.go`
- `wishlist_unit_test.go`
- `repository_test.go` Arquivo principal com a função `TestMain`

---

## Exemplos

### `user_postgres.go`

```go
func (r *PostgresUserRepository) GetInfoUser(ctx context.Context, id uint) (schemas.User, error) {
	query := r.db.WithContext(ctx).
		Model(&schemas.User{}).
		Preload("Role").
		Preload("Enterprise")

	user, err := getByID[schemas.User](query, id)
	if err != nil {
		return schemas.User{}, err
	}
	return user, nil
}
```

### `dao.go`

```go
func getByID[T any](db *gorm.DB, id uint) (T, error) {
	var model T
	err := db.First(&model, id).Error
	return model, err
}
