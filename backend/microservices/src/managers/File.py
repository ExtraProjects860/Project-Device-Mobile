import abc
import pathlib
import typing

import aiofiles


class File(abc.ABC):
    def __init__(self, path: pathlib.Path) -> None:
        self.__path: pathlib.Path = path
        self.__encoding: str = "utf-8"

    def file_exists(self):
        return self.__path.exists()

    @abc.abstractmethod
    async def read(self) -> str:
        msg = "Method not implemented yet."
        raise NotImplementedError(msg)

    @property
    def path(self) -> pathlib.Path:
        return self.__path

    @property
    def encoding(self) -> str:
        return self.__encoding


class FileHTML(File):
    def __init__(self, path: pathlib.Path) -> None:
        super().__init__(path)

    @typing.override
    async def read(self) -> str:
        if not self.file_exists():
            raise FileNotFoundError("Template html not found or invalid name")

        async with aiofiles.open(file=self.path, encoding=self.encoding) as file:
            return await file.read()
