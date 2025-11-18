from dataclasses import dataclass


@dataclass(frozen=True)
class GmailProvider:
    HOSTNAME: str = "smtp.gmail.com"
    PORT: int = 465
