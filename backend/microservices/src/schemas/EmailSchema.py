import pydantic


class EmailSchema(pydantic.BaseModel):
    email: str
