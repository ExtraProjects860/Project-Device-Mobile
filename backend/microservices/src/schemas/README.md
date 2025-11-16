# Visão Geral "Schemas.py"
Este arquivo contém modelos Pydantic, usados para validar a entrada e saída de dados nos endpoints do microserviço de email.

### Funcionamento:
```py
import pydantic
```

Pydantic é usado para:

* validação automática de tipos
* serialização de dados
* documentação automática no Swagger / OpenAPI

### Modelo: EmailSchema
```py
class EmailSchema(pydantic.BaseModel):
    subject: str
    send_to: str
    template_name: str
    data: dict[str, str]
```

### Campos:

- >subject > str | Assunto do email.
- >send_to > str | Para quem o email será enviado. Deve ser um email válido (pydantic validará formato).
- >template_name > str | Nome do arquivo HTML localizado em /static/{template_name}.html.
- >data > dict[str, str] | Dicionário com variáveis usadas para substituir placeholders no template HTML.

### Exemplo de corpo válido enviado ao endpoint:
```json
{
  "subject": "Confirmação de cadastro",
  "send_to": "user@gmail.com",
  "template_name": "confirmation",
  "data": {
    "username": "João",
    "code": "98271"
  }
}
```
# Modelo: Message
```py
class Message(pydantic.BaseModel):
    message: str
```
Esse é um modelo de resposta padrão.

### Local aplicado

Retorno do endpoint:
```py
return Message(message="Email enviado com sucesso!")
```

### Principal motivo do uso

- >Permite ao FastAPI validar automaticamente a entrada.
- >Garante consistência do formato das respostas.
- >Facilita geração automática de documentação Swagger.
- >Evita erros como valores faltando, tipos incorretos, etc.