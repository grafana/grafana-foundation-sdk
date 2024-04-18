from typing import Any, Self

from grafana_foundation_sdk.cog import builder
from grafana_foundation_sdk.cog import runtime as cogruntime
from grafana_foundation_sdk.builders.dashboard import Panel as PanelBuilder
from grafana_foundation_sdk.models import dashboard


class CustomPanelOptions:
    make_beautiful: bool

    def __init__(self, make_beautiful: bool = False):
        self.make_beautiful = make_beautiful

    def to_json(self) -> dict[str, object]:
        return {
            "makeBeautiful": self.make_beautiful,
        }

    @classmethod
    def from_json(cls, data: dict[str, Any]) -> Self:
        args: dict[str, Any] = {}

        if "makeBeautiful" in data:
            args["make_beautiful"] = data["makeBeautiful"]

        return cls(**args)


def custom_panel_variant_config() -> cogruntime.PanelCfgConfig:
    return cogruntime.PanelCfgConfig(
        # plugin ID
        identifier="custom-panel",
        options_from_json_hook=CustomPanelOptions.from_json,
    )


class CustomPanelBuilder(PanelBuilder, builder.Builder[dashboard.Panel]):
    def __init__(self):
        super().__init__()
        # plugin ID
        self._internal.type_val = "custom-panel"

    def make_beautiful(self) -> Self:
        if self._internal.options is None:
            self._internal.options = CustomPanelOptions()

        assert isinstance(self._internal.options, CustomPanelOptions)

        self._internal.options.make_beautiful = True

        return self
