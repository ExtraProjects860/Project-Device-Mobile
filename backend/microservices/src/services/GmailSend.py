from services.EmailSend import EmailSend
from services.EmailTransport import EmailTransport
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from jinja2 import Template
from managers.File import FileHTML
from config.Paths import STATIC_DIR

class GmailSend(EmailSend):

    def __init__(self, email: str, password: str, subject: str, transport: EmailTransport):
        self.email: str = email
        self.password: str = password
        self.subject: str = subject
        self.transport: EmailTransport = transport

    def formatter_variables(self, html:  str, vars_template: dict[str, str]):
        template: Template = Template(html)
        return template.render(vars_template)
    
    async def send_email(self, send_to: str, template_name: str, vars_template: dict[str, str]):
        html = await FileHTML(STATIC_DIR / f"{template_name}.html").read()
        if html == "":
            return
        
        html = self.formatter_variables(html, vars_template)

        message: MIMEMultipart = MIMEMultipart("alternative")
        message["From"] = self.email
        message["To"] = send_to
        message["Subject"] = self.subject
        message.attach(MIMEText(html, "html", "utf-8"))

        await self.transport.send(
            message,
            username=self.email,
            password=self.password
        )