from config.EmailProviders import GmailProvider
from config.Env import Env
from services.EmailTransport import EmailTransport
from services.GmailSend import GmailSend
from services.TLSTransport import TLSTransport


async def get_gmail_send() -> GmailSend:
    transport: EmailTransport = TLSTransport(GmailProvider.HOSTNAME, GmailProvider.PORT)
    return GmailSend(Env.EMAIL_USERNAME, Env.EMAIL_PASSWORD, transport)
