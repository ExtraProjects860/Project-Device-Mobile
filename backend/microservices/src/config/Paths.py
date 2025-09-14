import pathlib
from typing import Final

CURRENT_DIR: Final[pathlib.Path] = pathlib.Path.cwd()
STATIC_DIR: Final[pathlib.Path] = CURRENT_DIR / "static"
