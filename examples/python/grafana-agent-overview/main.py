from grafana_foundation_sdk.builders import dashboard
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.models.dashboard import (
    DataSourceRef,
    DashboardCursorSync,
    DashboardLinkType,
    VariableOption,
    VariableRefresh,
    VariableSort,
)

from src import overview, discovery, retrieval

if __name__ == "__main__":
    builder = (
        dashboard.Dashboard("[Example] Grafana Agent Overview")
        .uid("example-grafana-agent-overview")
        .tags(["generated", "grafana-foundation-sdk", "grafana-agent-integration"])
        .editable()
        .tooltip(DashboardCursorSync.CROSSHAIR)
        .refresh("30s")
        .time("now-30m", "now")
        .timezone("browser")
        .timepicker(
            dashboard.TimePicker().refresh_intervals(
                ["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"]
            )
        )
        # Links to other agent-related dashboards
        .link(
            dashboard.DashboardLink("Grafana Agent Dashboards")
            .type(DashboardLinkType.DASHBOARDS)
            .tags(["grafana-agent-integration"])
            .include_vars(True)
            .keep_time(True)
        )
        # "Data source" variable
        .with_variable(
            dashboard.DatasourceVariable("prometheus_datasource")
            .label("Data source")
            .type("prometheus")
            .regex("(?!grafanacloud-usage|grafanacloud-ml-metrics).+")
            .multi(False)
        )
        # "Job" variable
        .with_variable(
            dashboard.QueryVariable("job")
            .label("Job")
            .query("label_values(agent_build_info, job)")
            .datasource(DataSourceRef(uid="$prometheus_datasource"))
            .current(VariableOption(selected=True, text="All", value="$__all"))
            .refresh(VariableRefresh.ON_TIME_RANGE_CHANGED)
            .sort(VariableSort.ALPHABETICAL_ASC)
            .multi(True)
            .include_all(True)
        )
        # "Instance" variable
        .with_variable(
            dashboard.QueryVariable("instance")
            .label("Instance")
            .query('label_values(agent_build_info{job=~"$job"}, instance)')
            .datasource(DataSourceRef(uid="$prometheus_datasource"))
            .current(VariableOption(selected=True, text="All", value="$__all"))
            .refresh(VariableRefresh.ON_TIME_RANGE_CHANGED)
            .sort(VariableSort.ALPHABETICAL_ASC)
            .multi(True)
            .include_all(True)
        )
        # Panels
        .with_row(dashboard.Row("Overview"))
        .with_panel(overview.running_instances_table())
        .with_row(dashboard.Row("Prometheus Discovery"))
        .with_panel(discovery.target_sync_timeseries())
        .with_panel(discovery.target_timeseries())
        .with_row(dashboard.Row("Prometheus Retrieval"))
        .with_panel(retrieval.avg_scrape_interval_duration_timeseries())
        .with_panel(retrieval.scrape_failures_timeseries())
        .with_panel(retrieval.appended_samples_timeseries())
    )

    print(JSONEncoder(sort_keys=True, indent=2).encode(builder.build()))
