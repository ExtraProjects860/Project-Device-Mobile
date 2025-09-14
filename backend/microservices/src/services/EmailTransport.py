import abc
from email.mime.multipart import MIMEMultipart


class EmailTransport(abc.ABC):
    def __init__(self, hostname: str, port: int) -> None:
        self.__hostname: str = hostname
        self.__port: int = port

    @abc.abstractmethod
    async def send(
        self, message: MIMEMultipart, username: str, password: str
    ) -> None: ...

    @property
    def hostname(self) -> str:
        return self.__hostname

    @property
    def port(self) -> int:
        return self.__port
