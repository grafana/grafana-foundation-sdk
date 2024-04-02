# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
import enum
from ..cog import runtime as cogruntime


class UpdateConfig:
    render: bool
    data_changed: bool
    schema_changed: bool

    def __init__(self, render: bool = False, data_changed: bool = False, schema_changed: bool = False):
        self.render = render
        self.data_changed = data_changed
        self.schema_changed = schema_changed

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "render": self.render,
            "dataChanged": self.data_changed,
            "schemaChanged": self.schema_changed,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "render" in data:
            args["render"] = data["render"]
        if "dataChanged" in data:
            args["data_changed"] = data["dataChanged"]
        if "schemaChanged" in data:
            args["schema_changed"] = data["schemaChanged"]        

        return cls(**args)


class DebugMode(enum.StrEnum):
    RENDER = "render"
    EVENTS = "events"
    CURSOR = "cursor"
    STATE = "State"
    THROW_ERROR = "ThrowError"


class Options:
    mode: 'DebugMode'
    counters: typing.Optional['UpdateConfig']

    def __init__(self, mode: typing.Optional['DebugMode'] = None, counters: typing.Optional['UpdateConfig'] = None):
        self.mode = mode if mode is not None else DebugMode.RENDER
        self.counters = counters

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        if self.counters is not None:
            payload["counters"] = self.counters
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "counters" in data:
            args["counters"] = UpdateConfig.from_json(data["counters"])        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="debug",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
