# Package `enum`

## Visão Geral

O pacote `enum` centraliza **tipos enumerados** utilizados pela aplicação, com o objetivo de substituir o uso de strings **hardcoded** (ex: `"ADMIN"`, `"USER"`) em trechos sensíveis do código, como autenticação e autorização.

> ⚠️ A linguagem Go **não possui suporte nativo completo para enums**, como ocorre em linguagens como Java, C# ou Rust. Em Go, enums são implementados com **constantes nomeadas** de tipos customizados (`uint8`, `int`, etc.), o que oferece segurança limitada, mas ainda é útil para evitar erros comuns de digitação e melhorar a clareza do código.

> ⚠️ O ideal é que se evite o uso de enums, pois não são nativos, apenas utilizar para certos contextos e se extremamente necessário.

---

## Estrutura e Organização

Cada arquivo `.go` neste pacote representa um enum tendo um tipo, suas constantes e função String para puxar o enum como string.

### Arquivos atuais:

- `sqliteType.go` 
  Enum que estava sendo utilizado para testes na definição de criação de um sqlite em memória ou em arquivo, para evitar o uso do banco de dados postgres.
  > ⚠️ Observação: devido a dificuldades na hora de testar usando essa abordagem esse enum pode ser deletado por falta de uso ou contexto futuramente na aplicação. 

- `typeUser.go`
  Enum com utilidade para verificação de tipo do usuário por meio de middlewares ou diretamente por código. Utilizado também nas seeds para testes.

## Uso Atual

Atualmente o pacote está sendo utilizado apenas para o controle de **permissões e tipos de usuário**, substituindo o uso de strings diretas em middlewares ou validações.

Este modelo permite maior controle interno sobre as permissões do sistema, e facilita modificações futuras sem impactar diretamente os pontos de uso.

### Exemplo: `typeUser.go`

```go
package enum

type TypeUser uint8

const (
	SuperAdmin TypeUser = iota
	Admin
	User
)

func (t TypeUser) String() string {
	switch t {
	case SuperAdmin:
		return "SUPERADMIN"
	case Admin:
		return "ADMIN"
	case User:
		return "USER"
	default:
		return "UNKNOWN"
	}
}
