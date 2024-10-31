# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..cog import runtime as cogruntime


class TextMode(enum.StrEnum):
    HTML = "html"
    MARKDOWN = "markdown"
    CODE = "code"


class CodeLanguage(enum.StrEnum):
    JSON = "json"
    YAML = "yaml"
    XML = "xml"
    TYPESCRIPT = "typescript"
    SQL = "sql"
    GO = "go"
    MARKDOWN = "markdown"
    HTML = "html"
    PLAINTEXT = "plaintext"


class CodeOptions:
    # The language passed to monaco code editor
    language: 'CodeLanguage'
    show_line_numbers: bool
    show_mini_map: bool

    def __init__(self, language: typing.Optional['CodeLanguage'] = None, show_line_numbers: bool = False, show_mini_map: bool = False):
        self.language = language if language is not None else CodeLanguage.PLAINTEXT
        self.show_line_numbers = show_line_numbers
        self.show_mini_map = show_mini_map

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "language": self.language,
            "showLineNumbers": self.show_line_numbers,
            "showMiniMap": self.show_mini_map,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "language" in data:
            args["language"] = data["language"]
        if "showLineNumbers" in data:
            args["show_line_numbers"] = data["showLineNumbers"]
        if "showMiniMap" in data:
            args["show_mini_map"] = data["showMiniMap"]        

        return cls(**args)


class Options:
    mode: 'TextMode'
    code: typing.Optional['CodeOptions']
    content: str

    def __init__(self, mode: typing.Optional['TextMode'] = None, code: typing.Optional['CodeOptions'] = None, content: str = "# Title\n\nFor markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)"):
        self.mode = mode if mode is not None else TextMode.MARKDOWN
        self.code = code
        self.content = content

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "content": self.content,
        }
        if self.code is not None:
            payload["code"] = self.code
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "code" in data:
            args["code"] = CodeOptions.from_json(data["code"])
        if "content" in data:
            args["content"] = data["content"]        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="text",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
