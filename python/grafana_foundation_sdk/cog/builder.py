# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from abc import ABC, abstractmethod
from typing import Generic, TypeVar

T = TypeVar("T", covariant=True)


class Builder(Generic[T], ABC):
    @abstractmethod
    def build(self) -> T:
        pass
