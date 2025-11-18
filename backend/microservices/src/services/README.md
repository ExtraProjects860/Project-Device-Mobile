# 1. Visão Geral "EmailSend.py"
Este arquivo define uma classe base abstrata para todos os serviços de envio de email dentro da aplicação.
Ele serve como um contrato que classes como GmailSend, SendGridSend, etc. precisam seguir.

### Funcionamento
A classe "EmailSend" define como um serviço de envio deve se comportar, mas não implementa a lógica.
É o esqueleto do sistema de emails.

### Importações
```py
import abc
from services.EmailTransport import EmailTransport
```
- abc permite criar métodos abstratos.
- EmailTransport é uma camada de transporte (SSL, TLS etc.) responsável pela conexão SMTP.

### Relação com outras partes do sistema

O fluxo completo:

- >Rota FastAPI → email.py
- >Carrega HTML → FileHTML.read()
- >Aplica variáveis → EmailSend.formatter_variables()
- >Envia mensagem → EmailSend.send_email()
- >Transporte SMTP → EmailTransport, SSLTransport

##
# 2. Visão Geral "EmailTransport.py"
Este arquivo define a abstração da camada de transporte — ou seja, a classe responsável por abrir a conexão com o servidor SMTP e enviar o email.

Se a classe EmailSend descreve como montar emails, EmailTransport descreve como enviá-los pela rede.

Ela é um "provider de transporte" para email.

- > É onde classes concretas como
  - > SSLTransport
  - >TLSTransport
  - >DebugTransport

* vão implementar o envio real via SMTP.

Nota: Assim como EmailSend, ela é uma interface/contrato.

### Relação com o restante do sistema

Fluxo geral:
```scss
EmailSend (monta mensagem)
    ↓ chama
EmailTransport.send (envia via SMTP)
```

Ou seja:
- EmailSend cuida da lógica de negócio e formatação

- EmailTransport cuida da comunicação com o servidor SMTP

##
# 3. Visão Geral " GmailSend.py"

Função: Classe concreta responsável por montar e enviar emails via Gmail, usando o transporte definido (SSL/TLS/etc).

Essa classe implementa a interface EmailSend, definindo:

- Como substituir variáveis dentro do HTML
- Como montar o email (MIME)
- Como chamar a camada de transporte para enviar

### Importações
```py
import typing
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from jinja2 import Template
from services.EmailSend import EmailSend
from services.EmailTransport import EmailTransport
```
### O que cada uma faz:

- >typing.override → indica override de método (não obrigatório, mas bom para linting).
- >MIMEMultipart / MIMEText → objetos para construir emails compatíveis com SMTP.
- >jinja2.Template → motor de template usado para substituir variáveis no HTML.
- >EmailSend → classe abstrata que esta classe está implementando.
- >EmailTransport → tipo usado para enviar o email.

### Classe "GmailSend"
```py
class GmailSend(EmailSend):
```

Essa classe é especializada em:
- >preparar emails
- >aplicar variáveis no template
- >criar objeto MIME
- >delegar envio ao transporte

Ela não envia diretamente — apenas chama transport.send().

### Método: "formatter_variables"
```py
@typing.override
def formatter_variables(self, html: str, vars_template: dict[str, str]) -> str:
    template: Template = Template(html)
    return template.render(vars_template)
```

O que ele faz?
- Usa Jinja2 para substituir variáveis dentro do HTML.

Muito mais vantajoso já que usa .replace(), porque Jinja2 permite loops, blocos, ifs.

### Método: send_email
```py
@typing.override
async def send_email(self, send_to: str, subject: str, html: str) -> None:
    message: MIMEMultipart = MIMEMultipart("alternative")
    message["From"] = self.email
    message["To"] = send_to
    message["Subject"] = subject
    message.attach(MIMEText(html, "html", "utf-8"))

    await self.transport.send(message, username=self.email, password=self.password)
```
### Vantagens

Pode substituir o transporte facilmente:

- Gmail SSL
- Gmail TLS
- Amazon SES
- SMTP local
- Transporte gfalso para testes unitários

Tudo sem alterar esta classe.

### Resumo

- >GmailSend faz quatro coisas:
- >Recebe o HTML cru
- >Renderiza variáveis via Jinja2
- >Monta um email MIME completo
- >Entrega para o transporte enviar
- >Tudo de forma assíncrona.

##
# 4. Visão Geral "SLLTransport.py"

Função: Implementar o envio de email usando SMTP + SSL (TLS implícito) utilizando a biblioteca aiosmtplib.

A classe implementa a interface abstrata EmailTransport.

### Importações
```py
import typing
from email.mime.multipart import MIMEMultipart

import aiosmtplib

from services.EmailTransport import EmailTransport
```
- >typing.override → marca que estamos sobrescrevendo um método da classe abstrata.
- >MIMEMultipart → tipo da mensagem que será enviada (construída no GmailSend).
- >aiosmtplib → biblioteca assíncrona para envio de emails via SMTP.
- >EmailTransport → classe base abstrata que define o contrato.

### Método: send
```py
@typing.override
async def send(self, message: MIMEMultipart, username: str, password: str) -> None:
    await aiosmtplib.send(
        message,
        hostname=self.hostname,
        port=self.port,
        username=username,
        password=password,
        use_tls=True,
        timeout=30,
    )
```
Como funciona:
- > aiosmtplib.send() envia o email de forma assíncrona
- perfeito para FastAPI.

### Parâmetros enviados:

- >Parâmetro	Função
- >message	Objeto MIME pronto, vindo de GmailSend
- >hostname	Ex: smtp.gmail.com
- >port	Ex: 465
- >username	Email do remetente
- >password	Senha ou App Password
- >use_tls=True	Usa SSL/TLS implícito (porta 465)
- >timeout=30	Timeout padrão de 30s

* O transporte está configurado para conexão segura SSL.
* Não usa STARTTLS (que seria porta 587).

### Resumo

A classe SSLTransport:
- recebe o email já formatado
- se conecta ao servidor SMTP usando SSL
- envia a mensagem usando credenciais fornecidas
- funciona de forma assíncrona (não bloqueia o servidor)

##
# 5. Visão Geral "TLSTransport.py"

Função: Transportar emails via SMTP STARTTLS (criptografia iniciada após a conexão).

Ele implementa o mesmo contrato da classe abstrata EmailTransport, assim como o SSLTransport, mas muda como a conexão segura é estabelecida.

### Importações
```py
import typing
from email.mime.multipart import MIMEMultipart

import aiosmtplib

from services.EmailTransport import EmailTransport
```

Seguro e padronizado:
- typing.override para indicar sobrescrita
- MIMEMultipart para o conteúdo MIME
- aiosmtplib para envio assíncrono
- EmailTransport como contrato abstrato

### Diferença entre SSLTransport e TLSTransport
- >Tipo || Parâmetro chave || Porta comum || Como funciona
- >SSL (TLS implícito) || use_tls=True || 465 || Conexão já começa criptografada
- >STARTTLS (TLS explícito)|| start_tls=True || 587 || Começa em texto claro → negocia TLS depois

Logo:

- TLSTransport → envia email usando STARTTLS (587)
- SSLTransport → envia email usando conexão TLS direta (465)
  
Gmail aceita os dois, mas o mais moderno é STARTTLS (porta 587).

### Resumo

- A classe TLSTransport:
- Implementa envio via SMTP + STARTTLS.
- É equivalente ao SSLTransport, mas com TLS explícito.
- Ideal para servidores que exigem porta 587.
