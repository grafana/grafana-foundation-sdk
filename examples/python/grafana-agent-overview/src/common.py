from grafana_foundation_sdk.builders import (
    timeseries,
    common as common_builder,
    prometheus,
)
from grafana_foundation_sdk.models import common, prometheus as prom


def prometheus_query(query: str, legend: str) -> prometheus.Dataquery:
    return prometheus.Dataquery().expr(query).legend_format(legend)


def table_prometheus_query(query: str, ref_id: str) -> prometheus.Dataquery:
    return (
        prometheus.Dataquery()
        .expr(query)
        .format(prom.PromQueryFormat.TABLE)
        .instant()
        .interval_factor(2)
        .ref_id(ref_id)
    )


def default_timeseries() -> timeseries.Panel:
    return (
        timeseries.Panel()
        .height(7)
        .span(12)
        .line_width(1)
        .fill_opacity(0)
        .point_size(5)
        .show_points(common.VisibilityMode.AUTO)
        .draw_style(common.GraphDrawStyle.LINE)
        .gradient_mode(common.GraphGradientMode.NONE)
        .span_nulls(False)
        .axis_border_show(False)
        .legend(
            common_builder.VizLegendOptions()
            .display_mode(common.LegendDisplayMode.LIST)
            .placement(common.LegendPlacement.BOTTOM)
            .show_legend(True)
        )
        .tooltip(
            common_builder.VizTooltipOptions()
            .mode(common.TooltipDisplayMode.SINGLE)
            .sort(common.SortOrder.NONE)
        )
        .thresholds_style(
            common_builder.GraphThresholdsStyleConfig().mode(
                common.GraphThresholdsStyleMode.OFF
            )
        )
    )
