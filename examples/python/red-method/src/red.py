from grafana_foundation_sdk.builders import dashboard, testdata, timeseries
from grafana_foundation_sdk.models.dashboard import (
    DataSourceRef,
    DashboardCursorSync,
    DashboardLinkType,
)
from grafana_foundation_sdk.models import units

from .common import default_timeseries


def red(dashboard_title: str, service_ids: list[str]) -> dashboard.Dashboard:
    builder = (
        dashboard.Dashboard("[Example] %s" % dashboard_title)
        .uid("example-red-method")
        .tags(["generated", "red"])
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
        .link(
            dashboard.DashboardLink("RED method")
            .type(DashboardLinkType.LINK)
            .icon("question")
            .target_blank(True)
            .url(
                "https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method"
            )
        )
    )

    for service_id in service_ids:
        builder = (
            builder.with_row(dashboard.Row(service_id))
            .with_panel(request_rate_timeseries().span(8).height(8))
            .with_panel(error_rate_timeseries().span(8).height(8))
            .with_panel(duration_timeseries().span(8).height(8))
        )

    return builder


def request_rate_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Request rate")
        .description("Number of requests handled by the service, per second.")
        .unit(units.RequestsPerSecond)
        .with_target(
            testdata.Dataquery()
            .query_type("randomWalk")
            .datasource(DataSourceRef(uid="grafana", type_val="grafana"))
        )
    )


def error_rate_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Error rate")
        .description("Number of failed requests, per second.")
        .unit(units.RequestsPerSecond)
        .with_target(
            testdata.Dataquery()
            .query_type("randomWalk")
            .datasource(DataSourceRef(uid="grafana", type_val="grafana"))
        )
    )


def duration_timeseries() -> timeseries.Panel:
    return (
        default_timeseries()
        .title("Duration")
        .description("Time taken to process the requests.")
        .unit(units.Seconds)
        .with_target(
            testdata.Dataquery()
            .query_type("randomWalk")
            .datasource(DataSourceRef(uid="grafana", type_val="grafana"))
        )
    )
