import pydantic


class EmailSchema(pydantic.BaseModel):
    subject: str
    send_to: str
    template_name: str
    data: dict[str, str]


class Message(pydantic.BaseModel):
    message: str
