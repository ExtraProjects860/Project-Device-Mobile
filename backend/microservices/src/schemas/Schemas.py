from fastapi import HTTPException
import pydantic
import typing

class EmailSchema(pydantic.BaseModel):
    subject: str
    send_to: str
    template_name: str
    data: dict[str, str]
    


class ResponseSuccess(pydantic.BaseModel):
    message: str
    data: typing.Optional[dict[str, str]]
    
def send_success(message: str, data: typing.Optional[dict[str, str]] = None) -> ResponseSuccess:
    return ResponseSuccess(message=message, data=data)

def send_err(status_code: int, err: str) -> HTTPException:
    return HTTPException(status_code=status_code, detail=err)
