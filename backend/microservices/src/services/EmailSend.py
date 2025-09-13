import typing

from services.EmailTransport import EmailTransport


class EmailSend(typing.Protocol):
    email: str
    password: str
    subject: str
    trasport: EmailTransport

    def formatter_variables(self, html: str, **kwargs: dict[str, str]) -> str: ...

    async def send_mail(
        self, send_to: str, template_name: str, **kwargs: dict[str, str]
    ) -> None: ...
