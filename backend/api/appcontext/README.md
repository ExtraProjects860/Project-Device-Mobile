# Package `appcontext`

## Visão Geral

Este pacote fornece uma estrutura para gerenciamento de **contextos da aplicação**, comumente utilizados para armazenar e compartilhar recursos globais como:

- Conexões com **banco de dados**
- **Variáveis de ambiente**
- Outros recursos compartilhados entre diferentes partes da aplicação

Embora o uso mais comum seja para o **contexto global**, o pacote também suporta a criação de **contextos específicos** e isolados, oferecendo flexibilidade para diferentes necessidades da aplicação.

---

## Boas Práticas

- Utilize **structs e funções auxiliares** para encapsular os dados e evitar acesso direto às variáveis globais.
- **Evite mutações** nos estados internos dos contextos. Trabalhar com dados imutáveis ajuda a prevenir bugs, especialmente em ambientes concorrentes.
- Prefira a injeção de dependências via contexto ao invés de acessos diretos a instâncias globais.

---

## Estrutura Atual

- O pacote atualmente é composto por um único arquivo: `appcontext.go`.

---

## Exemplo de Uso *(opcional)*

```go
// Criação de um novo contexto com dependências
ctx := appcontext.New(appcontext.Config{
    DB: dbInstance,
    Env: os.Getenv("APP_ENV"),
})

// Acesso ao banco de dados a partir do contexto
db := appcontext.From(ctx).DB
