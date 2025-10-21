# Package `config`

## Visão Geral

O pacote `config` centraliza a **configuração principal do sistema**, incluindo:

- **Variáveis de ambiente**
- **Logger**
- **Banco de dados (SQL/Postgres)**

É o **ponto de partida da aplicação**, sendo responsável por inicializar e disponibilizar recursos essenciais via **dependency injection** (injeção de dependências) para os demais módulos do sistema.

---

## Responsabilidades

- Instanciar funcionalidades primordiais da aplicação
- Disponibilizar configurações reutilizáveis por outros pacotes
- Isolar a lógica de configuração em arquivos dedicados, facilitando manutenção e testes

---

## Boas Práticas

- Utilize as instâncias retornadas pela função `config.Init()` no `main.go` ou camada de inicialização da aplicação.
- Evite manipular diretamente recursos de configuração fora deste pacote.
- Organize configurações específicas (como de banco de dados ou logger) em arquivos separados dentro do pacote.

---

## Estrutura Atual

- `config.go`: Arquivo principal, contendo a função `Init`, responsável por:
  - Carregar variáveis de ambiente
  - Inicializar o banco de dados
  - Executar migrações com os modelos definidos
- `db.go`: Configuração do banco de dados PostgreSQL via GORM. O uso de `gorm.Dialector` permite extensão para outros bancos (NoSQL, etc.).
- `env_variables.go`: Responsável por carregar e validar variáveis de ambiente do `.env`.
- `logger.go`: Contém uma configuração básica de logger (ainda sem suporte a arquivo ou JSON de configuração).

---

## Exemplo de Uso

```go
package main

import (
	"log"

	"seu_modulo/config"
)

func main() {
	env, db, err := config.Init()
	if err != nil {
		log.Fatalf("Falha na inicialização do sistema: %v", err)
	}

	// Use env e db conforme necessário...
}
