# Defining a custom query type

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
