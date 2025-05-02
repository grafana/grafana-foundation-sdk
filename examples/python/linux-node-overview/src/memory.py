from grafana_foundation_sdk.builders import dashboard, stat, timeseries
from grafana_foundation_sdk.models.common import (
    BigValueColorMode,
    BigValueGraphMode,
    GraphGradientMode,
)
from grafana_foundation_sdk.models.dashboard import (
    DynamicConfigValue,
    ThresholdsMode,
    Threshold,
)
from grafana_foundation_sdk.models import units

from .common import prometheus_query, default_stat, default_timeseries


def total_stat() -> stat.Panel:
    return (
        default_stat()
        .title("Memory total")
        .description(
            "Amount of random-access memory (RAM) installed.\nIt represents the system's available working memory that applications and the operating system use to perform tasks.\nA higher memory total generally leads to better system performance and the ability to run more demanding applications and processes simultaneously.",
        )
        .with_target(
            prometheus_query(
                query='node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
            )
        )
        .decimals(2)
        .unit(units.BytesIEC)
        .color_mode(BigValueColorMode.NONE)
    )


def swap_total_stat() -> stat.Panel:
    return (
        default_stat()
        .title("Total swap")
        .description(
            "Total swap available.\n\nSwap is a space on a storage device (usually a dedicated swap partition or a swap file) \nused as virtual memory when the physical RAM (random-access memory) is fully utilized.\nSwap space helps prevent memory-related performance issues by temporarily transferring less-used data from RAM to disk,\nfreeing up physical memory for active processes and applications.",
        )
        .with_target(
            prometheus_query(
                query='node_memory_SwapTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
            )
        )
        .unit(units.BytesIEC)
        .color_mode(BigValueColorMode.NONE)
    )


def usage_stat() -> stat.Panel:
    query = """100 -
(
  avg by (instance) (node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) /
  avg by (instance) (node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"})
* 100
)"""

    return (
        default_stat()
        .title("Memory usage")
        .description(
            "RAM (random-access memory) currently in use by the operating system and running applications, in percent.",
        )
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
    mem_used_query = """(
  node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
  node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
  node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
  node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
)"""

    return (
        default_timeseries()
        .title("Memory usage")
        .description(
            "- Used: The amount of physical memory currently in use by the system.\n- Cached: The amount of physical memory used for caching data from disk. The Linux kernel uses available memory to cache data that is read from or written to disk. This helps speed up disk access times.\n- Free: The amount of physical memory that is currently not in use.\n- Buffers: The amount of physical memory used for temporary storage of data being transferred between devices or applications.\n- Available: The amount of physical memory that is available for use by applications. This takes into account memory that is currently being used for caching but can be freed up if needed.",
        )
        .span(18)
        .with_target(prometheus_query(query=mem_used_query, legend="Memory used"))
        .with_target(
            prometheus_query(
                'node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                "Memory cached",
            )
        )
        .with_target(
            prometheus_query(
                'node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                "Memory available",
            )
        )
        .with_target(
            prometheus_query(
                'node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                "Memory buffers",
            )
        )
        .with_target(
            prometheus_query(
                'node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                "Memory free",
            )
        )
        .with_target(
            prometheus_query(
                'node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                "Memory total",
            )
        )
        .decimals(2)
        .unit(units.BytesIEC)
        .gradient_mode(GraphGradientMode.OPACITY)
        .override_by_regexp(
            ".*(T|t)otal.*",
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
