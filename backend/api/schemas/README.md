# Package `schemas`

## Visão Geral

O pacote `schemas` contém as definições das **entidades do banco de dados** utilizadas pela aplicação, implementadas como structs Go com **tags GORM**, o ORM utilizado no projeto.

> Atualmente o sistema está configurado para usar **PostgreSQL** como banco de dados relacional.

O GORM é altamente compatível com bancos **relacionais** (PostgreSQL, MySQL, SQLite, SQL Server etc.). Ele **não é projetado para trabalhar com bancos NoSQL** diretamente. Para bancos NoSQL (como MongoDB ou Redis), o ideal é utilizar bibliotecas específicas.

---

## Estrutura e Organização

Cada arquivo `.go` representa um schema (tabela) específico. O nome do arquivo segue o padrão `nomeDaTabela.go`, e cada struct está associada diretamente a uma tabela do banco.

### Arquivos atuais:

- `enterprise.go`: Schema para empresas ou organizações
- `product.go`: Produtos
- `role.go`: Perfis/funções de usuário
- `schemas.go`: Inicialização e configuração do GORM (e possivelmente auto-migração)
- `tokenPassword.go`: Tokens de redefinição de senha ou autenticação
- `user.go`: Usuários do sistema
- `wishlist.go`: Lista de desejos ou favoritos

> ⚠️ Sempre que for adicionar uma nova tabela ao banco de dados, crie um novo arquivo seguindo o padrão `nomedatabela.go`.

---

## Diagrama do Banco de Dados

Visualize abaixo o modelo atual de dados:

![Diagrama do Banco de Dados](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/backend/api/schemas/diagram-db-sql.png)

---

## Boas Práticas

- Mantenha **um schema por arquivo**, nomeando conforme o nome da tabela.
- Use **tags GORM** para mapear corretamente as colunas, chaves estrangeiras e relacionamentos.
- Centralize a lógica de migração no arquivo `schemas.go` para manter consistência.
- Atualize o **diagrama do banco** sempre que novos schemas forem criados ou alterados.

---

## Exemplo de Schema com GORM

```go
// user.go
package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID         uint `gorm:"not null"`
	EnterpriseID   *uint
	Name           string `gorm:"not null"`
	Email          string `gorm:"uniqueIndex;not null"`
	Password       string `gorm:"not null"`
	Cpf            string `gorm:"uniqueIndex;not null"`
	RegisterNumber string `gorm:"not null"`
	PhotoUrl       *string

	Role            Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Enterprise      Enterprise    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenPassword   TokenPassword `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WishListEntries []WishList    `gorm:"foreignKey:UserID"`
}
