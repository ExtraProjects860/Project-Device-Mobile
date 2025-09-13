import asyncio
from fastapi import APIRouter, status
from schemas.Schemas import EmailSchema, send_success, send_err
from config.Env import Env
from config.EmailProviders import GmailProvider
from services.TLSTransport import TLSTransport
from services.GmailSend import GmailSend 
from services.EmailTransport import EmailTransport


router: APIRouter = APIRouter(prefix="/email", tags=["email"])


@router.post("/")
async def read_root(email_data: EmailSchema):
    try:
        transport: EmailTransport = TLSTransport(GmailProvider.HOSTNAME, GmailProvider.PORT)
        
        if Env.EMAIL_USERNAME == "" or Env.EMAIL_PASSWORD == "":
            return send_err(status.HTTP_500_INTERNAL_SERVER_ERROR, "Failed to send email!")
        
        gmail_send: GmailSend = GmailSend(Env.EMAIL_USERNAME, Env.EMAIL_PASSWORD, email_data.subject, transport)
        asyncio.create_task(
            gmail_send.send_email(email_data.send_to, email_data.template_name, email_data.data))
        return send_success("Email send success!")
    except Exception as e:
        print(f"Failed to send email, error: {e}")
        return send_err(status.HTTP_500_INTERNAL_SERVER_ERROR, f"Failed to send email, error: {e}")
