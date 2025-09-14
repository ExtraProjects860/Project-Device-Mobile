import abc

from services.EmailTransport import EmailTransport


class EmailSend(abc.ABC):
    def __init__(
        self, email: str, password: str, transport: EmailTransport
    ) -> None:
        self.__email: str = email
        self.__password: str = password
        self.__transport: EmailTransport = transport

    @abc.abstractmethod
    def formatter_variables(self, html: str, vars_template: dict[str, str]) -> str: ...

    @abc.abstractmethod
    async def send_email(self, send_to: str, subject: str, html: str) -> None: ...

    @property
    def email(self) -> str:
        return self.__email

    @property
    def password(self) -> str:
        return self.__password

    @property
    def transport(self) -> EmailTransport:
        return self.__transport
