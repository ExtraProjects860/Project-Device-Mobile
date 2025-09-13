import typing
from email.mime.multipart import MIMEMultipart

import aiosmtplib

from services.EmailTransport import EmailTransport


class TLSTransport(EmailTransport):
    def __init__(self, hostname: str, port: int) -> None:
        super().__init__(hostname, port)

    @typing.override
    async def send(self, message: MIMEMultipart, username: str, password: str) -> None:
        await aiosmtplib.send(
            message,
            hostname=self.hostname,
            port=self.port,
            username=username,
            password=password,
            start_tls=True,
            timeout=30,
        )
