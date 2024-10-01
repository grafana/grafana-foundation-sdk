# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from dataclasses import dataclass
from typing import Any, Callable, Optional, Self
from . import variants as cogvariants


@dataclass
class DataqueryConfig:
    identifier: str
    from_json_hook: Callable[[dict[str, Any]], cogvariants.Dataquery]


@dataclass
class PanelCfgConfig:
    identifier: str
    options_from_json_hook: Optional[Callable[[dict[str, Any]], Any]] = None
    field_config_from_json_hook: Optional[Callable[[dict[str, Any]], Any]] = None


class Runtime:
    _instance = None
    dataquery_variants: dict[str, DataqueryConfig]
    panelcfg_variants: dict[str, PanelCfgConfig]

    def __new__(cls, *args, **kwargs):
        if cls._instance is None:
            cls._instance = object.__new__(cls, *args, **kwargs)
            cls.dataquery_variants = {}
            cls.panelcfg_variants = {}

        return cls._instance

    def register_dataquery_variant(self, variant: DataqueryConfig):
        self.dataquery_variants[variant.identifier] = variant

    def register_panelcfg_variant(self, variant: PanelCfgConfig):
        self.panelcfg_variants[variant.identifier] = variant

    def dataquery_from_json(self, data: dict[str, Any], dataquery_type_hint: str) -> cogvariants.Dataquery:
        if dataquery_type_hint == "" and "datasource" in data and "type" in data["datasource"]:
            dataquery_type_hint = data["datasource"]["type"]

        if dataquery_type_hint != "" and dataquery_type_hint in self.dataquery_variants:
            return self.dataquery_variants[dataquery_type_hint].from_json_hook(data)

        # We have no idea what type the dataquery is: use our `UnknownDataquery` bag to not lose data.
        return UnknownDataquery(data)

    def panelcfg_config(self, variant: str) -> Optional[PanelCfgConfig]:
        return self.panelcfg_variants.get(variant, None)


class UnknownDataquery(cogvariants.Dataquery):
    data: dict[str, Any]

    def __init__(self, data: dict[str, Any]):
        self.data = data

    def to_json(self) -> dict[str, object]:
        return self.data

    @classmethod
    def from_json(cls, data: dict[str, Any]) -> Self:
        return cls(data)


def dataquery_from_json(data: dict[str, Any], dataquery_type_hint: str) -> cogvariants.Dataquery:
    return Runtime().dataquery_from_json(data, dataquery_type_hint)


def panelcfg_config(variant: str) -> Optional[PanelCfgConfig]:
    return Runtime().panelcfg_config(variant)


def register_panelcfg_variant(variant: PanelCfgConfig):
    Runtime().register_panelcfg_variant(variant)


def register_dataquery_variant(variant: DataqueryConfig):
    Runtime().register_dataquery_variant(variant)
