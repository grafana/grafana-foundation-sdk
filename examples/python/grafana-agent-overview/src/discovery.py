from grafana_foundation_sdk.builders import timeseries
from grafana_foundation_sdk.models.dashboard import DataSourceRef
from grafana_foundation_sdk.models import units

from .common import prometheus_query, default_timeseries


def target_sync_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Target Sync")
        .description("Actual interval to sync the scrape pool.")
        .datasource(DataSourceRef(uid="$prometheus_datasource"))
        .unit(units.Seconds)
        .with_target(
            prometheus_query(
                query='sum(rate(prometheus_target_sync_length_seconds_sum{job=~"$job", instance=~"$instance"}[$__rate_interval])) by (instance, scrape_job)',
                legend="{{instance}}/{{scrape_job}}",
            )
        )
    )


def target_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Targets")
        .description("Discovered targets by prometheus service discovery.")
        .datasource(DataSourceRef(uid="$prometheus_datasource"))
        .unit(units.Short)
        .with_target(
            prometheus_query(
                query='sum by (instance) (prometheus_sd_discovered_targets{job=~"$job", instance=~"$instance"})',
                legend="{{instance}}",
            )
        )
    )
