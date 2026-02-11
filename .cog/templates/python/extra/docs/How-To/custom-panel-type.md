# Defining a custom panel type

While the SDK ships with support for all core panels, it can be extended for
private/third-party plugins.

To do so, define a type and a builder for the custom panel's options:

```python
# src/custompanel.py
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
```

Register the type with cog, and use it as usual to build a dashboard:

```python
from grafana_foundation_sdk.builders.dashboard import Dashboard, Row
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.cog.plugins import register_default_plugins
from grafana_foundation_sdk.cog.runtime import register_panelcfg_variant

from src.custompanel import custom_panel_variant_config, CustomPanelBuilder


if __name__ == '__main__':
    # Required to correctly unmarshal panels and dataqueries
    register_default_plugins()

    # This lets cog know about the newly created panel type and how to unmarshal it.
    register_panelcfg_variant(custom_panel_variant_config())

    dashboard = (
        Dashboard("Custom panel type")
        .uid("test-custom-panel")
        .refresh("1m")
        .time("now-30m", "now")

        .with_row(Row("Overview"))
        .with_panel(
            CustomPanelBuilder()
            .title("Sample panel")
            .make_beautiful()
        )
    ).build()

    print(JSONEncoder(sort_keys=True, indent=2).encode(dashboard))
```
