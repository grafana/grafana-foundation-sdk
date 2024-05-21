# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime


class ArcOption:
    # Field from which to get the value. Values should be less than 1, representing fraction of a circle.
    field: typing.Optional[str]
    # The color of the arc.
    color: typing.Optional[str]

    def __init__(self, field: typing.Optional[str] = None, color: typing.Optional[str] = None):
        self.field = field
        self.color = color

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.color is not None:
            payload["color"] = self.color
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "color" in data:
            args["color"] = data["color"]        

        return cls(**args)


class NodeOptions:
    # Unit for the main stat to override what ever is set in the data frame.
    main_stat_unit: typing.Optional[str]
    # Unit for the secondary stat to override what ever is set in the data frame.
    secondary_stat_unit: typing.Optional[str]
    # Define which fields are shown as part of the node arc (colored circle around the node).
    arcs: typing.Optional[list['ArcOption']]

    def __init__(self, main_stat_unit: typing.Optional[str] = None, secondary_stat_unit: typing.Optional[str] = None, arcs: typing.Optional[list['ArcOption']] = None):
        self.main_stat_unit = main_stat_unit
        self.secondary_stat_unit = secondary_stat_unit
        self.arcs = arcs

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.main_stat_unit is not None:
            payload["mainStatUnit"] = self.main_stat_unit
        if self.secondary_stat_unit is not None:
            payload["secondaryStatUnit"] = self.secondary_stat_unit
        if self.arcs is not None:
            payload["arcs"] = self.arcs
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mainStatUnit" in data:
            args["main_stat_unit"] = data["mainStatUnit"]
        if "secondaryStatUnit" in data:
            args["secondary_stat_unit"] = data["secondaryStatUnit"]
        if "arcs" in data:
            args["arcs"] = data["arcs"]        

        return cls(**args)


class EdgeOptions:
    # Unit for the main stat to override what ever is set in the data frame.
    main_stat_unit: typing.Optional[str]
    # Unit for the secondary stat to override what ever is set in the data frame.
    secondary_stat_unit: typing.Optional[str]

    def __init__(self, main_stat_unit: typing.Optional[str] = None, secondary_stat_unit: typing.Optional[str] = None):
        self.main_stat_unit = main_stat_unit
        self.secondary_stat_unit = secondary_stat_unit

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.main_stat_unit is not None:
            payload["mainStatUnit"] = self.main_stat_unit
        if self.secondary_stat_unit is not None:
            payload["secondaryStatUnit"] = self.secondary_stat_unit
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mainStatUnit" in data:
            args["main_stat_unit"] = data["mainStatUnit"]
        if "secondaryStatUnit" in data:
            args["secondary_stat_unit"] = data["secondaryStatUnit"]        

        return cls(**args)


class Options:
    nodes: typing.Optional['NodeOptions']
    edges: typing.Optional['EdgeOptions']

    def __init__(self, nodes: typing.Optional['NodeOptions'] = None, edges: typing.Optional['EdgeOptions'] = None):
        self.nodes = nodes
        self.edges = edges

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.nodes is not None:
            payload["nodes"] = self.nodes
        if self.edges is not None:
            payload["edges"] = self.edges
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "nodes" in data:
            args["nodes"] = NodeOptions.from_json(data["nodes"])
        if "edges" in data:
            args["edges"] = EdgeOptions.from_json(data["edges"])        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="nodeGraph",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
