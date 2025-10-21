# Package `db`

## Visão Geral

O pacote `db` centraliza funcionalidades de **migração** e **seed (preenchimento inicial)** do banco de dados da aplicação. Ele é composto por dois módulos com `main.go` próprios:

- [`migrate`](./migrate): Executa a limpeza e criação das tabelas com base nos schemas.
- [`seed`](./seed): Popula o banco com dados fake para testes e desenvolvimento.

Esses scripts foram implementados manualmente, com foco em agilidade, considerando a escala reduzida do projeto. Para sistemas maiores, recomenda-se migrar para uma biblioteca de gerenciamento de migrations como:

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate): Suporte robusto a versões, reversões e integração com diversos bancos.

---

## Diagrama do Banco de Dados

![Diagrama do Banco de Dados](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/backend/api/schemas/diagram-db-sql.png)

---

## Responsabilidades

- **Migração:** Recriar estruturas do banco com base nos schemas definidos no pacote `schemas`.
- **Seed:** Popular o banco com dados simulados usando libs como [`brianvoe/gofakeit`](https://github.com/brianvoe/gofakeit) e geradores nativos (`math/rand`).
- **Apoiar desenvolvimento:** Facilitar testes e visualização da aplicação com dados pré-carregados.

---

## Boas Práticas

- Para bancos que crescem em complexidade, considere a adoção de uma biblioteca de migrations (ex: `golang-migrate`) para:
  - Histórico de versões
  - Reversões controladas
  - Ambientes CI/CD mais estáveis
- Mantenha o **diagrama do banco de dados sempre atualizado**:
  - Atualize a imagem `diagrama-banco-de-dados-sql.png` e o arquivo `schema.dbml` sempre que houver mudanças estruturais.

---

## Exemplo de Execução

```bash
# Executar migration
go run ./migrate ou make migrate

# Executar seed
go run ./seed ou make seeder
