import typing
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

from jinja2 import Template

from services.EmailSend import EmailSend
from services.EmailTransport import EmailTransport


class GmailSend(EmailSend):
    def __init__(
        self, email: str, password: str, transport: EmailTransport
    ) -> None:
        super().__init__(email, password, transport)

    @typing.override
    def formatter_variables(self, html: str, vars_template: dict[str, str]) -> str:
        template: Template = Template(html)
        return template.render(vars_template)

    @typing.override
    async def send_email(self, send_to: str, subject: str, html: str) -> None:
        message: MIMEMultipart = MIMEMultipart("alternative")
        message["From"] = self.email
        message["To"] = send_to
        message["Subject"] = subject
        message.attach(MIMEText(html, "html", "utf-8"))

        await self.transport.send(message, username=self.email, password=self.password)
