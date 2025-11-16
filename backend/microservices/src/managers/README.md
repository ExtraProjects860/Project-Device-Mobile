# Visão Geral "EmailTemplates.py"

Este arquivo define um Enum contendo os nomes dos arquivos HTML usados como templates de e-mail dentro do microserviço de e-mail.

Ele padroniza e evita erros de digitação quando outros componentes do backend precisam carregar ou enviar esses e-mails.

### Funcionamento
```py
import enum

class EmailTemplates(enum.Enum):
    CONFIRMATION = "confimation.html"
    CHANGING_PASSWORD = "changing_password.html"
```

### Uso do Enum

- "enum.Enum" -> permite criar valores nomeados imutáveis e seguros.

- >Vantagens
  - >Evita strings duplicadas espalhadas pelo código
  - >Evita erros como "comfirmation.html" vs "confirmation.html"
  - >Permite autocomplete e validação
  - >Facilita refatoração futura