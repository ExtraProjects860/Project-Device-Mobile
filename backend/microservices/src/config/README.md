# 1. Visão Geral "EmailProviders.py"

Este arquivo define provedores de e-mail utilizados pelo microserviço de envio de mensagens.
Atualmente, há apenas um provedor configurado: Gmail.

Ele utiliza o @dataclass do Python para garantir uma estrutura imutável, clara e segura para as credenciais do provedor SMTP.


### Funcionamento
```py
from dataclasses import dataclass


@dataclass(frozen=True)
class GmailProvider:
    HOSTNAME: str = "smtp.gmail.com"
    PORT: int = 465
```


- >"@dataclass(frozen=True)"
  - >Torna a classe imutável (read-only).

- Isso impede que o código altere acidentalmente o HOSTNAME ou PORT.
  - É ideal para dados de configuração que não devem mudar durante a execução.

- Classe GmailProvider
  - Representa a configuração SMTP do Gmail. 

### Objetivo da Classe

* Fornecer uma configuração centralizada e segura para serviços que vão:

- >Enviar e-mails via Gmail

- >Criar conexões SMTP

- >Configurar microserviços de comunicação
#
# 2. Visão Geral "Env.py"

Este arquivo é responsável por carregar variáveis de ambiente necessárias para o microserviço, como credenciais de e-mail.
Ele utiliza dotenv para carregar valores do arquivo .env e uma @dataclass imutável para armazenar essas variáveis com segurança.

A classe Env também implementa um iterador, permitindo percorrer automaticamente todos os campos definidos.

### Funcionamento

```py
import os
from collections.abc import Generator
from dataclasses import dataclass, fields

import dotenv

dotenv.load_dotenv()


@dataclass(frozen=True)
class Env:
    EMAIL_USERNAME: str = os.getenv("EMAIL_USERNAME", "")
    EMAIL_PASSWORD: str = os.getenv("EMAIL_PASSWORD", "")

    def __iter__(self) -> Generator[str]:
        for field in fields(class_or_instance=self):
            yield getattr(self, field.name)


if __name__ == "__main__":
    env = Env()
    for v in env:
        print(v)
```
### "dotenv.load_dotenv()"

- Carrega automaticamente variáveis do arquivo .env localizado na raiz que o Python consegue encontrar.
- Permite armazenar credenciais fora do código-fonte.

- >Exemplo do .env esperado:

EMAIL_USERNAME=seuemail@gmail.com
EMAIL_PASSWORD=suasenha

### Classe Env

Uma @dataclass imutável (frozen=True) que representa as configurações sensíveis da aplicação.

- Campos:
- >Campo	Valor padrão	Origem
- >EMAIL_USERNAME	Variável de ambiente
- >EMAIL_PASSWORD	 Variável de ambiente

Caso a variável não exista, uma string vazia é usada por padrão.

- Iterador __iter__

A classe implementa:
```py
def __iter__(self):
    for field in fields(self):
        yield getattr(self, field.name)
```

- Isso permite:
```py
env = Env()
for value in env:
    print(value)
```

- Saída (exemplo):
```css
seuemail@gmail.com
suasenha
```
- >Útil para depuração, testes ou logs internos (desde que não exponha credenciais em produção).

### Uso no microserviço

A classe Env é normalmente utilizada para inicializar serviços:

```py
from microservices.config.Env import Env

env = Env()

smtp_login = env.EMAIL_USERNAME
smtp_password = env.EMAIL_PASSWORD
```

#
# 3. Visão Geral "Paths.py"

Este arquivo visa centralizar caminhos importantes do projeto, principalmente o diretório raiz atual e a pasta static. Ele facilita o acesso a recursos estáticos (imagens, templates, logs, arquivos gerados, etc) sem precisar lidar com paths montados manualmente.

### Funcionamento
```py
import pathlib
from typing import Final

CURRENT_DIR: Final[pathlib.Path] = pathlib.Path.cwd()
STATIC_DIR: Final[pathlib.Path] = CURRENT_DIR / "static"
```

- > "pathlib"
  - > Biblioteca moderna do Python para manipulação de caminhos com segurança, portabilidade e sem precisar usar concatenar strings com /.

- > "CURRENT_DIR"
```py
CURRENT_DIR: Final[pathlib.Path] = pathlib.Path.cwd()
```

- Representa o diretório atual onde o processo está sendo executado.
- Ele não é necessariamente o diretório do arquivo Paths.py, mas sim o Current Working Directory (CWD).
- Em execução normal, será a raiz do backend ou do microserviço.

### Tipagem Final

Ambas as variáveis usam:

```py
from typing import Final
```

Isso indica ao type checker (ex: mypy) que esses valores não devem ser alterados em lugar algum do código:
```py
CURRENT_DIR = 123 
```

Melhora consistência e segurança.

# 4. Visão Geral "Paths.py"

Ele define caminhos (paths) globais para o microserviço.

##
# 5. Visão Geral "main.py"

```py
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from routers import email, health_check

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

base_path = "/api/v1"

app.include_router(router=email.router, prefix=base_path)
app.include_router(router=health_check.router, prefix=base_path)
```

### 1. Cria a instância principal do FastAPI
```py
app = FastAPI()
```
Este é o core do microserviço.

### 2. Configura CORS
```py
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
```
Significa:

- >permitir requisições de qualquer origem (allow_origins=["*"])
- >permitir envio de cookies e headers de autenticação
- >liberar todos os métodos (GET, POST, PUT, DELETE...)
- >liberar qualquer header