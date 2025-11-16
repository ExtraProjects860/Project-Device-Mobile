# 1. Visão Geral "dependencies_email.py"

Este arquivo define uma dependência injetável usada normalmente pelo FastAPI (ou framework semelhante) para fornecer uma instância configurada do serviço de envio de e-mail.

### É responsável por:

- Criar o transporte SMTP seguro
- Criar a instância do serviço Gmail
- Usar variáveis de ambiente
- Entregar tudo isso pronto para os endpoints que precisam enviar e-mail

### Funcionamento 
```py
from config.EmailProviders import GmailProvider
from config.Env import Env
from services.EmailTransport import EmailTransport
from services.GmailSend import GmailSend
from services.SSLTransport import SSLTransport
```

### Importações:

- GmailProvider: contém hostname e porta SMTP do Gmail
- Env: acessa variáveis do .env (EMAIL_USERNAME, EMAIL_PASSWORD)
- EmailTransport: interface/base do transporte
- GmailSend: classe que envia e-mails usando Gmail
- SSLTransport: implementação concreta de transporte SMTP com SSL

### Função principal: "get_gmail_send()"
```py
async def get_gmail_send() -> GmailSend:
    transport: EmailTransport = SSLTransport(GmailProvider.HOSTNAME, GmailProvider.PORT)
    return GmailSend(Env.EMAIL_USERNAME, Env.EMAIL_PASSWORD, transport)
```

Essa função é construída exatamente no padrão de uma dependência FastAPI:

### 1. Cria o transporte seguro (SSL)
```py
transport = SSLTransport(GmailProvider.HOSTNAME, GmailProvider.PORT)
```

- >Host: "smtp.gmail.com"
- >Porta: 465
- >Protocolo: SSL (não STARTTLS)

### 2. Cria uma instância do serviço Gmail
```py
return GmailSend(Env.EMAIL_USERNAME, Env.EMAIL_PASSWORD, transport)
```

- >Ele injeta:
  - >Usuário do e-mail (do .env)
  - >Senha do e-mail (do .env)
  - >Transporte configurado

### 3. Retorna a instância pronta

Essa instância é usada pelo endpoint que enviará e-mail, geralmente assim:
```py
@gmail_router.post("/send")
async def send_email(data: EmailData, gmail: GmailSend = Depends(get_gmail_send)):
    return await gmail.send_mail(data)
``` 

### Vantagens:

- > Facilita reutilização da mesma configuração em vários endpoints
- >Permite troca futura de transportes com mínima alteração de código
- >Evita duplicar configuração de SMTP
- >Permite criar mocks em testes unitários
- >Força uma arquitetura limpa e modular

##
# 2. Visão Geral "email.py"
Este arquivo define um router do FastAPI responsável por receber requisições HTTP para envio de e-mails.

Como:

- >validação dos dados de entrada via schema
- >leitura do template HTML
- >substituição de variáveis dentro do template
- >criação de uma tarefa assíncrona para envio de e-mail
- >tratamento completo de erros

### Importações:
```py
Importações
import asyncio

from fastapi import APIRouter, Depends, HTTPException, status

from config.Paths import STATIC_DIR
from managers.File import FileHTML
from routers.dependencies_email import get_gmail_send
from schemas.Schemas import EmailSchema, Message
from services.EmailSend import EmailSend
```

### Detalhamento:

- >Import | Função
- >asyncio | Para criar tarefas assíncronas paralelas
- >APIRouter | Cria o router do FastAPI
- >Depends | Injeta dependências (ex.: serviço de email)
- >HTTPException | Erros HTTP padrão do FastAPI
- >STATIC_DIR | Caminho para pasta /static onde ficam templates HTML
- >FileHTML | Classe responsável por ler arquivos HTML assíncronos
- >get_gmail_send | Dependência que fornece serviço Gmail configurado
- >EmailSchema | Schema de entrada (quem recebe + template + dados)
- >Message | Schema de saída
- >EmailSend | Interface do envio de email

### REsumo

- >Recebe JSON com email + template + variáveis
- >Localiza o template HTML
- >Lê o arquivo assíncrono
- >Substitui variáveis dentro do HTML
- >Dispara envio em background
- >Retorna imediatamente para o cliente
- >Lida com erros de forma apropriada

##
# 3. Visão Geral "health_check.py"

Este é um router extremamente simples, com uma finalidade muito específica:
fornecer um endpoint de health check para monitoramento e verificação do status da API.

### Funcionamento
```py
from fastapi import APIRouter

router: APIRouter = APIRouter(prefix="/healthcheck", tags=["healthcheck"])
```

O que isso faz:

- Cria um pequeno grupo de rotas com prefixo:
```py
/healthcheck/*
```
As rotas aparecerão na documentação Swagger na seção healthcheck.

### Endpoint principal
```py
@router.get("/")
def read_root():
    return {"Hello": "World"}
```

O que importa aqui:

- >Método: GET
- >Caminho final: /healthcheck/
- >Não usa async, porque não faz nenhuma operação assíncrona.
- >Retorna apenas um JSON fixo:

```json
{
  "Hello": "World"
}
```

### Finalidade real

Apesar de simples, esse endpoint é MUITO importante para:

- >Verificar se o microserviço está rodando

Serviços externos (como Kubernetes, Nginx, AWS, Docker, Traefik) fazem requisições periódicas aqui para decidir se:

* o container está saudável
* o serviço está pronto para receber tráfego
* a aplicação precisa ser reiniciada

### Testar se a API responde

Um monitor (New Relic, Datadog, UptimeRobot, etc.) pode apontar para aqui.

### CI/CD também usa

Pipelines podem chamar esse endpoint após o deploy para validar.