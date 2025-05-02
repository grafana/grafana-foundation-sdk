from grafana_foundation_sdk.builders.dashboard import Dashboard
from grafana_foundation_sdk.builders.timeseries import Panel as Timeseries
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.cog.plugins import register_default_plugins
from grafana_foundation_sdk.cog.runtime import register_dataquery_variant

from src.customquery import custom_query_variant_config, CustomQueryBuilder


if __name__ == "__main__":
    # Required to correctly unmarshal panels and dataqueries
    register_default_plugins()

    # This lets cog know about the newly created query type and how to unmarshal it.
    register_dataquery_variant(custom_query_variant_config())

    dashboard = (
        Dashboard("[Example] Custom query")
        .uid("example-custom-query")
        .with_panel(
            Timeseries()
            .title("Sample panel")
            .with_target(CustomQueryBuilder("query here"))
        )
    ).build()

    print(JSONEncoder(sort_keys=True, indent=2).encode(dashboard))
