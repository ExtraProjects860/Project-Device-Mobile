from fastapi import APIRouter, status
from schemas.Schemas import EmailSchema, send_success, send_err
from config.Env import Env
from config.EmailProviders import GmailProvider
from services.TLSTransport import TLSTransport
from services.GmailSend import GmailSend 
from services.EmailTransport import EmailTransport


router: APIRouter = APIRouter(prefix="/email", tags=["email"])


@router.post("/", status_code=status.HTTP_200_OK)
async def read_root(email_data: EmailSchema):
    try:
        transport: EmailTransport = TLSTransport(GmailProvider.HOSTNAME, GmailProvider.PORT)
        
        if Env.EMAIL_USERNAME == "" or Env.EMAIL_PASSWORD == "":
            return send_err(status.HTTP_500_INTERNAL_SERVER_ERROR, "Failed to send email!")
        
        gmail_send: GmailSend = GmailSend(Env.EMAIL_USERNAME, Env.EMAIL_PASSWORD, email_data.subject, transport)
        await gmail_send.send_email(EmailSchema.send_to, EmailSchema.template_name, EmailSchema.data)
        return send_success("Email send success!")
    except Exception as e:
        print(f"Failed to send email, error: {e}")
        return send_err(status.HTTP_500_INTERNAL_SERVER_ERROR, f"Failed to send email, error: {e}")
