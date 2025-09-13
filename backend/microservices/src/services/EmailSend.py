import typing

from services.EmailTransport import EmailTransport


class EmailSend(typing.Protocol):
    email: str
    password: str
    subject: str
    transport: EmailTransport

    def formatter_variables(self, html: str, vars_template: dict[str, str]) -> str: ...

    async def send_email(
        self, send_to: str, template_name: str, vars_template: dict[str, str]
    ) -> None: ...
