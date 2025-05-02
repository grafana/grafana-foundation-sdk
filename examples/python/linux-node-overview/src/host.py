from grafana_foundation_sdk.builders import common, dashboard, stat
from grafana_foundation_sdk.models.common import BigValueColorMode
from grafana_foundation_sdk.models.dashboard import ThresholdsMode, Threshold
from grafana_foundation_sdk.models import units

from .common import prometheus_query, uname_stat, default_stat


def uptime_stat() -> stat.Panel:
    return (
        default_stat()
        .title("Uptime")
        .description(
            "The duration of time that has passed since the last reboot or system start."
        )
        .unit(units.DurationSeconds)
        .with_target(
            prometheus_query(
                query='time() - node_boot_time_seconds{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
            )
        )
        .thresholds(
            dashboard.ThresholdsConfig()
            .mode(ThresholdsMode.ABSOLUTE)
            .steps(
                [
                    Threshold(color="orange"),
                    Threshold(value=600, color="text"),
                ]
            )
        )
    )


def os_stat() -> stat.Panel:
    return (
        default_stat()
        .title("OS")
        .description("Operating system.")
        .with_target(
            prometheus_query(
                query='node_os_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
            )
        )
        .reduce_options(
            common.ReduceDataOptions().calcs(["lastNotNull"]).fields("/^pretty_name$/")
        )
        .color_mode(BigValueColorMode.NONE)
    )


def hostname_stat() -> stat.Panel:
    return uname_stat(
        title="Hostname",
        description="System's hostname.",
        field="nodename",
    )


def kernel_version_stat() -> stat.Panel:
    return uname_stat(
        title="Kernel version",
        description="Kernel version of linux host.",
        field="release",
    )
