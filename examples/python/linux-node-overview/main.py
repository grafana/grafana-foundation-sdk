from grafana_foundation_sdk.builders import dashboard
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.models.dashboard import (
    DashboardCursorSync,
    DashboardLinkType,
    VariableHide,
)

from src import common, cpu, disk, host, memory, network

builder = (
    dashboard.Dashboard("[Example] Linux node / overview")
    .uid("example-integration-linux-node")
    .tags(["generated", "grafana-foundation-sdk", "linux-node-integration"])
    .editable()
    .tooltip(DashboardCursorSync.OFF)
    .refresh("30s")
    .time("now-30m", "now")
    .timezone("browser")
    .timepicker(
        dashboard.TimePicker().refresh_intervals(
            ["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"]
        )
    )
    # "Back to fleet" link
    .link(
        dashboard.DashboardLink("Back to Linux node / fleet")
        .type(DashboardLinkType.LINK)
        .url("/d/node-fleet")
        .tags(["grafana-agent-integration"])
        .keep_time(True)
    )
    # Links to other linux-node-related dashboards
    .link(
        dashboard.DashboardLink("All Linux node /  dashboards")
        .type(DashboardLinkType.DASHBOARDS)
        .tags(["linux-node-integration"])
        .include_vars(True)
        .as_dropdown(True)
        .keep_time(True)
    )
    # "Data source" variable
    .with_variable(
        dashboard.DatasourceVariable("datasource")
        .label("Data source")
        .type("prometheus")
        .regex("(?!grafanacloud-usage|grafanacloud-ml-metrics).+")
        .multi(False)
    )
    # "Loki data source" variable
    .with_variable(
        dashboard.DatasourceVariable("loki_datasource")
        .label("Loki data source")
        .type("loki")
        .regex("(?!grafanacloud.+usage-insights|grafanacloud.+alert-state-history).+")
        .multi(False)
        .hide(VariableHide.HIDE_VARIABLE)
    )
    # "Cluster" variable
    .with_variable(
        common.query_variable(
            "cluster",
            "Cluster",
            'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)"}, cluster)',
        ).all_value(".*")
    )
    # "Job" variable
    .with_variable(
        common.query_variable(
            "job",
            "Job",
            'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster"}, job)',
        ).all_value(".+")
    )
    # "Instance" variable
    .with_variable(
        common.query_variable(
            "instance",
            "Instance",
            'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job"}, instance)',
        ).all_value(".+")
    )
    # Panels
    .with_row(dashboard.Row("Overview"))
    .with_panel(host.uptime_stat())
    .with_panel(host.hostname_stat())
    .with_panel(host.kernel_version_stat())
    .with_panel(host.os_stat())
    .with_panel(cpu.count_stat())
    .with_panel(memory.total_stat())
    .with_panel(memory.swap_total_stat())
    .with_panel(disk.root_mount_size_stat())
    .with_row(dashboard.Row("CPU"))
    .with_panel(cpu.usage_stat().height(6))
    .with_panel(cpu.usage_timeseries().height(6))
    .with_panel(cpu.load_average_timeseries().height(6).span(6))
    .with_row(dashboard.Row("Memory"))
    .with_panel(memory.usage_stat().height(6))
    .with_panel(memory.usage_timeseries().height(6))
    .with_row(dashboard.Row("Disk"))
    .with_panel(disk.read_write_timeseries().height(8))
    .with_panel(disk.space_usage_table().height(8))
    .with_row(dashboard.Row("Network"))
    .with_panel(network.traffic_timeseries().height(8))
    .with_panel(network.errors_timeseries().height(8))
)

if __name__ == "__main__":
    print(JSONEncoder(sort_keys=True, indent=2).encode(builder.build()))
