# Building a dashboard

```python
from grafana_foundation_sdk.builders.dashboard import Dashboard, Row
from grafana_foundation_sdk.builders import prometheus, timeseries
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.models.common import TimeZoneBrowser
from grafana_foundation_sdk.models import units


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
            timeseries.Panel()
            .title("Network Received")
            .unit(units.BitsPerSecondSI)
            .min(0)
            .with_target(
                prometheus.Dataquery()
                .expr(
                    'rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8'
                )
                .legend_format("{{ device }}")
            )
        )
    )

    return builder


if __name__ == "__main__":
    dashboard = build_dashboard().build()
    encoder = JSONEncoder(sort_keys=True, indent=2)

    print(encoder.encode(dashboard))
```
