import abc
import pathlib
import typing

import aiofiles


class File(abc.ABC):
    def __init__(self, path: pathlib.Path) -> None:
        self.__path: pathlib.Path = path
        self.__encoding: str = "utf-8"

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
        async with aiofiles.open(file=self.path, encoding=self.encoding) as file:
            return await file.read()
