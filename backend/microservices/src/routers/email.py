import asyncio

from fastapi import APIRouter, Depends, HTTPException, status

from config.Paths import STATIC_DIR
from managers.File import FileHTML
from routers.dependencies_email import get_gmail_send
from schemas.Schemas import EmailSchema, Message
from services.EmailSend import EmailSend

router: APIRouter = APIRouter(prefix="/email", tags=["email"])


@router.post("/", response_model=Message)
async def send_email_to(
    email_data: EmailSchema, gmail_send: EmailSend = Depends(get_gmail_send)
):
    try:
        html: str = await FileHTML(
            STATIC_DIR / f"{email_data.template_name}.html"
        ).read()
        html = gmail_send.formatter_variables(html, email_data.data)

        asyncio.create_task(
            coro=gmail_send.send_email(email_data.send_to, email_data.subject, html)
        )
        return Message(message=f"Email {email_data.send_to} send success!")
    except FileNotFoundError as e:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=str(e))
    except Exception as e:
        return HTTPException(
            status.HTTP_500_INTERNAL_SERVER_ERROR, f"Failed to send email, error: {e}"
        )
