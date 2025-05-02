from grafana_foundation_sdk.builders import timeseries
from grafana_foundation_sdk.models.dashboard import DataSourceRef
from grafana_foundation_sdk.models import units

from .common import prometheus_query, default_timeseries


def avg_scrape_interval_duration_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Average Scrape Interval Duration")
        .description("Actual intervals between scrapes.")
        .datasource(DataSourceRef(uid="$prometheus_datasource"))
        .unit(units.Seconds)
        .with_target(
            prometheus_query(
                query='rate(prometheus_target_interval_length_seconds_sum{job=~"$job", instance=~"$instance"}[$__rate_interval]) / rate(prometheus_target_interval_length_seconds_count{job=~"$job", instance=~"$instance"}[$__rate_interval])',
                legend="{{instance}} {{interval}} configured",
            ).interval_factor(2)
        )
    )


def scrape_failures_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Scrape failures")
        .description(
            "Shows all scrape failures (sample limit exceeded, duplicate, out of bounds, out of order)."
        )
        .datasource(DataSourceRef(uid="$prometheus_datasource"))
        .unit(units.Short)
        .with_target(
            prometheus_query(
                query='sum by (job) (rate(prometheus_target_scrapes_exceeded_sample_limit_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                legend="exceeded sample limit: {{job}}",
            )
        )
        .with_target(
            prometheus_query(
                query='sum by (job) (rate(prometheus_target_scrapes_sample_duplicate_timestamp_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                legend="duplicate timestamp: {{job}}",
            )
        )
        .with_target(
            prometheus_query(
                query='sum by (job) (rate(prometheus_target_scrapes_sample_out_of_bounds_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                legend="out of bounds: {{job}}",
            )
        )
        .with_target(
            prometheus_query(
                query='sum by (job) (rate(prometheus_target_scrapes_sample_out_of_order_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                legend="out of order: {{job}}",
            )
        )
    )


def appended_samples_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Appended Samples")
        .description("Total number of samples appended to the WAL.")
        .datasource(DataSourceRef(uid="$prometheus_datasource"))
        .unit(units.Short)
        .with_target(
            prometheus_query(
                query='sum by (job, instance_group_name) (rate(agent_wal_samples_appended_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                legend="{{job}} {{instance_group_name}}",
            )
        )
    )
