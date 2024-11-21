# Grafana Foundation SDK – Python

A set of tools, types and *builder libraries* for building and manipulating Grafana objects in Python.

> [!NOTE]
> This branch contains **types and builders generated for Grafana v11.0.x.**
> Other supported versions of Grafana can be found at [this repository's root](https://github.com/grafana/grafana-foundation-sdk/).

## Installing

```shell
python3 -m pip install 'grafana_foundation_sdk==1732188160!11.0.0'
```

## Example usage

### Building a dashboard

```python
from grafana_foundation_sdk.builders.dashboard import Dashboard, Row
from grafana_foundation_sdk.builders.prometheus import Dataquery as PrometheusQuery
from grafana_foundation_sdk.builders.timeseries import Panel as Timeseries
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.models.common import TimeZoneBrowser

def build_dashboard() -> Dashboard:
    builder = (
        Dashboard("[TEST] Node Exporter / Raspberry")
        .uid("test-dashboard-raspberry")
        .tags(["generated", "raspberrypi-node-integration"])

        .refresh("1m")
        .time("now-30m", "now")
        .timezone(TimeZoneBrowser)

        .with_row(Row("Overview"))
        .with_panel(
            Timeseries()
            .title("Network Received")
            .unit("bps")
            .min_val(0)
            .with_target(
                PrometheusQuery()
                .expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
                .legend_format("{{ device }}")
            )
        )
    )

    return builder


if __name__ == '__main__':
    dashboard = build_dashboard().build()
    encoder = JSONEncoder(sort_keys=True, indent=2)

    print(encoder.encode(dashboard))
```

### Unmarshaling a dashboard

```python
import json

from grafana_foundation_sdk.cog.plugins import register_default_plugins
from grafana_foundation_sdk.models.dashboard import Dashboard as DashboardModel


if __name__ == '__main__':
    # Required to correctly unmarshal panels and dataqueries
    register_default_plugins()

    with open("dashboard.json", "r") as f:
        decoded_dashboard = DashboardModel.from_json(json.load(f))
        print(decoded_dashboard)
```

### Defining a custom query type

While the SDK ships with support for all core datasources and their query types,
it can be extended for private/third-party plugins.

To do so, define a type and a builder for the custom query:

```python
# src/customquery.py
from typing import Any, Optional, Self

from grafana_foundation_sdk.cog import variants as cogvariants
from grafana_foundation_sdk.cog import runtime as cogruntime
from grafana_foundation_sdk.cog import builder


class CustomQuery(cogvariants.Dataquery):
    # ref_id and hide are expected on all queries
    ref_id: Optional[str]
    hide: Optional[bool]

    # query is specific to the CustomQuery type
    query: str

    def __init__(self, query: str, ref_id: Optional[str] = None, hide: Optional[bool] = None):
        self.query = query
        self.ref_id = ref_id
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "query": self.query,
        }
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, Any]) -> Self:
        args: dict[str, Any] = {}

        if "query" in data:
            args["query"] = data["query"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]

        return cls(**args)


def custom_query_variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        # datasource plugin ID
        identifier="custom-query",
        from_json_hook=CustomQuery.from_json,
    )


class CustomQueryBuilder(builder.Builder[CustomQuery]):
    _internal: CustomQuery

    def __init__(self, query: str):
        self._internal = CustomQuery(query=query)

    def build(self) -> CustomQuery:
        return self._internal

    def ref_id(self, ref_id: str) -> Self:
        self._internal.ref_id = ref_id

        return self

    def hide(self, hide: bool) -> Self:
        self._internal.hide = hide

        return self
```

Register the type with cog, and use it as usual to build a dashboard:

```python
from grafana_foundation_sdk.builders.dashboard import Dashboard, Row
from grafana_foundation_sdk.builders.timeseries import Panel as Timeseries
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.cog.plugins import register_default_plugins
from grafana_foundation_sdk.cog.runtime import register_dataquery_variant

from src.customquery import custom_query_variant_config, CustomQueryBuilder


if __name__ == '__main__':
    # Required to correctly unmarshal panels and dataqueries
    register_default_plugins()

    # This lets cog know about the newly created query type and how to unmarshal it.
    register_dataquery_variant(custom_query_variant_config())

    dashboard = (
        Dashboard("Custom query type")
        .uid("test-custom-query")
        .refresh("1m")
        .time("now-30m", "now")

        .with_row(Row("Overview"))
        .with_panel(
            Timeseries()
            .title("Sample panel")
            .with_target(
                CustomQueryBuilder("query here")
            )
        )
    ).build()

    print(JSONEncoder(sort_keys=True, indent=2).encode(dashboard))
```

### Defining a custom panel type

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

## Maturity

The code in this repository should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](./LICENSE)
