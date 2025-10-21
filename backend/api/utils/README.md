# Package `utils`

## Visão Geral

O pacote `utils` contém **funções utilitárias** utilizadas em diferentes partes da aplicação. Essas funções servem como ferramentas auxiliares para tarefas comuns e recorrentes no código, tais como:

- Geração e validação de **CPF** (para testes)
- Geração e comparação de **hashes com bcrypt**
- Instanciação global de **validações customizadas**
- Geração de imagens fake via serviço externo (`picsum.photos`)

---

## Estrutura e Organização

Cada arquivo `.go` neste pacote representa uma responsabilidade utilitária bem definida. Essa separação por responsabilidade facilita:

- Clareza e organização do código
- Reutilização em múltiplas camadas da aplicação
- Redução de duplicidade de lógica comum

### Arquivos atuais:

- `cpf.go`  
  Funções para gerar, formatar e validar CPFs (útil em ambiente de testes)

- `hash.go`  
  Geração e verificação de **hashes bcrypt**.  
  > ⚠️ Observação: caso seja necessário utilizar criptografia (e não apenas hash), recomenda-se criar um pacote separado com o uso do pacote `crypto`.

- `utils.go`  
  Funções mais gerais, incluindo:
  - Instanciação global de **validações com `validator`**
  - Geração de imagens fake usando o serviço **Picsum (https://picsum.photos/)**

- `validators.go`  
  Acesso direto a instâncias e configurações de validadores personalizados

---

## Boas Práticas

- **Isolamento de responsabilidades**: mantenha funções com responsabilidades distintas em arquivos separados, com nomes que expressem claramente seu propósito (`cpf.go`, `hash.go`, etc).

- **Evite lógica duplicada**: sempre utilize os utilitários centralizados (ex: `utils.HashPassword()`, `utils.IsValidCpf()`) ao invés de duplicar a lógica nos handlers ou serviços.

- **Criptografia ≠ Hash**: o uso de `bcrypt` (em `hash.go`) serve para armazenamento seguro de senhas. Para **criptografia real (reversível)**, utilizar pacotes como `crypto/aes` e organizar em outro módulo.

- **Validações globais**: centralize a configuração dos validadores no `utils.go` ou `validators.go`, para evitar múltiplas instâncias conflitantes ou revalidações desnecessárias.

---

## Exemplo de Uso

```go
import (
	"project/internal/utils"
)

// Verificação de CPF
if !utils.ValidateCPF("123.456.789-09") {
	log.Println("CPF inválido")
}

// Geração de hash
hashed, _ := utils.GenerateHashPassword("senha123")

// Comparação
if err := utils.VerifyHashedPassword("senha123", hashed); err != nil {
	log.Println("Senha válida")
}
