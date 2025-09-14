import os
from collections.abc import Generator
from dataclasses import dataclass, fields

import dotenv

dotenv.load_dotenv()


@dataclass(frozen=True)
class Env:
    EMAIL_USERNAME: str = os.getenv("EMAIL_USERNAME", "")
    EMAIL_PASSWORD: str = os.getenv("EMAIL_PASSWORD", "")

    def __iter__(self) -> Generator[str]:
        for field in fields(class_or_instance=self):
            yield getattr(self, field.name)


if __name__ == "__main__":
    env = Env()
    for v in env:
        print(v)
