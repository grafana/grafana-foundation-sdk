# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
import enum
from ..cog import runtime as cogruntime


class Options:
    view: 'MapViewConfig'
    controls: 'ControlsOptions'
    basemap: common.MapLayerOptions
    layers: list[common.MapLayerOptions]
    tooltip: 'TooltipOptions'

    def __init__(self, view: typing.Optional['MapViewConfig'] = None, controls: typing.Optional['ControlsOptions'] = None, basemap: typing.Optional[common.MapLayerOptions] = None, layers: typing.Optional[list[common.MapLayerOptions]] = None, tooltip: typing.Optional['TooltipOptions'] = None):
        self.view = view if view is not None else MapViewConfig()
        self.controls = controls if controls is not None else ControlsOptions()
        self.basemap = basemap if basemap is not None else common.MapLayerOptions()
        self.layers = layers if layers is not None else []
        self.tooltip = tooltip if tooltip is not None else TooltipOptions()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "view": self.view,
            "controls": self.controls,
            "basemap": self.basemap,
            "layers": self.layers,
            "tooltip": self.tooltip,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "view" in data:
            args["view"] = MapViewConfig.from_json(data["view"])
        if "controls" in data:
            args["controls"] = ControlsOptions.from_json(data["controls"])
        if "basemap" in data:
            args["basemap"] = common.MapLayerOptions.from_json(data["basemap"])
        if "layers" in data:
            args["layers"] = data["layers"]
        if "tooltip" in data:
            args["tooltip"] = TooltipOptions.from_json(data["tooltip"])        

        return cls(**args)


class MapViewConfig:
    id_val: str
    lat: typing.Optional[int]
    lon: typing.Optional[int]
    zoom: typing.Optional[int]
    min_zoom: typing.Optional[int]
    max_zoom: typing.Optional[int]
    padding: typing.Optional[int]
    all_layers: typing.Optional[bool]
    last_only: typing.Optional[bool]
    layer: typing.Optional[str]
    shared: typing.Optional[bool]

    def __init__(self, id_val: str = "zero", lat: typing.Optional[int] = 0, lon: typing.Optional[int] = 0, zoom: typing.Optional[int] = 1, min_zoom: typing.Optional[int] = None, max_zoom: typing.Optional[int] = None, padding: typing.Optional[int] = None, all_layers: typing.Optional[bool] = True, last_only: typing.Optional[bool] = None, layer: typing.Optional[str] = None, shared: typing.Optional[bool] = None):
        self.id_val = id_val
        self.lat = lat
        self.lon = lon
        self.zoom = zoom
        self.min_zoom = min_zoom
        self.max_zoom = max_zoom
        self.padding = padding
        self.all_layers = all_layers
        self.last_only = last_only
        self.layer = layer
        self.shared = shared

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
        }
        if self.lat is not None:
            payload["lat"] = self.lat
        if self.lon is not None:
            payload["lon"] = self.lon
        if self.zoom is not None:
            payload["zoom"] = self.zoom
        if self.min_zoom is not None:
            payload["minZoom"] = self.min_zoom
        if self.max_zoom is not None:
            payload["maxZoom"] = self.max_zoom
        if self.padding is not None:
            payload["padding"] = self.padding
        if self.all_layers is not None:
            payload["allLayers"] = self.all_layers
        if self.last_only is not None:
            payload["lastOnly"] = self.last_only
        if self.layer is not None:
            payload["layer"] = self.layer
        if self.shared is not None:
            payload["shared"] = self.shared
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "lat" in data:
            args["lat"] = data["lat"]
        if "lon" in data:
            args["lon"] = data["lon"]
        if "zoom" in data:
            args["zoom"] = data["zoom"]
        if "minZoom" in data:
            args["min_zoom"] = data["minZoom"]
        if "maxZoom" in data:
            args["max_zoom"] = data["maxZoom"]
        if "padding" in data:
            args["padding"] = data["padding"]
        if "allLayers" in data:
            args["all_layers"] = data["allLayers"]
        if "lastOnly" in data:
            args["last_only"] = data["lastOnly"]
        if "layer" in data:
            args["layer"] = data["layer"]
        if "shared" in data:
            args["shared"] = data["shared"]        

        return cls(**args)


class ControlsOptions:
    # Zoom (upper left)
    show_zoom: typing.Optional[bool]
    # let the mouse wheel zoom
    mouse_wheel_zoom: typing.Optional[bool]
    # Lower right
    show_attribution: typing.Optional[bool]
    # Scale options
    show_scale: typing.Optional[bool]
    # Show debug
    show_debug: typing.Optional[bool]
    # Show measure
    show_measure: typing.Optional[bool]

    def __init__(self, show_zoom: typing.Optional[bool] = None, mouse_wheel_zoom: typing.Optional[bool] = None, show_attribution: typing.Optional[bool] = None, show_scale: typing.Optional[bool] = None, show_debug: typing.Optional[bool] = None, show_measure: typing.Optional[bool] = None):
        self.show_zoom = show_zoom
        self.mouse_wheel_zoom = mouse_wheel_zoom
        self.show_attribution = show_attribution
        self.show_scale = show_scale
        self.show_debug = show_debug
        self.show_measure = show_measure

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.show_zoom is not None:
            payload["showZoom"] = self.show_zoom
        if self.mouse_wheel_zoom is not None:
            payload["mouseWheelZoom"] = self.mouse_wheel_zoom
        if self.show_attribution is not None:
            payload["showAttribution"] = self.show_attribution
        if self.show_scale is not None:
            payload["showScale"] = self.show_scale
        if self.show_debug is not None:
            payload["showDebug"] = self.show_debug
        if self.show_measure is not None:
            payload["showMeasure"] = self.show_measure
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "showZoom" in data:
            args["show_zoom"] = data["showZoom"]
        if "mouseWheelZoom" in data:
            args["mouse_wheel_zoom"] = data["mouseWheelZoom"]
        if "showAttribution" in data:
            args["show_attribution"] = data["showAttribution"]
        if "showScale" in data:
            args["show_scale"] = data["showScale"]
        if "showDebug" in data:
            args["show_debug"] = data["showDebug"]
        if "showMeasure" in data:
            args["show_measure"] = data["showMeasure"]        

        return cls(**args)


class TooltipOptions:
    mode: 'TooltipMode'

    def __init__(self, mode: typing.Optional['TooltipMode'] = None):
        self.mode = mode if mode is not None else TooltipMode.NONE

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]        

        return cls(**args)


class TooltipMode(enum.StrEnum):
    NONE = "none"
    DETAILS = "details"


class MapCenterID(enum.StrEnum):
    ZERO = "zero"
    COORDS = "coords"
    FIT = "fit"





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="geomap",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
