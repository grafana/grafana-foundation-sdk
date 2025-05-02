from grafana_foundation_sdk.builders import dashboard, stat, timeseries
from grafana_foundation_sdk.models.common import BigValueColorMode, BigValueGraphMode
from grafana_foundation_sdk.models.dashboard import (
    DynamicConfigValue,
    ThresholdsMode,
    Threshold,
)
from grafana_foundation_sdk.models import units

from .common import prometheus_query, default_stat, default_timeseries


def count_stat() -> stat.Panel:
    return (
        default_stat()
        .title("CPU count")
        .description(
            "CPU count is the number of processor cores or central processing units (CPUs) in a computer,\ndetermining its processing capability and ability to handle tasks concurrently."
        )
        .with_target(
            prometheus_query(
                query='count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})',
                legend="Cores",
            )
        )
        .decimals(0)
        .color_mode(BigValueColorMode.NONE)
    )


def usage_stat() -> stat.Panel:
    query = """(((count by (instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance)))
    -
    avg by (instance) (sum by (instance, mode)(irate(node_cpu_seconds_total{mode='idle',job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])))) * 100)
    /
    count by(instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance))"""

    description = """Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.
It represents the combined load placed on all CPU cores or processors.

For instance, if the total CPU utilization percent is 50%, it means that,
on average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high."""

    return (
        default_stat()
        .title("CPU usage")
        .description(description)
        .with_target(prometheus_query(query=query))
        .min(0)
        .max(100)
        .unit(units.Percent)
        .thresholds(
            dashboard.ThresholdsConfig()
            .mode(ThresholdsMode.ABSOLUTE)
            .steps(
                [
                    Threshold(color="green"),
                    Threshold(value=80, color="red"),
                ]
            )
        )
        .graph_mode(BigValueGraphMode.AREA)
    )


def usage_timeseries() -> timeseries.Panel:
    query = """(
  (1 - sum without (mode) (rate(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode=~"idle|iowait|steal"}[$__rate_interval])))
/ ignoring(cpu) group_left
  count without (cpu, mode) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})
) * 100"""

    description = """Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.
It represents the combined load placed on all CPU cores or processors.

For instance, if the total CPU utilization percent is 50%, it means that,
on average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high."""

    return (
        default_timeseries()
        .title("CPU usage")
        .description(description)
        .unit(units.Percent)
        .decimals(1)
        .with_target(prometheus_query(query=query, legend="CPU {{cpu}}"))
        .min(0)
        .max(100)
        .thresholds(
            dashboard.ThresholdsConfig()
            .mode(ThresholdsMode.ABSOLUTE)
            .steps(
                [
                    Threshold(color="green"),
                    Threshold(value=80, color="red"),
                ]
            )
        )
    )


def load_average_timeseries() -> timeseries.Panel:
    description = """System load average over the previous 1, 5, and 15 minute ranges.

A measurement of how many processes are waiting for CPU cycles. The maximum number is the number of CPU cores for the node."""

    return (
        default_timeseries()
        .title("Load average")
        .description(description)
        .with_target(
            prometheus_query(
                query='node_load1{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                legend="1m",
            )
        )
        .with_target(
            prometheus_query(
                query='node_load5{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                legend="5m",
            )
        )
        .with_target(
            prometheus_query(
                query='node_load15{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                legend="15m",
            )
        )
        .with_target(
            prometheus_query(
                query='count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})',
                legend="Cores",
            )
        )
        .unit(units.Short)
        .decimals(2)
        .min(0)
        .thresholds(
            dashboard.ThresholdsConfig()
            .mode(ThresholdsMode.ABSOLUTE)
            .steps(
                [
                    Threshold(color="green"),
                    Threshold(value=80, color="red"),
                ]
            )
        )
        .override_by_regexp(
            "Cores",
            [
                DynamicConfigValue(
                    id_val="color",
                    value={"fixedColor": "light-orange", "mode": "fixed"},
                ),
                DynamicConfigValue(id_val="custom.fillOpacity", value=0),
                DynamicConfigValue(
                    id_val="custom.lineStyle",
                    value={"dash": [10, 10], "fill": "dash"},
                ),
            ],
        )
    )
