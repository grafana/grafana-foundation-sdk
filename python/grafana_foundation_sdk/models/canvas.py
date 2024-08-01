# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..models import common
from ..cog import runtime as cogruntime


class HorizontalConstraint(enum.StrEnum):
    LEFT = "left"
    RIGHT = "right"
    LEFT_RIGHT = "leftright"
    CENTER = "center"
    SCALE = "scale"


class VerticalConstraint(enum.StrEnum):
    TOP = "top"
    BOTTOM = "bottom"
    TOP_BOTTOM = "topbottom"
    CENTER = "center"
    SCALE = "scale"


class Constraint:
    horizontal: typing.Optional['HorizontalConstraint']
    vertical: typing.Optional['VerticalConstraint']

    def __init__(self, horizontal: typing.Optional['HorizontalConstraint'] = None, vertical: typing.Optional['VerticalConstraint'] = None):
        self.horizontal = horizontal
        self.vertical = vertical

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.horizontal is not None:
            payload["horizontal"] = self.horizontal
        if self.vertical is not None:
            payload["vertical"] = self.vertical
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "horizontal" in data:
            args["horizontal"] = data["horizontal"]
        if "vertical" in data:
            args["vertical"] = data["vertical"]        

        return cls(**args)


class Placement:
    top: typing.Optional[float]
    left: typing.Optional[float]
    right: typing.Optional[float]
    bottom: typing.Optional[float]
    width: typing.Optional[float]
    height: typing.Optional[float]
    rotation: typing.Optional[float]

    def __init__(self, top: typing.Optional[float] = None, left: typing.Optional[float] = None, right: typing.Optional[float] = None, bottom: typing.Optional[float] = None, width: typing.Optional[float] = None, height: typing.Optional[float] = None, rotation: typing.Optional[float] = None):
        self.top = top
        self.left = left
        self.right = right
        self.bottom = bottom
        self.width = width
        self.height = height
        self.rotation = rotation

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.top is not None:
            payload["top"] = self.top
        if self.left is not None:
            payload["left"] = self.left
        if self.right is not None:
            payload["right"] = self.right
        if self.bottom is not None:
            payload["bottom"] = self.bottom
        if self.width is not None:
            payload["width"] = self.width
        if self.height is not None:
            payload["height"] = self.height
        if self.rotation is not None:
            payload["rotation"] = self.rotation
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "top" in data:
            args["top"] = data["top"]
        if "left" in data:
            args["left"] = data["left"]
        if "right" in data:
            args["right"] = data["right"]
        if "bottom" in data:
            args["bottom"] = data["bottom"]
        if "width" in data:
            args["width"] = data["width"]
        if "height" in data:
            args["height"] = data["height"]
        if "rotation" in data:
            args["rotation"] = data["rotation"]        

        return cls(**args)


class BackgroundImageSize(enum.StrEnum):
    ORIGINAL = "original"
    CONTAIN = "contain"
    COVER = "cover"
    FILL = "fill"
    TILE = "tile"


class BackgroundConfig:
    color: typing.Optional[common.ColorDimensionConfig]
    image: typing.Optional[common.ResourceDimensionConfig]
    size: typing.Optional['BackgroundImageSize']

    def __init__(self, color: typing.Optional[common.ColorDimensionConfig] = None, image: typing.Optional[common.ResourceDimensionConfig] = None, size: typing.Optional['BackgroundImageSize'] = None):
        self.color = color
        self.image = image
        self.size = size

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.color is not None:
            payload["color"] = self.color
        if self.image is not None:
            payload["image"] = self.image
        if self.size is not None:
            payload["size"] = self.size
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "color" in data:
            args["color"] = common.ColorDimensionConfig.from_json(data["color"])
        if "image" in data:
            args["image"] = common.ResourceDimensionConfig.from_json(data["image"])
        if "size" in data:
            args["size"] = data["size"]        

        return cls(**args)


class LineConfig:
    color: typing.Optional[common.ColorDimensionConfig]
    width: typing.Optional[float]
    radius: typing.Optional[float]

    def __init__(self, color: typing.Optional[common.ColorDimensionConfig] = None, width: typing.Optional[float] = None, radius: typing.Optional[float] = None):
        self.color = color
        self.width = width
        self.radius = radius

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.color is not None:
            payload["color"] = self.color
        if self.width is not None:
            payload["width"] = self.width
        if self.radius is not None:
            payload["radius"] = self.radius
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "color" in data:
            args["color"] = common.ColorDimensionConfig.from_json(data["color"])
        if "width" in data:
            args["width"] = data["width"]
        if "radius" in data:
            args["radius"] = data["radius"]        

        return cls(**args)


class HttpRequestMethod(enum.StrEnum):
    GET = "GET"
    POST = "POST"
    PUT = "PUT"


class ConnectionCoordinates:
    x: float
    y: float

    def __init__(self, x: float = 0, y: float = 0):
        self.x = x
        self.y = y

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "x": self.x,
            "y": self.y,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "x" in data:
            args["x"] = data["x"]
        if "y" in data:
            args["y"] = data["y"]        

        return cls(**args)


class ConnectionPath(enum.StrEnum):
    STRAIGHT = "straight"


class CanvasConnection:
    source: 'ConnectionCoordinates'
    target: 'ConnectionCoordinates'
    target_name: typing.Optional[str]
    path: 'ConnectionPath'
    color: typing.Optional[common.ColorDimensionConfig]
    size: typing.Optional[common.ScaleDimensionConfig]
    vertices: typing.Optional[list['ConnectionCoordinates']]
    source_original: typing.Optional['ConnectionCoordinates']
    target_original: typing.Optional['ConnectionCoordinates']

    def __init__(self, source: typing.Optional['ConnectionCoordinates'] = None, target: typing.Optional['ConnectionCoordinates'] = None, target_name: typing.Optional[str] = None, path: typing.Optional['ConnectionPath'] = None, color: typing.Optional[common.ColorDimensionConfig] = None, size: typing.Optional[common.ScaleDimensionConfig] = None, vertices: typing.Optional[list['ConnectionCoordinates']] = None, source_original: typing.Optional['ConnectionCoordinates'] = None, target_original: typing.Optional['ConnectionCoordinates'] = None):
        self.source = source if source is not None else ConnectionCoordinates()
        self.target = target if target is not None else ConnectionCoordinates()
        self.target_name = target_name
        self.path = path if path is not None else ConnectionPath.STRAIGHT
        self.color = color
        self.size = size
        self.vertices = vertices
        self.source_original = source_original
        self.target_original = target_original

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "source": self.source,
            "target": self.target,
            "path": self.path,
        }
        if self.target_name is not None:
            payload["targetName"] = self.target_name
        if self.color is not None:
            payload["color"] = self.color
        if self.size is not None:
            payload["size"] = self.size
        if self.vertices is not None:
            payload["vertices"] = self.vertices
        if self.source_original is not None:
            payload["sourceOriginal"] = self.source_original
        if self.target_original is not None:
            payload["targetOriginal"] = self.target_original
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "source" in data:
            args["source"] = ConnectionCoordinates.from_json(data["source"])
        if "target" in data:
            args["target"] = ConnectionCoordinates.from_json(data["target"])
        if "targetName" in data:
            args["target_name"] = data["targetName"]
        if "path" in data:
            args["path"] = data["path"]
        if "color" in data:
            args["color"] = common.ColorDimensionConfig.from_json(data["color"])
        if "size" in data:
            args["size"] = common.ScaleDimensionConfig.from_json(data["size"])
        if "vertices" in data:
            args["vertices"] = data["vertices"]
        if "sourceOriginal" in data:
            args["source_original"] = ConnectionCoordinates.from_json(data["sourceOriginal"])
        if "targetOriginal" in data:
            args["target_original"] = ConnectionCoordinates.from_json(data["targetOriginal"])        

        return cls(**args)


class CanvasElementOptions:
    name: str
    type_val: str
    # TODO: figure out how to define this (element config(s))
    config: typing.Optional[object]
    constraint: typing.Optional['Constraint']
    placement: typing.Optional['Placement']
    background: typing.Optional['BackgroundConfig']
    border: typing.Optional['LineConfig']
    connections: typing.Optional[list['CanvasConnection']]

    def __init__(self, name: str = "", type_val: str = "", config: typing.Optional[object] = None, constraint: typing.Optional['Constraint'] = None, placement: typing.Optional['Placement'] = None, background: typing.Optional['BackgroundConfig'] = None, border: typing.Optional['LineConfig'] = None, connections: typing.Optional[list['CanvasConnection']] = None):
        self.name = name
        self.type_val = type_val
        self.config = config
        self.constraint = constraint
        self.placement = placement
        self.background = background
        self.border = border
        self.connections = connections

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "type": self.type_val,
        }
        if self.config is not None:
            payload["config"] = self.config
        if self.constraint is not None:
            payload["constraint"] = self.constraint
        if self.placement is not None:
            payload["placement"] = self.placement
        if self.background is not None:
            payload["background"] = self.background
        if self.border is not None:
            payload["border"] = self.border
        if self.connections is not None:
            payload["connections"] = self.connections
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "config" in data:
            args["config"] = data["config"]
        if "constraint" in data:
            args["constraint"] = Constraint.from_json(data["constraint"])
        if "placement" in data:
            args["placement"] = Placement.from_json(data["placement"])
        if "background" in data:
            args["background"] = BackgroundConfig.from_json(data["background"])
        if "border" in data:
            args["border"] = LineConfig.from_json(data["border"])
        if "connections" in data:
            args["connections"] = data["connections"]        

        return cls(**args)


class Options:
    # Enable inline editing
    inline_editing: bool
    # Show all available element types
    show_advanced_types: bool
    # Enable pan and zoom
    pan_zoom: bool
    # Enable infinite pan
    infinite_pan: bool
    # The root element of canvas (frame), where all canvas elements are nested
    # TODO: Figure out how to define a default value for this
    root: 'CanvasOptionsRoot'

    def __init__(self, inline_editing: bool = True, show_advanced_types: bool = True, pan_zoom: bool = True, infinite_pan: bool = True, root: typing.Optional['CanvasOptionsRoot'] = None):
        self.inline_editing = inline_editing
        self.show_advanced_types = show_advanced_types
        self.pan_zoom = pan_zoom
        self.infinite_pan = infinite_pan
        self.root = root if root is not None else CanvasOptionsRoot()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "inlineEditing": self.inline_editing,
            "showAdvancedTypes": self.show_advanced_types,
            "panZoom": self.pan_zoom,
            "infinitePan": self.infinite_pan,
            "root": self.root,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "inlineEditing" in data:
            args["inline_editing"] = data["inlineEditing"]
        if "showAdvancedTypes" in data:
            args["show_advanced_types"] = data["showAdvancedTypes"]
        if "panZoom" in data:
            args["pan_zoom"] = data["panZoom"]
        if "infinitePan" in data:
            args["infinite_pan"] = data["infinitePan"]
        if "root" in data:
            args["root"] = CanvasOptionsRoot.from_json(data["root"])        

        return cls(**args)


class CanvasOptionsRoot:
    # Name of the root element
    name: str
    # Type of root element (frame)
    type_val: typing.Literal["frame"]
    # The list of canvas elements attached to the root element
    elements: list['CanvasElementOptions']

    def __init__(self, name: str = "", elements: typing.Optional[list['CanvasElementOptions']] = None):
        self.name = name
        self.type_val = "frame"
        self.elements = elements if elements is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "type": self.type_val,
            "elements": self.elements,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "elements" in data:
            args["elements"] = data["elements"]        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="canvas",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
