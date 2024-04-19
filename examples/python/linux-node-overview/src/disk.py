from grafana_foundation_sdk.builders import dashboard, stat, table, timeseries
from grafana_foundation_sdk.models.common import BigValueColorMode, GraphGradientMode
from grafana_foundation_sdk.models.dashboard import (
    DataSourceRef, DynamicConfigValue, MatcherConfig, ThresholdsMode, Threshold,
    DataTransformerConfig,
)

from .common import prometheus_query, table_prometheus_query, default_stat, default_timeseries


def root_mount_size_stat() -> stat.Panel:
    return (
        default_stat()
        .title("Root mount size")
        .description("Total capacity on the primary mount point /.")
        .with_target(prometheus_query(
            query='node_filesystem_size_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mountpoint="/", fstype!="rootfs"}',
        ))
        .unit("bytes")
        .color_mode(BigValueColorMode.NONE)
    )


def read_write_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Disk reads/writes")
        .description("Disk read/writes in bytes per second.")
        .with_target(prometheus_query(
            'irate(node_disk_read_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
            "{{device}} read",
        ))
        .with_target(prometheus_query(
            'irate(node_disk_written_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
            "{{device}} written",
        ))
        .with_target(prometheus_query(
            'irate(node_disk_io_time_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
            "{{device}} io util",
        ))
        .decimals(0)
        .unit("Bps")
        .fill_opacity(1)
        .gradient_mode(GraphGradientMode.OPACITY)
        .with_override(
            matcher=MatcherConfig(id_val="byRegexp", options="/time|used|busy|util/"),
            properties=[
                DynamicConfigValue(id_val="custom.axisSoftMax", value=100),
                DynamicConfigValue(id_val="custom.drawStyle", value="points"),
                DynamicConfigValue(id_val="unit", value="percent"),
            ],
        )
    )


def space_usage_table() -> table.Panel:
    return (
        table.Panel()
        .title("Disk space usage")
        .description(
            "Disk utilisation in percent, by mountpoint. Some duplication can occur if the same filesystem is mounted in multiple locations.")
        .datasource(DataSourceRef(uid="$datasource"))
        .with_target(table_prometheus_query(
            'node_filesystem_size_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
            "TOTAL",
        ).instant(True))
        .with_target(table_prometheus_query(
            'node_filesystem_avail_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
            "FREE",
        ).instant(True).legend_format("{{ mountpoint }} free"))
        .unit("bytes")
        .thresholds(
            dashboard.ThresholdsConfig()
            .mode(ThresholdsMode.ABSOLUTE)
            .steps([
                Threshold(color="light-blue"),
                Threshold(value=0.8, color="light-yellow"),
                Threshold(value=0.9, color="light-red"),
            ])
        )
        # Transformations
        .with_transformation(DataTransformerConfig(
            id_val="groupBy",
            options={
                "fields": {
                    "Value #FREE": {
                        "aggregations": ["lastNotNull"],
                        "operation": "aggregate",
                    },
                    "Value #TOTAL": {
                        "aggregations": ["lastNotNull"],
                        "operation": "aggregate",
                    },
                    "mountpoint": {
                        "aggregations": [],
                        "operation": "groupby",
                    },
                },
            },
        ))
        .with_transformation(DataTransformerConfig(id_val="merge", options={}))
        .with_transformation(DataTransformerConfig(
            id_val="calculateField",
            options={
                "alias": "Used",
                "binary": {
                    "left": "Value #TOTAL (lastNotNull)",
                    "operator": "-",
                    "reducer": "sum",
                    "right": "Value #FREE (lastNotNull)",
                },
                "mode": "binary",
                "reduce": {"reducer": "sum"},
            },
        ))
        .with_transformation(DataTransformerConfig(
            id_val="calculateField",
            options={
                "alias": "Used, %",
                "binary": {
                    "left": "Used",
                    "operator": "/",
                    "reducer": "sum",
                    "right": "Value #TOTAL (lastNotNull)",
                },
                "mode": "binary",
                "reduce": {"reducer": "sum"},
            },
        ))
        .with_transformation(DataTransformerConfig(
            id_val="organize",
            options={
                "excludeByName": {},
                "indexByName": {
                    "Used": 3,
                    "Used, %": 4,
                    "Value #FREE (lastNotNull)": 2,
                    "Value #TOTAL (lastNotNull)": 1,
                    "mountpoint": 0,
                },
                "renameByName": {
                    "Value #FREE (lastNotNull)": "Available",
                    "Value #TOTAL (lastNotNull)": "Size",
                    "mountpoint": "Mounted on",
                },
            },
        ))
        .with_transformation(DataTransformerConfig(
            id_val="sortBy",
            options={
                "fields": {},
                "sort": [
                    {"field": "Mounted on", "desc": False},
                ],
            },
        ))
        # Overrides
        .with_override(
            matcher=MatcherConfig(id_val="byName", options="Mounted on"),
            properties=[DynamicConfigValue(id_val="custom.width", value="260")],
        )
        .with_override(
            matcher=MatcherConfig(id_val="byName", options="Size"),
            properties=[DynamicConfigValue(id_val="custom.width", value="80")],
        )
        .with_override(
            matcher=MatcherConfig(id_val="byName", options="Used"),
            properties=[DynamicConfigValue(id_val="custom.width", value="80")],
        )
        .with_override(
            matcher=MatcherConfig(id_val="byName", options="Available"),
            properties=[DynamicConfigValue(id_val="custom.width", value="80")],
        )
        .with_override(
            matcher=MatcherConfig(id_val="byName", options="Used, %"),
            properties=[
                DynamicConfigValue(
                    id_val="custom.cellOptions",
                    value={
                        "mode": "basic",
                        "type": "gauge",
                    },
                ),
                DynamicConfigValue(id_val="min", value=0),
                DynamicConfigValue(id_val="max", value=1),
                DynamicConfigValue(id_val="unit", value="percentunit"),
            ],
        )
    )
