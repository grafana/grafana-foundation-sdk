from grafana_foundation_sdk.builders import timeseries
from grafana_foundation_sdk.models.common import AxisPlacement, GraphGradientMode
from grafana_foundation_sdk.models.dashboard import DynamicConfigValue
from grafana_foundation_sdk.models import units

from .common import prometheus_query, default_timeseries


def traffic_timeseries() -> timeseries.Panel:
    received_query = """irate(node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8
# only show interfaces that had traffic change at least once during selected dashboard interval:
and
increase(
    node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]
    ) > 0"""

    transmitted_query = """irate(node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8
# only show interfaces that had traffic change at least once during selected dashboard interval:
and
increase(
    node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]
    ) > 0"""

    return (
        default_timeseries()
        .title("Network traffic")
        .description(
            "Network traffic (bits per sec) measures data transmitted and received."
        )
        .axis_label("out(-) | in(+)")
        .axis_placement(AxisPlacement.AUTO)
        .with_target(prometheus_query(received_query, "{{ device }} received"))
        .with_target(prometheus_query(transmitted_query, "{{ device }} transmitted"))
        .decimals(1)
        .unit(units.BitsPerSecondSI)
        .gradient_mode(GraphGradientMode.OPACITY)
        .override_by_regexp(
            "/transmit|tx|out/",
            [DynamicConfigValue(id_val="custom.transform", value="negative-Y")],
        )
    )


def errors_timeseries() -> timeseries.Panel:
    description = """**Network errors**:

Network errors refer to issues that occur during the transmission of data across a network. 

These errors can result from various factors, including physical issues, jitter, collisions, noise and interference.

Monitoring network errors is essential for diagnosing and resolving issues, as they can indicate problems with network hardware or environmental factors affecting network quality.

**Dropped packets**:

Dropped packets occur when data packets traveling through a network are intentionally discarded or lost due to congestion, resource limitations, or network configuration issues. 

Common causes include network congestion, buffer overflows, QoS settings, and network errors, as corrupted or incomplete packets may be discarded by receiving devices.

Dropped packets can impact network performance and lead to issues such as degraded voice or video quality in real-time applications.
"""

    return (
        default_timeseries()
        .title("Network errors and dropped packets")
        .description(description)
        .axis_label("out(-) | in(+)")
        .axis_placement(AxisPlacement.AUTO)
        .with_target(
            prometheus_query(
                'irate(node_network_transmit_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0',
                "{{ device }} errors transmitted",
            )
        )
        .with_target(
            prometheus_query(
                'irate(node_network_receive_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0',
                "{{ device }} errors received",
            )
        )
        .with_target(
            prometheus_query(
                'irate(node_network_transmit_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0',
                "{{ device }} transmitted dropped",
            )
        )
        .with_target(
            prometheus_query(
                'irate(node_network_receive_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0',
                "{{ device }} received dropped",
            )
        )
        .decimals(1)
        .unit(units.PacketsPerSecond)
        .gradient_mode(GraphGradientMode.OPACITY)
        .override_by_regexp(
            "/transmit|tx|out/",
            [DynamicConfigValue(id_val="custom.transform", value="negative-Y")],
        )
    )
