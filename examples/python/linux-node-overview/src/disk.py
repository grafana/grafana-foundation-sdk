from grafana_foundation_sdk.builders import dashboard, stat, table, timeseries
from grafana_foundation_sdk.models.common import BigValueColorMode, GraphGradientMode
from grafana_foundation_sdk.models.dashboard import (
    DataSourceRef,
    DynamicConfigValue,
    ThresholdsMode,
    Threshold,
    DataTransformerConfig,
)
from grafana_foundation_sdk.models import units

from .common import (
    prometheus_query,
    table_prometheus_query,
    default_stat,
    default_timeseries,
)


def root_mount_size_stat() -> stat.Panel:
    return (
        default_stat()
        .title("Root mount size")
        .description("Total capacity on the primary mount point /.")
        .with_target(
            prometheus_query(
                query='node_filesystem_size_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mountpoint="/", fstype!="rootfs"}',
            )
        )
        .unit(units.BytesIEC)
        .color_mode(BigValueColorMode.NONE)
    )


def read_write_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Disk reads/writes")
        .description("Disk read/writes in bytes per second.")
        .with_target(
            prometheus_query(
                'irate(node_disk_read_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
                "{{device}} read",
            )
        )
        .with_target(
            prometheus_query(
                'irate(node_disk_written_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
                "{{device}} written",
            )
        )
        .with_target(
            prometheus_query(
                'irate(node_disk_io_time_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
                "{{device}} io util",
            )
        )
        .decimals(0)
        .unit("Bps")
        .unit(units.BytesPerSecondSI)
        .fill_opacity(1)
        .gradient_mode(GraphGradientMode.OPACITY)
        .override_by_regexp(
            "/time|used|busy|util/",
            [
                DynamicConfigValue(id_val="custom.axisSoftMax", value=100),
                DynamicConfigValue(id_val="custom.drawStyle", value="points"),
                DynamicConfigValue(id_val="unit", value=units.Percent),
            ],
        )
    )


def space_usage_table() -> table.Panel:
    return (
        table.Panel()
        .title("Disk space usage")
        .description(
            "Disk utilisation in percent, by mountpoint. Some duplication can occur if the same filesystem is mounted in multiple locations."
        )
        .datasource(DataSourceRef(uid="$datasource"))
        .with_target(
            table_prometheus_query(
                'node_filesystem_size_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                "TOTAL",
            ).instant()
        )
        .with_target(
            table_prometheus_query(
                'node_filesystem_avail_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                "FREE",
            )
            .instant()
            .legend_format("{{ mountpoint }} free")
        )
        .unit(units.BytesIEC)
        .thresholds(
            dashboard.ThresholdsConfig()
            .mode(ThresholdsMode.ABSOLUTE)
            .steps(
                [
                    Threshold(color="light-blue"),
                    Threshold(value=0.8, color="light-yellow"),
                    Threshold(value=0.9, color="light-red"),
                ]
            )
        )
        # Transformations
        .with_transformation(
            DataTransformerConfig(
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
            )
        )
        .with_transformation(DataTransformerConfig(id_val="merge", options={}))
        .with_transformation(
            DataTransformerConfig(
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
            )
        )
        .with_transformation(
            DataTransformerConfig(
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
            )
        )
        .with_transformation(
            DataTransformerConfig(
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
            )
        )
        .with_transformation(
            DataTransformerConfig(
                id_val="sortBy",
                options={
                    "fields": {},
                    "sort": [
                        {"field": "Mounted on", "desc": False},
                    ],
                },
            )
        )
        # Overrides
        .override_by_name(
            "Mounted on",
            [DynamicConfigValue(id_val="custom.width", value="260")],
        )
        .override_by_name(
            "Size",
            [DynamicConfigValue(id_val="custom.width", value="80")],
        )
        .override_by_name(
            "Used",
            [DynamicConfigValue(id_val="custom.width", value="80")],
        )
        .override_by_name(
            "Available",
            [DynamicConfigValue(id_val="custom.width", value="80")],
        )
        .override_by_name(
            "Used, %",
            [
                DynamicConfigValue(
                    id_val="custom.cellOptions",
                    value={"mode": "basic", "type": "gauge"},
                ),
                DynamicConfigValue(id_val="min", value=0),
                DynamicConfigValue(id_val="max", value=1),
                DynamicConfigValue(id_val="unit", value=units.PercentUnit),
            ],
        )
    )
