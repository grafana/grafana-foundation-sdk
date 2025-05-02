from grafana_foundation_sdk.builders.dashboard import Dashboard
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.cog.plugins import register_default_plugins
from grafana_foundation_sdk.cog.runtime import register_panelcfg_variant

from src.custompanel import custom_panel_variant_config, CustomPanelBuilder


if __name__ == "__main__":
    # Required to correctly unmarshal panels and dataqueries
    register_default_plugins()

    # This lets cog know about the newly created panel type and how to unmarshal it.
    register_panelcfg_variant(custom_panel_variant_config())

    dashboard = (
        Dashboard("[Example] Custom panel")
        .uid("example-custom-panel")
        .with_panel(CustomPanelBuilder().title("Sample custom panel").make_beautiful())
    ).build()

    print(JSONEncoder(sort_keys=True, indent=2).encode(dashboard))
