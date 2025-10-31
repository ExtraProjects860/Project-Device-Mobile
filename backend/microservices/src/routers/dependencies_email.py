from config.EmailProviders import GmailProvider
from config.Env import Env
from services.EmailTransport import EmailTransport
from services.GmailSend import GmailSend
from services.SSLTransport import SSLTransport


async def get_gmail_send() -> GmailSend:
    transport: EmailTransport = SSLTransport(GmailProvider.HOSTNAME, GmailProvider.PORT)
    return GmailSend(Env.EMAIL_USERNAME, Env.EMAIL_PASSWORD, transport)
